/*
Copyright © 2021 SF Chong <soonfook11@hotmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt <input>",
	Short: "Encrypt a string",
	Long:  `Encrypt a string. This will return an encrypted string and key to decrypt.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		encryptedText, key := encrypt(args[0])
		fmt.Println("Encrypted Text:")
		fmt.Println(encryptedText)
		fmt.Println("Key to decrypt:")
		fmt.Println(key)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
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
