package myMuetx

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// WaitGroup 1.20直接操作底层指针获取
type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) CountCounter() uint64 {
	state := atomic.LoadUint64((*uint64)(unsafe.Pointer(&wg.WaitGroup)))
	v := uint64(state >> 32)
	return v
}
func (wg *WaitGroup) CountWaiter() uint64 {
	state := atomic.LoadUint64((*uint64)(unsafe.Pointer(&wg.WaitGroup)))
	v := state & 0xFFFFFFFF
	return v
}

// TrueWaitGroup 这个才是兼容多版本的添加功能后的并发原语
type TrueWaitGroup struct {
	sync.WaitGroup
	count uint32
}

func (wg *TrueWaitGroup) Add(delta int) {
	wg.WaitGroup.Add(delta)
	atomic.AddUint32(&wg.count, uint32(delta))
}

func (wg *TrueWaitGroup) Done() {
	wg.WaitGroup.Done()
	atomic.AddUint32(&wg.count, -1)
}

func (wg *TrueWaitGroup) Count() uint32 {
	return atomic.LoadUint32(&wg.count)
}
