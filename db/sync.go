package db

type brokenLocker struct{}

func newBrokenLocker() brokenLocker { return brokenLocker{} }

func (brokenLocker) Lock() {}

func (brokenLocker) Unlock() {}
