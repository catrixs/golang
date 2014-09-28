package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"os"
)

func fileSync(input chan string, out chan string) {
    f, err := os.Create("/data1/chenfei3/data/a.txt")
    if err != nil {
        panic(err)
    }

    w := bufio.NewWriter(f)

    for {
        s := <-input
        if s == "EXIT" {
            w.Flush()
            out <- "DONE"
        } else {
            w.WriteString(s)
            w.WriteString("\n")
        }
    }
}

func testSync(input chan string, out chan string) {
    f, err := os.Open("/data1/target")
    if err != nil {
        panic(err)
    }

    r := bufio.NewReader(f)
    for {
        s := <-input
        if s == "EXIT" {
            out <- "DONE"
        } else {
            line, _, err := r.ReadLine()
            if err != nil {
                panic(err)
            }

            if (s != string(line)) {
                fmt.Println("expect:", string(line), ", actual:", s)
            } else {
                //fmt.Println("[ok] actual:", s)
            }
        }
    }
}

func brickSync(input chan string) <-chan string {
	out := make(chan string)
	go func() {
		testSync(input, out)
        //fileSync(input, out)
	}()
	return out
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}

	var service = os.Args[1]
	fmt.Println("Connecting to server at ", service)
	conn, err := net.Dial("udp", service)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to server at ", service)
	defer conn.Close()

	input := make(chan string)
	c := brickSync(input)

	fmt.Println("Begin to transfer brick")
	req := 1
	var buf []byte = make([]byte, 200+1)
    receiveSum := 0
    start := time.Now()
	for {
		_, err := conn.Write([]byte(fmt.Sprintf("%d", req)))
		if err != nil {
			panic(err)
		}

		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
        receiveSum += n

		s := string(buf[0:n])
		if s == "EOF" {
			input <- "EXIT"
			<-c
			fmt.Println("Transfer OK")
			break
		} else {
            input <- s
        }

		req++
	}
    elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("elapsed = ", elapsed, " us")
    fmt.Println("bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}
