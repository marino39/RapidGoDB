package db

import (
	"sync"
	"testing"
)

func TestEagerLockerTest(t *testing.T) {
	locker := newEagerLocker()
	lockedTwice := false

	locker.Lock()

	go func() {
		locker.Lock()
		lockedTwice = true
	}()

	if lockedTwice {
		t.Fail()
	}
}

func BenchmarkEagerLockerSimpleTest(b *testing.B) {
	locker := newEagerLocker()
	for n := 0; n < b.N; n++ {
		locker.Lock()
		locker.Unlock()
	}
}

func BenchmarkMutexSimpleTest(b *testing.B) {
	locker := sync.Mutex{}
	for n := 0; n < b.N; n++ {
		locker.Lock()
		locker.Unlock()
	}
}

func BenchmarkEagerLockerTest(b *testing.B) {
	locker := newEagerLocker()
	for n := 0; n < b.N; n++ {
		locker.Lock()
		go func() { locker.Unlock() }()
		locker.Lock()
		locker.Unlock()
	}
}

func BenchmarkMutexTest(b *testing.B) {
	locker := sync.Mutex{}
	for n := 0; n < b.N; n++ {
		locker.Lock()
		go func() { locker.Unlock() }()
		locker.Lock()
		locker.Unlock()
	}
}
