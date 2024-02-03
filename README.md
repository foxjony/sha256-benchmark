# Golang SHA256 Benchmark CPU Test
- This is a very simple and reliable performance test of your processor that shows its speed compared to other processors.
- This test of the performance of your device (CPU and RAM test) differs from others in that the same simple program is used in the same way on very different devices. You can test both powerful computers and laptops, as well as tablets, mobile phones and TV-boxes.
- This is a good example to start learning about a very fast programming language Golang.
- To do this, you need to compile and run the file: main.go
- This program does not use third-party libraries, it uses only three standart libraries of the Golang compiler.
- Algorithm: the program generates 1,000,000 hash sums by adding them to the table in RAM. This is repeated 9 times and each time the time spent on the action is displayed. The shorter the time, the faster the device. Regardless of the number of cores on your processor, only one core is tested.
- I collect test results in an excel table. And I will be grateful to all people who repeat this test and send me tests of their processors in the comments. Also show this test to your friends.

![Table Сollect resultsl](https://github.com/foxjony/sha256-benchmark/blob/main/Table.png)

![Source](https://github.com/foxjony/sha256-benchmark/blob/main/Source.png)

![Table Сollect resultsl](https://github.com/foxjony/sha256-benchmark/blob/main/Result.png)

Go in 100 Seconds
- https://www.youtube.com/watch?v=446E-r0rXHI

Golang download and install
- https://go.dev/doc/install

For Windows build.bat
```
go build main.go
main
pause
```

For Linux and Mac Termanal
```
go build main.go
./main
```

The result of the program
```
System:   Windows 10
Name:     PC
CPU:      AMD Ryzen5 3600 6-Core
RAM:      DDR4 16GB

Average time: 	9.5ms	    (n = 100 000)
Average time:  	 85ms	  (n = 1 000 000)
Average time: 	850ms	 (n = 10 000 000)
Average time: 	8.33s	(n = 100 000 000)
```

```
Note: The first test usually takes longer to generate, so I don't count it.

Golang SHA256 Benchmark CPU Test 100000
1  time: 11.5231ms
2  time: 9.8688ms
3  time: 9.3218ms
4  time: 9.8494ms
5  time: 9.3222ms
6  time: 9.8709ms
7  time: 7.0961ms
8  time: 5.5373ms
9  time: 9.9796ms

Golang SHA256 Benchmark CPU Test 1000000
1  time: 100.2465ms
2  time: 90.4809ms
3  time: 90.4033ms
4  time: 81.5508ms
5  time: 90.7188ms
6  time: 81.236ms
7  time: 90.0365ms
8  time: 90.6217ms
9  time: 81.1459ms

Golang SHA256 Benchmark CPU Test 10000000
1  time: 990.4237ms
2  time: 882.8575ms
3  time: 856.1492ms
4  time: 855.4965ms
5  time: 866.774ms
6  time: 869.1547ms
7  time: 867.7644ms
8  time: 870.4721ms
9  time: 867.1598ms

Golang SHA256 Benchmark CPU Test 100000000
1  time: 9.3205342s
2  time: 8.6575299s
3  time: 8.3266816s
4  time: 8.3360652s
5  time: 8.361239s
6  time: 8.3720439s
7  time: 8.3616938s
8  time: 8.3711706s
9  time: 8.3740316s
```

Golang SHA256 Benchmark CPU Test
- https://www.youtube.com/watch?v=La8AW8Ea8oo
