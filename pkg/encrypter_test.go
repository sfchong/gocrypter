package encrypter

import "testing"

func TestEncrypt(t *testing.T) {
	input := "test"

	_, _, err := Encrypt(input)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDecrypt(t *testing.T) {
	input := "test"
	encrypted := "0bab69a3011319f9c10a34487e9e83acc62089a120b79a39617c18b8985fd972"
	key := "76f966ae407e08cc29f8c6cd66c95e0f9297c75bed913b333d75cdd7270334bc"

	decrypted, err := Decrypt(encrypted, key)

	if err != nil {
		t.Fatal(err)
	}

	if input != decrypted {
		t.Fatal("Decrypted result is not matched with encrypted input!")
	}
}
