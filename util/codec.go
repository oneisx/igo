package util

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base64"
    "encoding/hex"
    "io"
)

func SHA1(str string) string {
    s := sha1.New()
    s.Write([]byte(str))
    return hex.EncodeToString(s.Sum(nil))
}

func HmacSHA1(str string, key string) string {
    h := hmac.New(sha1.New, []byte(key))
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func SHA256(str string) string {
    s := sha256.New()
    s.Write([]byte(str))
    return hex.EncodeToString(s.Sum(nil))
}

func HmacSHA256(str string, key string) string {
    h := hmac.New(sha256.New, []byte(key))
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func SHA512(str string) string {
    s := sha512.New()
    s.Write([]byte(str))
    return hex.EncodeToString(s.Sum(nil))
}

func HmacSHA512(str string, key string) string {
    h := hmac.New(sha512.New, []byte(key))
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func MD5(str string) string {
    m := md5.New()
    m.Write([]byte(str))
    return hex.EncodeToString(m.Sum(nil))
}

func HmacMD5(str string, key string) string {
    h := hmac.New(md5.New, []byte(key))
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func EncryptBase64(s string) string {
    return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecryptBase64(s string) (string, error) {
    result, err := base64.StdEncoding.DecodeString(s)
    return string(result), err
}

func EncryptAES(str string, key string) (string, error) {
    plainText := []byte(str)
    keyByte := []byte(key)
    block, err := aes.NewCipher(keyByte)
    if err != nil {
        return "", err
    }
    plainText = pkcs5Padding(plainText, block.BlockSize())
    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]
    _, _ = io.ReadFull(rand.Reader, iv)
    blockMode := cipher.NewCBCEncrypter(block, iv)
    blockMode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
    return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptAES(str string, key string) (string, error) {
    cipherText, _ := base64.StdEncoding.DecodeString(str)
    keyByte := []byte(key)
    block, err := aes.NewCipher(keyByte)
    if err != nil {
        return "", err
    }
    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]
    blockMode := cipher.NewCBCDecrypter(block, iv)
    blockMode.CryptBlocks(cipherText, cipherText)
    cipherText = pkcs5UnPadding(cipherText)
    return string(cipherText), nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
    padding := blockSize - len(cipherText)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(cipherText, padText...)
}

func pkcs5UnPadding(decrypted []byte) []byte {
    length := len(decrypted)
    unPadding := int(decrypted[length-1])
    return decrypted[:(length - unPadding)]
}
