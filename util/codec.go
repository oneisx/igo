package util

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base64"
    "encoding/hex"
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

func Base64Encode(s string) string {
    return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
    result, err := base64.StdEncoding.DecodeString(s)
    return string(result), err
}
