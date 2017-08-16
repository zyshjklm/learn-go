package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/jungle85gopy/learn-go/51reboot/cryptoSocks5-V2/mycrypto"
)

var (
	socks5 = flag.String("s", ":8022", "port for socks5")
)

const key = "123456"

func mustReadByte(r *bytes.Buffer) (b byte) {
	// log.Println("-- start mustReadByte")
	var err error
	for {
		if b, err = r.ReadByte(); err == io.EOF {
			time.Sleep(time.Millisecond * 10)
			continue
		} else if err != nil {
			log.Printf("mustReadByte err:%s\n", err.Error())
			panic(err)
		}
		break
	}
	return
}

// 握手。进行授权认证
func handshake(r *bytes.Buffer, wr io.Writer) (err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()

	// log.Println("start handshake.")
	// version
	version := mustReadByte(r)
	// log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	// nmethods
	nmethods := mustReadByte(r)
	// log.Printf("nmethods: %d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	// log.Printf("methods type:%v", buf)
	// response
	resp := []byte{5, 0}
	n, err := wr.Write(resp[:])
	log.Printf("handshake write back [%v] with %d byte", n == len(resp), n)
	return
}

// 获取需要代理的域名
func readAddr(r *bytes.Buffer) (addrPort string, err error) {
	defer func() {
		e := recover() // interface{}
		if e != nil {
			err = e.(error)
		}
	}()

	funcName := "readAddr():"
	// log.Printf("run %s ...", funcName)
	version := mustReadByte(r)
	// log.Printf("%s version:%d", funcName, version)
	if version != 5 {
		return "", errors.New(funcName + " bad version")
	}
	cmd := mustReadByte(r)
	// log.Printf("%s cmd:%d", funcName, cmd)
	if cmd != 1 {
		return "", errors.New(funcName + " bad cmd")
	}
	// skip rsv
	mustReadByte(r)
	// addr type
	addrType := mustReadByte(r)
	// log.Printf("%s addrType:%d", funcName, addrType)
	if addrType != 3 {
		return "", errors.New(funcName + " bad addr type")
	}
	// domain name. 域名是变长的，读取一个字节的数据，代表后面紧跟着的域名的长度
	addrLen := mustReadByte(r)
	addr := make([]byte, addrLen)
	io.ReadFull(r, addr)
	// log.Printf("%s domain len:%d", funcName, addrLen)
	// log.Printf("%s domain:%s", funcName, addr)
	// port
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	addrPort = fmt.Sprintf("%s:%d", addr, port)
	return
}

func socks5Auth(r *bytes.Buffer, wr io.Writer) (addr string, err error) {
	if err = handshake(r, wr); err != nil {
		log.Println("handshake return err...")
		return "", err
	}
	if addr, err = readAddr(r); err != nil {
		log.Println("readAddr return err...")
		return "", err
	}
	// resp 响应客户端连接成功
	resp := []byte{0x05, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	wr.Write(resp)
	return
}

func handleConn(conn net.Conn) {
	log.Println("start handle ...")
	defer conn.Close()

	wg := new(sync.WaitGroup)
	wg.Add(4)

	// 对client端请求的数据进行解密
	decryptBuf := new(bytes.Buffer)
	decryptWr := mycrypto.NewCryptoWriter(decryptBuf, key)
	// 对server端返回的数据进行加密
	encryptBuf := new(bytes.Buffer)
	encryptWr := mycrypto.NewCryptoWriter(encryptBuf, key)

	// from conn to decrypt buf
	go func() {
		defer wg.Done()
		io.Copy(decryptWr, conn)
		log.Println("++++ go 1 over!")
	}()
	log.Println("go start decrypt...")

	// from encrypt buf to conn
	go func() {
		defer wg.Done()
		for {
			io.Copy(conn, encryptBuf)
			// log.Printf("~~~~ go routine 2 Copy over with %d bytes!", n)
			time.Sleep(time.Millisecond * 90)
		}
		log.Println("++++ go 2 over!")
	}()
	log.Println("go start encrypt...")

	// 说明：
	//    从解密后的buffer读数据进行握手和转发，是直接使用的*bytes.Buffer，
	//    从加密后的buffer读数据写回client端 ，是直接使用的*bytes.Buffer
	//    这里实验过，如果使用如下2句，构造2次Reader，得到的version是242
	// decryptRd := mycrypto.NewCryptoReader(decryptBuf, key)
	// r := bufio.NewReader(conn)

	// read from decryptBuf, write to encryptWr
	addr, err := socks5Auth(decryptBuf, encryptWr)
	// addr, err := socks5Auth(r, encryptWr)
	if err != nil {
		log.Printf("get addr err:%s\n", err.Error())
		return
	}
	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("dial remote err:%s", err.Error())
		conn.Close()
		return
	}

	// from decrypt to remote
	go func() {
		defer wg.Done()
		for {
			io.Copy(remote, decryptBuf)
			time.Sleep(time.Millisecond * 90)
		}
		log.Println("++++ go 3 over!")
	}()
	// from remote to encrypt
	go func() {
		defer wg.Done()
		for {
			io.Copy(encryptWr, remote)
			time.Sleep(time.Millisecond * 90)
		}
		log.Println("++++ go 4 over!")
	}()

	wg.Wait()
	log.Printf("shut of %s", conn.RemoteAddr().String())
	conn.Close()
	remote.Close()
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", *socks5)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("new connection from %s\n\n", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
