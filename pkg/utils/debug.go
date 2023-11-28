package utils

import (
	"fmt"
	"sync"
	"time"
)

var TCCPrinterMutex sync.Mutex

type TCCPrinter func(format string, args ...interface{})

type TCCPointDetail struct {
	name string
	time int64
}

type TCCPrintType uint

const (
	TCCPrintTypeNs TCCPrintType = iota
	TCCPrintTypeMs
	TCCPrintTypeS
)

type TimeCostCounter struct {
	printLevel TCCPrintType
	tag        string
	start      int64
	points     []*TCCPointDetail
	end        int64
}

func NewMsTcc(tag string) *TimeCostCounter {
	return NewTccWithPrintLevel(tag, TCCPrintTypeMs)
}

func NewTccWithPrintLevel(tag string, tp TCCPrintType) *TimeCostCounter {
	return &TimeCostCounter{
		printLevel: tp,
		tag:        tag,
	}
}

func (tcc *TimeCostCounter) Start() *TimeCostCounter {
	tcc.start = time.Now().UnixNano()
	return tcc
}

func (tcc *TimeCostCounter) Point(pName string) *TimeCostCounter {
	tcc.points = append(tcc.points, &TCCPointDetail{
		name: pName,
		time: time.Now().UnixNano(),
	})
	return tcc
}

func (tcc *TimeCostCounter) End() *TimeCostCounter {
	tcc.end = time.Now().UnixNano()
	return tcc
}

func (tcc *TimeCostCounter) Print() {
	go func() {
		TCCPrinterMutex.Lock()
		prevPointTime := tcc.start
		for _, point := range tcc.points {
			pointNanoSeconds := float64(point.time - prevPointTime)
			tcc.PrintNs("[**TCC** -%v:%v-] cost time : %fns\n", tcc.tag, point.name, pointNanoSeconds)
			tcc.PrintMs("[**TCC** -%v:%v-] cost time(ms) :[%fns %fms]\n", tcc.tag, point.name, pointNanoSeconds, pointNanoSeconds/1e6)
			tcc.PrintS("[**TCC** -%v:%v-] cost time(s) :[%fns %fms %fs]\n", tcc.tag, point.name, pointNanoSeconds, pointNanoSeconds/1e6, pointNanoSeconds/1e9)
			prevPointTime = point.time
		}
		nanoSeconds := float64(tcc.end - tcc.start)
		tcc.PrintNs("[**TCC** -%v-] cost time : %vns\n", tcc.tag, nanoSeconds)
		tcc.PrintMs("[**TCC** -%v-] cost time(ms) :%vms\n", tcc.tag, nanoSeconds/1e6)
		tcc.PrintS("[**TCC** -%v-] cost time(s) :%vs\n", tcc.tag, nanoSeconds/1e9)

		TCCPrinterMutex.Unlock()
	}()
}

func (tcc *TimeCostCounter) PrintNs(format string, args ...interface{}) {
	if tcc.printLevel == TCCPrintTypeNs {
		fmt.Printf(format, args...)
	}
}

func (tcc *TimeCostCounter) PrintMs(format string, args ...interface{}) {
	if tcc.printLevel == TCCPrintTypeMs {
		fmt.Printf(format, args...)
	}
}

func (tcc *TimeCostCounter) PrintS(format string, args ...interface{}) {
	if tcc.printLevel == TCCPrintTypeS {
		fmt.Printf(format, args...)
	}
}

// 简版执行用时
//
// 用法：defer TimeCost("test")()
func TimeCost(tag string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%v 执行用时 %+v\n", tag, time.Since(start))
	}
}
