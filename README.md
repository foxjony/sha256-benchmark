# Golang SHA256 Benchmark CPU Test
- This is a very simple and reliable performance test of your processor that shows its speed compared to other processors.
- This test of the performance of your device (CPU and RAM test) differs from others in that the same simple program is used in the same way on very different devices. You can test both powerful computers and laptops, as well as tablets, mobile phones and TV-boxes.
- This is a good example to start learning about a very fast programming language Golang.
- To do this, you need to compile and run the file: main.go
- This program does not use third-party libraries, it uses only three standart libraries of the Golang compiler.
- Algorithm: the program generates 1,000,000 hash sums by adding them to the table in RAM. This is repeated 9 times and each time the time spent on the action is displayed. The shorter the time, the faster the device. Regardless of the number of cores on your processor, only one core is tested.
- I collect test results in an excel table. And I will be grateful to all people who repeat this test and send me tests of their processors in the comments. Also show this test to your friends.
- I noticed one peculiarity. Testing different devices with my test, I saw different deltas of discrepancy in the 9 lines of the test. On some OS, the numbers are the same (or almost the same), and on others with a large discrepancy. This indicates an interruption of the OS to other processes. I know that Linux is not a real-time OS, but still some OSes work without heavy interruptions. This can be used to detect some type of viruses embedded in the OS.

## Table of results: Golang SHA256 Benchmark CPU Test.xlsx
![Table of results](https://github.com/foxjony/sha256-benchmark/blob/main/Table.png)

## Source code: main.go 
![Source code](https://github.com/foxjony/sha256-benchmark/blob/main/Source.png)

## The result of the program
![The result of the program](https://github.com/foxjony/sha256-benchmark/blob/main/Result.png)

## Instruction
Go in 100 Seconds
- https://www.youtube.com/watch?v=446E-r0rXHI

Хауди Хо: Как выучить GO? Самый аху#### способ!
- https://www.youtube.com/watch?v=sP-wNhqgp9U

Гоша Дударь: Изучение Golang (Создание веб сайта) / #1 – Введение в язык Go
- https://www.youtube.com/watch?v=OcWMQPPOq-0

Golang download and install
- https://go.dev/doc/install

Golang Online Test Programm
- https://go.dev/play/p/I4NnsnNQ91-

For Windows build.bat
```
go build main.go
main
pause
```

For Linux and Mac Termanal
```
apt-get install golang -y
go vetsion
mkdir home/go && cd home/go
wget https://raw.githubusercontent.com/foxjony/sha256-benchmark/main/main.go
go build main.go
./main
```

- Note: The first test usually takes longer to generate, so I don't count it.

## Debian Linux image for Android TV boxes with Amlogic
- https://github.com/devmfc/debian-on-amlogic

## Wokwi
Algorithm SHA256 fof "ESP32", "ESP32-C3", "Raspberry Pi Pico" and "STM32" in Online Web Simulator Wokwi:
- https://wokwi.com/projects/398636708003334145 - SHA-256 for ESP32
- https://wokwi.com/projects/399641258558275585 - SHA-256 for ESP32-C3
- https://wokwi.com/projects/399642127436776449 - SHA-256 for Raspberry Pi Pico (RP2040)
- https://wokwi.com/projects/399643889913932801 - SHA-256 for STM32

## Youtube
Youtube "Golang SHA256 Benchmark CPU Test"
- (EN) https://youtu.be/zDBve2iOUuk
- (UA) https://youtu.be/S-YLmRWMlq0
- (RU) https://youtu.be/jBGwUaxNwfA
