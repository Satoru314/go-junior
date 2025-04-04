package middlewares

import (
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	var no int
	mu.Lock()
	no = logNo
	logNo++
	mu.Unlock()
	return no
}
