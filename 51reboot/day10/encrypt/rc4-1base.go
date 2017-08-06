package main

import (
	"crypto/rc4"
	"log"
)

func main() {
	key := "123456"
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	buf := []byte("hello")
	cipher.XORKeyStream(buf, buf)
	log.Printf("crypt#:%s#\n", string(buf))

	{
		cipher, err := rc4.NewCipher([]byte(key))
		if err != nil {
			log.Fatal(err)
		}
		cipher.XORKeyStream(buf, buf)
		log.Print(string(buf))
	}

}
