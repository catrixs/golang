package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func rotate(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func brickWorker() <-chan string {
	c := make(chan string)
	go func() {
		f, err := os.Open("/data1/t.txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		var lineNo int = 1
		reader := bufio.NewReader(f)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					c <- "EOF"
					return
				} else {
					panic(err)
				}
			}

			n := len(line)
			s := n / 3
			if s != 0 {
				ss := line[0 : n-s]
				copy(line[s:len(ss)], line[s+s:n])
				rotate(ss)
				ret := fmt.Sprintf("%d%s", lineNo, string(ss))
				c <- ret
			} else {
				rotate(line)
				c <- fmt.Sprintf("%d%s", lineNo, string(line))
			}
			lineNo++
		}
	}()
	return c
}

func main() {
	port := "127.0.0.1:1200"

	udpAddress, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var buf []byte = make([]byte, 1500)
	c := brickWorker()
	for {
		n, address, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		if address != nil {
			//fmt.Println("got message from ", address, " with req = ", string(buf))
			if n > 0 {
				//fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
				out := []byte(<-c)
				//fmt.Println("write:%s", string(out), " byte=", out)
				conn.WriteToUDP(out, address)
			}
		}
	}
}
