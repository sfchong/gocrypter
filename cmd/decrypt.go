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
	"encoding/hex"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt <input>",
	Short: "Decrypt a string with a key",
	Long:  `Decrypt a string with a key. This will return a decrypted string.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		decryptedText := decrypt(args[0], key)
		fmt.Println(decryptedText)
	},
}

var key string

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringVarP(&key, "key", "k", "", "key to decrypt")
	decryptCmd.MarkFlagRequired("key")
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
