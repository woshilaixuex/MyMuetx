package myMuetx

import "sync"

type RWMutex struct {
	w           sync.RWMutex
	writerSem   uint32
	readerSem   uint32
	readerCount int32
	readerWait  int32
}

const rwMutexMaxReaders = 1 << 30
