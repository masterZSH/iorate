package iorate

import (
	"context"
	"io"
	"time"

	"golang.org/x/time/rate"
)

type LimitWriter struct {
	w       io.Writer
	limiter *rate.Limiter
	ctx     context.Context
}

func (w *LimitWriter) Write(p []byte) (n int, err error) {
	if w.limiter == nil {
		return w.w.Write(p)
	}
	n, err = w.w.Write(p)
	if err != nil {
		return
	}
	err = w.limiter.WaitN(w.ctx, n)
	return n, err
}

func NewWriter(w io.Writer, limit ...float64) *LimitWriter {
	writer := &LimitWriter{
		w:   w,
		ctx: context.Background(),
	}
	if len(limit) > 0 {
		writer.limiter = rate.NewLimiter(rate.Limit(limit[0]), defaultBursts)
		if res := writer.limiter.ReserveN(time.Now(), defaultBursts); res.OK() {
			time.Sleep(res.Delay())
		}
	}
	return writer
}

func NewWriterWithContext(w io.Writer, ctx context.Context, limit ...float64) *LimitWriter {
	writer := &LimitWriter{
		w:   w,
		ctx: ctx,
	}
	if len(limit) > 0 {
		writer.limiter = rate.NewLimiter(rate.Limit(limit[0]), defaultBursts)
	}
	return writer
}

func (r *LimitWriter) SetLimit(limit float64) {
	r.limiter = rate.NewLimiter(rate.Limit(limit), defaultBursts)
	if res := r.limiter.ReserveN(time.Now(), defaultBursts); res.OK() {
		time.Sleep(res.Delay())
	}
}

func (r *LimitWriter) Clean() {
	r.limiter = nil
}
