package queue

import (
	"errors"
	"sync"
)

// Request türü, kuyruğa konulacak istekleri temsil eder.
type Request struct {
	// İstekle ilgili verileri buraya ekleyin
	Payload interface{}
}

// Queue türü, talepleri sıraya almak için kullanılır.
type Queue struct {
	Requests     chan Request
	maxQueueSize int
	mtx          sync.Mutex
}

// NewQueue, yeni bir talep kuyruğu örneği oluşturur.
func NewQueue(maxQueueSize int) *Queue {
	return &Queue{
		Requests:     make(chan Request, maxQueueSize),
		maxQueueSize: maxQueueSize,
	}
}

// Enqueue, yeni bir talebi kuyruğa ekler.
func (q *Queue) Enqueue(request Request) error {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.Requests) == q.maxQueueSize {
		return errors.New("kuyruk dolu")
	}
	q.Requests <- request
	return nil
}

// Dequeue, kuyruktan bir talebi çıkarır ve döndürür.
func (q *Queue) Dequeue() (Request, error) {
	select {
	case req := <-q.Requests:
		return req, nil
	default:
		return Request{}, errors.New("kuyruk boş")
	}
}

// ProcessQueue, kuyruktaki talepleri işleyen bir fonksiyondur.
// Bu fonksiyon, genellikle bir goroutine içinde çalıştırılmalıdır.
func (q *Queue) ProcessQueue(processFunc func(Request)) {
	for req := range q.Requests {
		processFunc(req)
	}
}
