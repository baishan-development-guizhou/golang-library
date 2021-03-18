package oticker

import (
	"context"
	"time"
)

//TickerAdvance execute once before and execute functions in the interval time cycle
func TickerAdvance(interval time.Duration, function func()) {
	TickerAdvanceCtx(context.Background(), interval, function)
}

//TickerAdvanceCtx execute functions in the interval time cycle
func TickerAdvanceCtx(ctx context.Context, interval time.Duration, function func()) {
	function()
	TickerCtx(ctx, interval, function)
}

//Ticker execute functions in the interval time cycle
func Ticker(interval time.Duration, function func()) {
	TickerCtx(context.Background(), interval, function)
}

//TickerCtx execute functions in the interval time cycle
func TickerCtx(ctx context.Context, interval time.Duration, function func()) {
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
