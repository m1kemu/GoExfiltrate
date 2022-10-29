package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io/ioutil"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

func DecryptAES(ciphertext, key []byte) []byte {
	c, _ := aes.NewCipher(key)

	IV := []byte("1234567812345678")

	stream := cipher.NewCTR(c, IV)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

func DecryptXOR(ciphertext, key []byte) []byte {
	plaintext := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = ciphertext[i] ^ key[i%len(key)]
	}

	return plaintext
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func RandStringBytes(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = ALPHABET[rand.Intn(len(ALPHABET))]
	}

	return string(b)
}

func AESKeyPadding(byte_array []byte) []byte {
	if len(byte_array) == 32 {
		return byte_array
	} else {
		if len(byte_array) < 32 {
			bytes_to_add := 32 - len(byte_array)

			for i := 0; i < bytes_to_add; i++ {
				byte_array = append(byte_array, byte('A'))
			}

			return byte_array
		} else {
			byte_array = byte_array[0:32]
		}

		return byte_array
	}
}

func PlaintextB64Encode(shellcode_file string) string {
	log.Info("Not performing any encryption. Encoding the file.")

	plaintext, _ := ioutil.ReadFile(shellcode_file)

	plaintext_b64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	log.Info("Generated encoded text: " + plaintext_b64)

	return plaintext_b64
}

func EncryptAES(shellcode_file string, key string) string {
	log.Info("Performing AES encryption on file " + shellcode_file + " using key: " + key)

	plaintext, _ := ioutil.ReadFile(shellcode_file)
	key_byte := []byte(key)

	c, err := aes.NewCipher(key_byte)
	if err != nil {
		log.Error(err)
	}

	IV := []byte("1234567812345678")

	stream := cipher.NewCTR(c, IV)
	stream.XORKeyStream(plaintext, plaintext)
	ciphertext := plaintext

	ciphertext_b64 := base64.StdEncoding.EncodeToString(ciphertext)

	log.Info("Generated encoded text: " + ciphertext_b64)

	return ciphertext_b64
}

func EncryptXOR(shellcode_file string, key string) string {
	log.Info("Performing XOR encryption on file " + shellcode_file + " using key: " + key)

	plaintext, _ := ioutil.ReadFile(shellcode_file)
	key_byte := []byte(key)

	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key_byte[i%len(key_byte)]
	}

	ciphertext_b64 := base64.StdEncoding.EncodeToString([]byte(ciphertext))

	log.Info("Generated ciphertext: " + ciphertext_b64)

	return ciphertext_b64
}

func CopyFile(src, dest string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
