package parawalk

import (
	"io/fs"
	"path/filepath"
	"sync"

	"go.uber.org/multierr"
)

type WalkFunc func(path string, info fs.FileInfo, err error) error

func Walk(root string, fn WalkFunc) error {
	var wg sync.WaitGroup
	errCh := make(chan error)
	rerr := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		// エラー処理が必要またはディレクトリの場合はそのまま処理
		if err != nil || info.IsDir() {
			return fn(path, info, err)
		}

		// ファイルの場合はゴールーチンで処理
		wg.Add(1)
		go func() {
			// TODO: deferでwg.Doneを呼ぶ
			defer wg.Done()

			err := fn(path, info, err)
			if err != nil {
				// TODO: エラーチャネル(errCh)にエラーを送信
				errCh <- err
			}
		}()

		return nil
	})

	// TODO: 以下の関数呼び出しを別のゴールーチンで起動する
	go func() {
		for err := range errCh {
			rerr = multierr.Append(rerr, err)
		}
	}()

	wg.Wait()
	// TODO: errChをクローズする
	close(errCh)
	return rerr
}
