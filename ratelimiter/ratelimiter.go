package ratelimiter

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	*rate.Limiter

	sema chan struct{}
}

// NewLimiter returns a new HostLimiter that allows events up to rate r and permits
// bursts of at most b tokens per host.
func NewLimiter(r rate.Limit, b int) *Limiter {
	return &Limiter{
		Limiter: rate.NewLimiter(r, b),
		sema:    make(chan struct{}, 100),
	}
}

func (l *Limiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiterCtx, cancel := context.WithTimeout(r.Context(), time.Second*15)
		defer cancel()

		err := l.Wait(limiterCtx)
		if err != nil {
			w.Header().Set("Retry-After", time.Now().Add(time.Second*120).Format(http.TimeFormat))
			w.Header().Add("Retry-After", "120")

			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		if l.sema != nil {
			select {
			case <-limiterCtx.Done():
				w.Header().Set("Retry-After", time.Now().Add(time.Second*120).Format(http.TimeFormat))
				w.Header().Add("Retry-After", "120")

				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			case l.sema <- struct{}{}:
				defer func() {
					select {
					case <-l.sema:
						//
					default:
						//
					}
				}()
			}
		}

		next.ServeHTTP(w, r)
	})
}
