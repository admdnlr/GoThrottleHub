package ratelimiter

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// Limiter struct'ı her IP adresi için bir rate.Limiter tutar.
type Limiter struct {
	visitors map[string]*rate.Limiter
	mtx      sync.Mutex
}

// Yeni bir Limiter örneği oluşturur.
func NewLimiter(r rate.Limit, b int) *Limiter {
	return &Limiter{
		visitors: make(map[string]*rate.Limiter),
	}
}

// GetLimiter IP adresine özgü rate.Limiter döndürür, eğer yoksa yenisini oluşturur.
func (l *Limiter) GetLimiter(ip string) *rate.Limiter {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	limiter, exists := l.visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(1), 5) // Her IP için saniyede 1 token, burst boyutu 5
		l.visitors[ip] = limiter
	}

	return limiter
}

// Middleware oran limitlendiriciyi HTTP isteklerine uygular.
func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		if !l.GetLimiter(ip).Allow() {
			http.Error(w, "Çok fazla istek yaptınız, lütfen sonra tekrar deneyin.", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
