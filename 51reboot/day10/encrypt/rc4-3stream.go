package main

import (
	"crypto/md5"
	"crypto/rc4"
	"flag"
	"io"
	"log"
	"os"
	"time"
)

var 	key = flag.String("k", "", "secret key")

func crypto(w io.Writer, r io.Reader, key string) {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err != nil || err == io.EOF {
			log.Print(err)
			break
		}
	
		cipher.XORKeyStream(buf[:n], buf[:n])
		w.Write(buf[:n])
		time.Sleep(time.Second)
	}
}

func main() {
	flag.Parse()
	crypto(os.Stdout, os.Stdin, *key)
}
