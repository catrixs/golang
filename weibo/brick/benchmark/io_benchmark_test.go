package benchmark

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"time"
	//"bytes"
	"net/textproto"
)

func rotate(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
$ go test -benchmem -bench ReadLargeFile

BenchmarkReadLargeFile	receive =  1063164745  Bytes
elapsed =  2.174487509s
bandwidth =  466.27672369632364  MB/s
       1	2174919710 ns/op	    4840 B/op	      18 allocs/op
ok  	github.com/feiyang21687/golang/weibo/brick/benchmark	2.191s
*/
func BenchmarkReadLargeFile(b *testing.B) {
	f, err := os.Open("/data1/t.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lineNo int = 1
	start := time.Now()
	receiveSum := 0

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		n := len(line)
		receiveSum += n

		s := n / 3
		if s != 0 {
			ss := line[0 : n-s]
			copy(line[s:len(ss)], line[s+s:n])
			rotate(ss)
		} else {
			rotate(line)
		}

		lineNo++
	}
	elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("elapsed = ", elapsed)
    fmt.Println("bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}

/*
$ go test -benchmem -bench IoCopy

BenchmarkIoCopy	client start reciving.....
receive =  1063164745  Bytes
receive elapsed =  3.065908458s
receive bandwidth =  330.7055397461251  MB/s
       1	3066049413 ns/op	1139665176 B/op	10275724 allocs/op
ok  	github.com/feiyang21687/golang/weibo/brick/benchmark	3.074s
*/
func BenchmarkIoCopy(b *testing.B) {
	laddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:1200")
	l, _ := net.ListenTCP("tcp", laddr)
	defer l.Close()

	BUF_SIZE := 100 * 1024 * 1024
	go func() {
		conn, err := l.AcceptTCP()
		if err != nil {panic(err)}
		defer conn.Close()

		conn.SetWriteBuffer(BUF_SIZE)

		f, err := os.Open("/data1/t.txt")
		if err != nil { panic(err) }
		io.Copy(conn, bufio.NewReader(f))
	}()

	start := time.Now()
	receiveSum := 0
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1200")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	conn.SetReadBuffer(BUF_SIZE)
	fmt.Println("client start reciving.....")
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
	    // read one line (ended with \n or \r\n)
	    line, err := tp.ReadLine()
	    if err != nil {break}
	    receiveSum += len(line)
	    // do something with data here, concat, handle and etc... 
	}
	
	elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("receive elapsed = ", elapsed)
    fmt.Println("receive bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}


func BenchmarkPingPong(b *testing.B) {
	laddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:1200")
	l, _ := net.ListenTCP("tcp", laddr)
	defer l.Close()

	BUF_SIZE := 100 * 1024 * 1024
	go func() {
		conn, err := l.AcceptTCP()
		if err != nil {panic(err)}
		defer conn.Close()

		conn.SetWriteBuffer(BUF_SIZE)

		f, err := os.Open("/data1/t.txt")
		if err != nil { panic(err) }

		reqQ := make([]byte, 256)
		for {
			conn.Read(reqQ)
			n, err := io.CopyN(conn, bufio.NewReader(f), 200)
			if n == 0 || err != nil {break}
		}	
	}()

	start := time.Now()
	receiveSum := 0
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1200")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	conn.SetReadBuffer(BUF_SIZE)

	go func () {
		for {
			_, err := conn.Write([]byte("1"))
			if err != nil {return}
		}
	}()

	fmt.Println("client start reciving.....")
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
	    // read one line (ended with \n or \r\n)
	    line, err := tp.ReadLine()
	    if err != nil {break}
	    receiveSum += len(line)
	    // do something with data here, concat, handle and etc... 
	}
	
	elapsed := time.Since(start)
    fmt.Println("receive = ", receiveSum, " Bytes")
    fmt.Println("receive elapsed = ", elapsed)
    fmt.Println("receive bandwidth = ", (float64(receiveSum))/elapsed.Seconds()/1024/1024, " MB/s")
}

