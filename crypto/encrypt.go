package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha512"
	"io"
	"math/rand"
	"time"
)

const IV_SIZE int = 16
const VERSION_ONE byte = 0x1

func Encrypt(in io.Reader, out io.Writer, aesKey, hmacKey []byte) (err error) {
	iv := make([]byte, IV_SIZE)
	rand.Seed(time.Now().Unix())
	_, err = rand.Read(iv)
	if err != nil {
		return err
	}
	aes, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}
	ctr := cipher.NewCTR(aes, iv)
	hmacHash := hmac.New(sha512.New, hmacKey)

	out.Write([]byte{VERSION_ONE})

	w := io.MultiWriter(out, hmacHash)
	w.Write(iv)
}
