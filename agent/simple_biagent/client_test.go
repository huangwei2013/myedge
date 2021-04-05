package main


import (
	"testing"
)

func Benchmark_RunAgent(b *testing.B){
	b.N = 1000
	for i := 0; i < b.N; i++ {
		b.ReportAllocs() // 这里可以直接调用 ReportAllocs 方法，就省去了再命令行中输入 -benchmem ，用于查看内存分配的大小和次数
		run()
	}

}