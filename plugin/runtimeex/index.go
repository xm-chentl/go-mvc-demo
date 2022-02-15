package runtimeex

import (
	"fmt"
	"runtime"
)

func ShowGC() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("此时占用内存: %vKB, GC次数: %v\n", m.Alloc/1024, m.NumGC)
}
