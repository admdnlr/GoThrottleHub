// middleware.go
package http

import (
	"net/http"

	"github.com/admdnlr/GoThrottleHub/internal/queue"
	"github.com/admdnlr/GoThrottleHub/internal/ratelimiter"
)

// RateLimiterMiddleware, oran limitlendiriciyi HTTP isteklerine uygular.
func RateLimiterMiddleware(limiter *ratelimiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.GetLimiter(r.RemoteAddr).Allow() {
				http.Error(w, "Çok fazla istek yaptınız, lütfen sonra tekrar deneyin.", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// QueueMiddleware, gelen istekleri talep kuyruğuna ekler.
func QueueMiddleware(q *queue.Queue) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := queue.Request{Payload: r} // Burada Payload, http.Request olabilir veya isteği temsil eden başka bir yapı olabilir.
			err := q.Enqueue(req)
			if err != nil {
				http.Error(w, "Şu anda işlem kapasitemiz dolu, lütfen daha sonra tekrar deneyiniz.", http.StatusServiceUnavailable)
				return
			}

			// Kuyrukta işlenmeyi bekleyen isteklere göre burada bir mekanizma uygulamanız gerekecek.
			// Bu, kuyruk tarafından bir goroutine'de işlenmesini bekleyebilir veya hemen işleyebilir.
			// Bu örnek, kuyruk işlemesini simüle etmek için sadece bir sonraki handler'ı çağırır.
			next.ServeHTTP(w, r)
		})
	}
}
