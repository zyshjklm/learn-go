package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"
	"log"
)

// refer: http://localhost:6060/pkg/crypto/cipher/#example_NewCFBEncrypter
func main() {
	// defined length. 128 / 8 = 16 byte
	key := "123456"
	plaintext := []byte("hello golang is the orgin content!")

	md5sum := md5.Sum([]byte(key))
	block, err := aes.NewCipher(md5sum[:]) // aes-128 16Byte
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	// cipher [iv : ciphered]
	cv := make([]byte, aes.BlockSize+len(plaintext))
	iv := cv[:aes.BlockSize]
	ciphertext := cv[aes.BlockSize:]

	log.Printf("iv:%x\n", iv)
	io.ReadFull(rand.Reader, iv)
	log.Printf("iv:%x\n", iv)

	// 加密
	streamEn := cipher.NewCFBEncrypter(block, iv)
	log.Printf("orgin :%x\n", plaintext)
	streamEn.XORKeyStream(ciphertext, plaintext)
	log.Printf("cipher:%x\n", ciphertext)

	// 解密
	streamDe := cipher.NewCFBDecrypter(block, iv)
	streamDe.XORKeyStream(ciphertext, ciphertext)
	log.Printf("result:%x\n", ciphertext)
}
