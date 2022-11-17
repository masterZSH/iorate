package iorate

import (
	"context"
	"io"
	"time"

	"golang.org/x/time/rate"
)

type LimitReader struct {
	r       io.Reader
	limiter *rate.Limiter
	ctx     context.Context
}

func (r *LimitReader) Read(p []byte) (n int, err error) {
	if r.limiter == nil {
		return r.r.Read(p)
	}
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	err = r.limiter.WaitN(r.ctx, n)
	return n, err
}

func NewReader(r io.Reader, limit ...float64) *LimitReader {
	reader := &LimitReader{
		r:   r,
		ctx: context.Background(),
	}
	if len(limit) > 0 {
		reader.limiter = rate.NewLimiter(rate.Limit(limit[0]), defaultBursts)
		if res := reader.limiter.ReserveN(time.Now(), defaultBursts); res.OK() {
			time.Sleep(res.Delay())
		}
	}
	return reader
}

func NewReaderWithContext(r io.Reader, ctx context.Context, limit ...float64) *LimitReader {
	reader := &LimitReader{
		r:   r,
		ctx: ctx,
	}
	// 每秒限制数
	if len(limit) > 0 {
		reader.limiter = rate.NewLimiter(rate.Limit(limit[0]), defaultBursts)
	}
	return reader
}

func (r *LimitReader) SetLimit(limit float64) {
	r.limiter = rate.NewLimiter(rate.Limit(limit), defaultBursts)
	if res := r.limiter.ReserveN(time.Now(), defaultBursts); res.OK() {
		time.Sleep(res.Delay())
	}
}

func (r *LimitReader) Clean() {
	r.limiter = nil
}
