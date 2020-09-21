package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("SHIFT CIPHER")

	n := 127

	plaintext := flag.String("encrypt", "", "Plain text to be encrypted")
	ciphertext := flag.String("decrypt", "", "Cipher text to be decrypted")
	key := flag.Int("key", 1, "shift length")

	flag.Parse()

	if *key <= 0 {
		*key = rand.Intn(100)
		fmt.Println("Invalid Key Input, randomly chooses", *key)
		fmt.Println("==========")
	}
	if *plaintext != "" {
		fmt.Println("Plain Text:", *plaintext)
		fmt.Println("Key:", *key)
		fmt.Println("ENKRIPSI")
		ciphertext := enc(*plaintext, n, *key)
		fmt.Println("Cipher Text:", ciphertext)
		fmt.Println("==========")
	}
	if *ciphertext != "" {
		fmt.Println("Cipher Text:", *ciphertext)
		fmt.Println("Key:", *key)
		fmt.Println("DEKRIPSI")
		plaintext := dec(*ciphertext, n, *key)
		fmt.Println("Plain Text:", plaintext)
		fmt.Println("==========")
	}

}

func enc(plaintext string, n int, k int) string {
	var plainbyte = []byte(plaintext)
	var cb = make([]byte, len(plainbyte))
	for i := 0; i < len(plaintext); i++ {
		fmt.Printf("x[%d]: %d %s\n", i, int(plainbyte[i]), string(plainbyte[i]))
		cb[i] = (plainbyte[i] + byte(k)) % byte(n)
		if cb[i] < 32 {
			cb[i] += 32
		}
	}
	return string(cb)
}

func dec(ciphertext string, n int, k int) string {
	plainbyte := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		var temp int
		fmt.Printf("x[%d]: %d %s\n", i, int(ciphertext[i]), string(ciphertext[i]))
		temp = int(ciphertext[i]) - k
		if temp < 32 {
			temp -= 32
		}
		if temp < 0 {
			temp = (temp%n + n) % n
		} else {
			temp = (temp % n)
		}
		plainbyte[i] = byte(temp)
	}

	return string(plainbyte)
}
