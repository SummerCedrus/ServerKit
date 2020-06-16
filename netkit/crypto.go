package netkit

import (
	"encoding/hex"
	"crypto/des"
	"crypto/cipher"
	"bytes"
	"fmt"
)

const(
	CIPHER_CODE ="MFwwDQYJ"
)

//CBC加密
func EncryptDES_CBC(origData []byte) ([]byte,error) {
	block, err := des.NewCipher([]byte(CIPHER_CODE))
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	iv := createIv([]byte(CIPHER_CODE))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//CBC解密
func DecryptDES_CBC(cryptedData []byte) ([]byte,error) {
	block, err := des.NewCipher([]byte(CIPHER_CODE))
	if err != nil {
		return nil, err
	}
	iv := createIv([]byte(CIPHER_CODE))

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cryptedData))
	blockMode.CryptBlocks(origData, cryptedData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//ECB加密
func EncryptDES_ECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

//ECB解密
func DecryptDES_ECB(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

//明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
//
func createIv(key []byte) []byte {
	return key
}
