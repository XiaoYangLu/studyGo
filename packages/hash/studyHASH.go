package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	hash "hash"
)

/*
 * hash算法可以将任意长度的二进制值(明文)映射为较短的固定长度的二进制值(hash值)，
 * hash值又称为数字指纹、数字摘要
 */
func StudyHASH(str string, hashType string) string {
	var hash hash.Hash
	switch hashType {
	case "md5":
		hash = md5.New()
	case "sha256":
		hash = sha256.New()
	case "sha512":
		sha512.New()
	case "sha1":
		sha1.New()
	}
	hash.Write([]byte(str))
	hashByte := hash.Sum([]byte(nil))
	hashStr := hex.EncodeToString(hashByte)
	return hashStr
}
