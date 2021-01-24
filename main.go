package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func main() {
	input := "password"
	encryptedText, key := encrypt(input)
	fmt.Println("Encrypted Text:")
	fmt.Println(encryptedText)
	fmt.Println("Key to decrypt:")
	fmt.Println(key)

	// input := "0a325301db062009ee7c7d1f08f686ec560d6f6919648cceae10aed4d0a3701f054ee275"
	// key := "bf3f38922cdabb42e8808fa274aceb1dd1949717875168778f11b5849332eb0c"
	// decryptedText := decrypt(input, key)
	// fmt.Println(decryptedText)
}

func encrypt(inputStr string) (string, string) {
	input := []byte(inputStr)

	// Generate random 32 byte key
	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		log.Fatal(err)
	}

	// Encode key bytes to string
	key := hex.EncodeToString(keyBytes)

	// Create cipher block from key
	cipherBlock, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Create gcm from cipher block
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatal(err)
	}

	// Create nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// Encrypt data, add nonce as prefix (first nonce argument) to the encrypted data
	encryptedByte := gcm.Seal(nonce, nonce, input, nil)

	// Convert encrypted byte to base 16
	encryptedText := fmt.Sprintf("%x", encryptedByte)

	return encryptedText, key
}

func decrypt(inputStr string, keyStr string) string {
	encryptedByte, _ := hex.DecodeString(inputStr)
	keyByte, _ := hex.DecodeString(keyStr)

	// Create cipher block from key
	cipherBlock, err := aes.NewCipher(keyByte)
	if err != nil {
		log.Fatal(err)
	}

	// Create gcm from cipher block
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatal(err)
	}

	// Get nonce size
	nonceSize := gcm.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, encryptedByte := encryptedByte[:nonceSize], encryptedByte[nonceSize:]

	// Decrypt data
	decryptedByte, err := gcm.Open(nil, nonce, encryptedByte, nil)

	return string(decryptedByte)
}
