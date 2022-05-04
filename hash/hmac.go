package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}
//hmac is a warpper around crypto/hmac package
// making it easier to use it in our code
type HMAC struct {
	hmac hash.Hash
}
//Hash will hash the provided input string using HMAC with 
//the secret key provided when the HMAC object was created

func (h HMAC) Hash(input string) string{
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)

}