package common

import (
	"context"
	"time"
)

type _ticker uint

func (t _ticker) TickerAdvance(interval time.Duration, function func()) {
	function()
	t.Ticker(context.Background(), interval, function)
}

func (t _ticker) TickerDelayed(interval time.Duration, function func()) {
	t.Ticker(context.Background(), interval, function)
}

func (t _ticker) Ticker(ctx context.Context, interval time.Duration, function func()) {
	go func() {
		eventsTick := time.NewTicker(interval)
		defer eventsTick.Stop()
		for {
			select {
			case <-eventsTick.C:
				function()
			case <-ctx.Done():
				return
			}
		}
	}()
}
