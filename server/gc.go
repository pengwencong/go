package server

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

const (
	ALLMEN = 1024 * 1024 * 1024 * 6
)

type finalizer struct {
	ch chan time.Time
	ref *finalizerRef
}

type finalizerRef struct {
	parent *finalizer
}

func finalizerHandle(f *finalizerRef) {

	//select {
	//case <- f.parent.ch:
	//
	//}

	var memStats = &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	fmt.Printf("GC: %+v\n", memStats)

	if (memStats.Alloc * 200) > ALLMEN {
		multi := ALLMEN / memStats.Alloc
		debug.SetGCPercent( int(multi) )
	} else {
		debug.SetGCPercent(200)
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