package logs

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TraceID represents unique 128bit identifier of a trace
type TraceID struct {
	High, Low uint64
}

func (t TraceID) String() string {
	return fmt.Sprintf("%016x%016x", t.High, t.Low)
}

// NewTraceID generate rand traceID
func NewTraceID() string {
	trace := TraceID{
		High: Generator(),
		Low:  Generator(),
	}
	return trace.String()
}

// Generator generate rand uint64
func Generator() uint64 {
	seedGenerator := NewRand(time.Now().UnixNano())
	pool := sync.Pool{
		New: func() interface{} {
			return rand.NewSource(seedGenerator.Int63())
		},
	}
	generator := pool.Get().(rand.Source)
	number := uint64(generator.Int63())
	pool.Put(generator)
	return number
}
