package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(inputStr string) (string, string, error) {
	input := []byte(inputStr)

	// Generate random 32 byte key
	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		return "", "", err
	}

	// Encode key bytes to string
	key := hex.EncodeToString(keyBytes)

	// Create cipher block from key
	cipherBlock, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", "", err
	}

	// Create gcm from cipher block
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", "", err
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

	return encryptedText, key, nil
}

func Decrypt(inputStr string, keyStr string) (string, error) {
	encryptedByte, err := hex.DecodeString(inputStr)
	if err != nil {
		return "", err
	}

	keyByte, err := hex.DecodeString(keyStr)
	if err != nil {
		return "", err
	}

	// Create cipher block from key
	cipherBlock, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	// Create gcm from cipher block
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	// Get nonce size
	nonceSize := gcm.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, encryptedByte := encryptedByte[:nonceSize], encryptedByte[nonceSize:]

	// Decrypt data
	decryptedByte, err := gcm.Open(nil, nonce, encryptedByte, nil)

	return string(decryptedByte), nil
}
