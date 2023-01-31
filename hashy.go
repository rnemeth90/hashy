package hashy

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// OpFunc represents a hashing function defined below
type OpFunc func(string) string

// GenerateSHA224 is used to generate a sha224 hash
func GenerateSHA224(input string) string {
	hashy := sha256.New224()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateSHA256 is used to generate a sha256 hash
func GenerateSHA256(input string) string {
	hashy := sha256.New()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateSHA384 is used to generate a sha384 hash
func GenerateSHA384(input string) string {
	hashy := sha512.New384()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateSHA512 is used to generate a sha512 hash
func GenerateSHA512(input string) string {
	hashy := sha512.New()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateSHA512224 is used to generate a sha512_224 hash
func GenerateSHA512224(input string) string {
	hashy := sha512.New512_224()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateSHA512256 is used to generate a sha512_256 hash
func GenerateSHA512256(input string) string {
	hashy := sha512.New512_256()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}

// GenerateMD5 is used to generate an MD5 hash
func GenerateMD5(input string) string {
	hashy := md5.New()
	hashy.Write([]byte(input))
	str := fmt.Sprintf("%x\n", hashy.Sum(nil))
	return str
}
