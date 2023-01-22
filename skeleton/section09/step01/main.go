package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	const url = "https://connpass.com/api/v1/event/?keyword=golang"
	// TODO: GETメソッドのリクエストを生成する
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// TODO: http.DefaultClientでリクエストを送る
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 標準出力にダンプする
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return err
	}

	return nil
}
