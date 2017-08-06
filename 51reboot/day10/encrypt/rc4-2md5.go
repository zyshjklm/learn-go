package main

import (
	"crypto/md5"
	"crypto/rc4"
	"log"
)

func main() {
	key := "123456"
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil {
		log.Fatal(err)
	}
	buf := []byte("hello")
	cipher.XORKeyStream(buf, buf)
	log.Printf(string(buf))

	{
		cipher, err := rc4.NewCipher([]byte(md5sum[:]))
		if err != nil {
			log.Fatal(err)
		}
		cipher.XORKeyStream(buf, buf)
		log.Print(string(buf))
	}

}
