package server

import (
	"runtime"
	"runtime/debug"
	"time"
)

var LastGCTime time.Time

const (
	MENLIMIT = 1024 * 1024 * 1024 * 6
)

type finalizer struct {
	ch chan time.Time
	ref *finalizerRef
}

type finalizerRef struct {
	parent *finalizer
}

func finalizerHandle(f *finalizerRef) {
	var memStats = &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	now := time.Now()
	gcSubTime := now.Sub(LastGCTime)
	LastGCTime = now

	if gcSubTime.Minutes() < 1.97 {
		perSecondsMem := 1024 * 1024//memStats.NextGC / ( uint64( gcSubTime.Seconds()/2 ) )
		twoMinuteMem := perSecondsMem * 120

		var gcPercentge uint64
		if twoMinuteMem > MENLIMIT {
			gcPercentge = MENLIMIT / memStats.NextGC
		} else {
			gcPercentge = uint64( twoMinuteMem ) / memStats.NextGC
		}

		debug.SetGCPercent( int(gcPercentge) * 100 )
	}

	runtime.SetFinalizer(f, finalizerHandle)
}

func GCTicker() {
	f := &finalizer{
		ch: make(chan time.Time, 1),
	}
	f.ref = &finalizerRef{parent:f}

	runtime.SetFinalizer(f.ref, finalizerHandle)

	f.ref = nil
}