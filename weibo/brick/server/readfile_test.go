package main

import (
	//"bufio"
	"fmt"
	//"io"
	//"os"
	//"time"
	"testing"
	//"net"
	//"syscall"
	"encoding/binary"
	//"bytes"
)

/*
func TestReadLargeFile(t *testing.T) {
	f, err := os.Open("/data1/t.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lineNo int = 1
	start := time.Now()
	receiveSum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		n := len(line)
		receiveSum += n
		
		fmt.Sprintf("%d%s", lineNo, string(line))
		lineNo++
	}
	elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("elapsed = ", elapsed)
    fmt.Println("bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}
*/

/*
func TestReadLargeFileWithChannel(t *testing.T) {
	c := make(chan *string)

	start := time.Now()
	receiveSum := 0	

	go func() {
		f, err := os.Open("/data1/t.txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		var lineNo int = 1
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			s := fmt.Sprintf("%d%s", lineNo, string(line))
			c <- &s
			lineNo++
		}
		s := "EOF"
		c <- &s
	}()

	for {
		line := <- c
		if *line == "EOF" {break}
		receiveSum += len(*line)
	}

	elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("elapsed = ", elapsed)
    fmt.Println("bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}
*/

/*
func TestUdpServer(t *testing.T) {
	c := make(chan string)
	go func() {
		port := "0.0.0.0:1200"
		udpAddress, err := net.ResolveUDPAddr("udp4", port)
		if err != nil {
			panic(err)
		}

		conn, err := net.ListenUDP("udp", udpAddress)
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		conn.SetWriteBuffer(100*1024*1024)

		c <- "start"

		var buf []byte = make([]byte, 1500)
		for {
			n, address, err := conn.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}
			if address != nil {
				if n > 0 {
					conn.WriteToUDP([]byte("1qdY!#eow>2SCQu.:h]VR049EIq?-j0zv6-vJs9SK1ba2&c8-fMqwQ0x.seD}V45nK~M8=vfdafdasfdasfdasfsafdsafdsa"), address)
					//conn.WriteToUDP([]byte("1"), address)
				}
			}
		}
	}()

	<- c

	go func () {
		var service = "10.209.1.98:1200"
		fmt.Println("Connecting to server at ", service)
		udpAddress, err := net.ResolveUDPAddr("udp4", service)
		if err != nil { panic(err) }
		conn, err := net.DialUDP("udp", nil, udpAddress)
		if err != nil {
			panic(err)
		}

		fmt.Println("Connected to server at ", service)
		defer conn.Close()
		conn.SetReadBuffer(100*1024*1024)

		go func() {
			var buf []byte = make([]byte, 200+1)
			receiveSum := 0
	    	start := time.Now()

	    	for {	
				n, err := conn.Read(buf)
				if err != nil {
					break
				}
				//fmt.Println("receive = ", n)
		        receiveSum += n
	    	}

	    	elapsed := time.Since(start)
		    fmt.Println("receive = ", receiveSum, " Bytes")
		    fmt.Println("elapsed = ", elapsed)
		    fmt.Println("bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
		    c <- "EXIT"
		}()

		req := 1
		for i := 0; i < 1000000; i++{
			_, err := conn.Write([]byte(fmt.Sprintf("%d", req)))
			if err != nil {
				panic(err)
			}	
			req++
		}
		fmt.Println("write done")
	}()

	<- c
}
*/

func TestBinaryEncoding(t *testing.T) {
	b := []byte{0, 0, 0, 0, 0, 0}
	bs := b[0:4]
	//buf := bytes.NewBuffer(b)

	var i int = 31
	var j int = 0
	binary.BigEndian.PutUint32(bs, uint32(i))
	j = int(binary.BigEndian.Uint32(bs))
	fmt.Println(j)
}
