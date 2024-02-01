package main

// Golang SHA256 Benchmark CPU Test

import (
	"fmt"
	"time"
	"crypto/sha256"
)

func benchmark(n int) {
	h := []byte("Benchmark")
	hash := sha256.Sum256(h)
	//fmt.Printf("sha256: %x\n\n", hash)
	for j := 1; j<10; j++ {
		start := time.Now()
		m := [][32]byte{}
		for i := 0; i<n; i++ {
			hash = sha256.Sum256(h)
			m = append(m, hash)
		}
		m = nil
		end := time.Now()
		fmt.Println(j, " time:", end.Sub(start))
	}
}

func main() {
	fmt.Println("\nGolang SHA256 Benchmark CPU Test 100000")
	benchmark(100000)
	fmt.Println("\nGolang SHA256 Benchmark CPU Test 1000000")
	benchmark(1000000)
	fmt.Println("\nGolang SHA256 Benchmark CPU Test 10000000")
	benchmark(10000000)
	//fmt.Println("\nGolang SHA256 Benchmark CPU Test 100000000")
	//benchmark(100000000)	// If RAM >8GB
}
