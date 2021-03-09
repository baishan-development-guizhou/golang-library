package ticker

import (
	"context"
	"time"
)

func TickerAdvance(interval time.Duration, function func()) {
	function()
	Ticker(context.Background(), interval, function)
}

func TickerDelayed(interval time.Duration, function func()) {
	Ticker(context.Background(), interval, function)
}

func Ticker(ctx context.Context, interval time.Duration, function func()) {
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
