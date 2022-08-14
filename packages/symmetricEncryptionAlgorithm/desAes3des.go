package symmetricEncryptionAlgorithm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func StudyAES(str, check string, key []byte) {
	de, _ := DesEn([]byte(str), key, check)
	fmt.Printf("%v 加密：%x \n", check, de)
	de, _ = DesDe(de, key, check)
	fmt.Println(check, "解密：", string(de))

	d, _ := DesEnStr(str, check, key)
	fmt.Println(check, "加密：", d)
	ds, _ := DesDeStr(d, check, key)
	fmt.Println(check, "解密：", string(ds))
}

// DES加密字节数组，返回字节数组
func DesEn(enBytes, key []byte, check string) ([]byte, error) {
	var block cipher.Block
	switch check {
	case "des":
		block, _ = des.NewCipher(key)
		pkcsBytes := PKCS5Padding(enBytes, block.BlockSize())
		blockMode := cipher.NewCBCEncrypter(block, key)
		pkcsArr := make([]byte, len(pkcsBytes))
		blockMode.CryptBlocks(pkcsArr, pkcsBytes)
		return pkcsArr, nil
	case "aes":
		block, _ = aes.NewCipher(key)

	case "3des":
		block, _ = des.NewTripleDESCipher(key)
	}
	return nil, nil
}

// DES加密字符串，返回字符串
func DesEnStr(enText, check string, key []byte) (string, error) {
	desDe, err := DesEn([]byte(enText), key, check)
	if err != nil {
		fmt.Println("2:", err)
	}
	base64str := base64.StdEncoding.EncodeToString(desDe)
	return base64str, nil
}

// DES解密字节数组，返回字节数组
func DesDe(deBtytes, key []byte, check string) ([]byte, error) {
	var block cipher.Block
	switch check {
	case "des":
		block, _ = des.NewCipher(key)
		blockMode := cipher.NewCBCDecrypter(block, key)
		deText := make([]byte, len(deBtytes))
		blockMode.CryptBlocks(deText, deBtytes)
		deText = PKCS5UnPadding(deText)
		return deText, nil
	case "aes":
		block, _ = aes.NewCipher(key)
	case "3des":
		block, _ = des.NewTripleDESCipher(key)
	}
	return nil, nil
}

// DES解密字符串，返回字符串
func DesDeStr(desDe, check string, key []byte) ([]byte, error) {
	base64Str, err := base64.StdEncoding.DecodeString(desDe)
	if err != nil {
		fmt.Println("4:", err)
	}
	des, err := DesDe(base64Str, key, check)
	if err != nil {
		fmt.Println("5:", err)
	}
	return des, nil
}

// 尾部填充
func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtest := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtest...) //解压压缩切片padtest
}
func PKCS5UnPadding(test []byte) []byte {
	leng := len(test)
	//去掉最后一个字节Unpadding次
	unpadding := int(test[leng-1])
	return test[:(leng - unpadding)]
}
