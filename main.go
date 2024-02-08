package main

// Golang SHA256 Benchmark CPU Test

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nGolang SHA256 Benchmark CPU Test\n")
	var timeAverage float64
	h := []byte("Benchmark")
	hash := sha256.Sum256(h)
	for j := 1; j < 10; j++ {
		start := time.Now()
		m := [][32]byte{}
		for i := 0; i < 1000000; i++ {
			hash = sha256.Sum256(h)
			m = append(m, hash)
		}
		m = nil
		end := time.Now()
		var duration time.Duration = end.Sub(start)
		var di64 = float64(duration)
		if (j > 1) {timeAverage += di64/1000000}
		fmt.Printf("%d time: %.2f ms\n", j, di64/1000000)
	}
	fmt.Printf("\nTime Average: %.2f ms\n\n", timeAverage/8)
	var a string
	fmt.Scan(&a)
}
