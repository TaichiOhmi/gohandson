package eventcal

import "time"

var DefaultClock Clock = ClockFunc(time.Now) /* TODO: time.Now関数をClockFunc型に変換して代入する */

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

// TODO: ClockFunc型にClockインタフェースを実装させる
func (cf ClockFunc) Now() time.Time {
	return cf()
}
