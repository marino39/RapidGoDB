package db

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type RWLocker interface {
	sync.Locker
	RLock()
	RUnlock()
}

type brokenLocker struct{}

func newBrokenLocker() brokenLocker { return brokenLocker{} }

func (brokenLocker) Lock() {}

func (brokenLocker) Unlock() {}

func (brokenLocker) RLock() {}

func (brokenLocker) RUnlock() {}

type eagerLocker struct {
	val uint32
}

func newEagerLocker() *eagerLocker { return &eagerLocker{} }

func (el *eagerLocker) Lock() {
	for !atomic.CompareAndSwapUint32(&el.val, 0, 1) {
		runtime.Gosched()
	}
}

func (el *eagerLocker) Unlock() {
	atomic.StoreUint32(&el.val, 0)
}

func (el *eagerLocker) RLock() {
	el.Lock()
}

func (el *eagerLocker) RUnlock() {
	el.Unlock()
}

type locker struct {
	mutex sync.Mutex
}

func newLocker() *locker { return &locker{} }

func (l *locker) Lock() {
	l.mutex.Lock()
}

func (l *locker) Unlock() {
	l.mutex.Unlock()
}

func (l *locker) RLock() {
	l.mutex.Lock()
}

func (l *locker) RUnlock() {
	l.mutex.Unlock()
}
