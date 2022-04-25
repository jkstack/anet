package anet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var key = []byte{
	0x76, 0x64, 0x71, 0x45,
	0x51, 0x64, 0x74, 0x32,
	0x43, 0x52, 0x44, 0x61,
	0x57, 0x41, 0x4c, 0x42,
	0x50, 0x44, 0x4c, 0x70,
	0x48, 0x53, 0x41, 0x67,
	0x45, 0x62, 0x52, 0x6a,
	0x53, 0x7a, 0x4d, 0x6b,
}

var iv = []byte{
	0x57, 0x59, 0x34, 0x4a,
	0x50, 0x35, 0x69, 0x41,
	0x4c, 0x77, 0x77, 0x79,
	0x47, 0x35, 0x36, 0x4b,
}

func pkcs7padding(data []byte, block int) []byte {
	padding := block - len(data)%block
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7unpadding(data []byte) []byte {
	padding := data[len(data)-1]
	return data[:len(data)-int(padding)]
}

func Encrypt(data []byte) ([]byte, error) {
	enc, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = pkcs7padding(data, enc.BlockSize())
	block := cipher.NewCBCEncrypter(enc, iv)
	ret := make([]byte, len(data))
	block.CryptBlocks(ret, data)
	return ret, nil
}

func Decrypt(data []byte) ([]byte, error) {
	dec, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	block := cipher.NewCBCDecrypter(dec, iv)
	raw := make([]byte, len(data))
	block.CryptBlocks(raw, data)
	return pkcs7unpadding(raw), nil
}
