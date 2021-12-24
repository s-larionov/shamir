package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(key, text []byte) (encrypted []byte, nonce []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	var gcm cipher.AEAD
	if gcm, err = cipher.NewGCM(block); err != nil {
		return
	}

	if nonce, err = RandBytes(gcm.NonceSize()); err != nil {
		return
	}

	encrypted = gcm.Seal(nil, nonce, text, nil)

	return
}

func EncryptSimple(key, text []byte) ([]byte, error) {
	encrypted, nonce, err := Encrypt(key, text)
	if err != nil {
		return nil, err
	}

	result := make([]byte, 0, len(nonce)+len(encrypted))
	result = append(result, nonce...)
	result = append(result, encrypted...)

	return result, nil
}

func RandBytes(size int) ([]byte, error) {
	value := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, value); err != nil {
		return nil, err
	}

	return value, nil
}
