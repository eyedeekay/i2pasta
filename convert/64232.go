package i2pconv ///vert

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"github.com/eyedeekay/i2pasta/nup"
	"strings"
    "log"
)

type I2pconv struct {
	l i2pasta.I2plog
}

func (i *I2pconv) I2p64to32(b64 string) (string, error) {
	raw64, err := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~").DecodeString(b64)
	if i.l.Error(err, "i2pdig.go Base64 Conversion", string(raw64)) {
		hash := sha256.New()
		_, err := hash.Write([]byte(raw64)) //sha256.Sum256(raw64)
		if i.l.Error(err, "i2pdig.go Base32 Conversion") {
			b32 := strings.ToLower(strings.Replace(base32.StdEncoding.EncodeToString(hash.Sum(nil)), "=", "", -1))
			return b32, err
		} else {
			return "", err
		}
	}
	return "", err
}

func I2p64to32(b64 string) (string, error) {
	raw64, err := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~").DecodeString(b64)
	if err == nil {
		hash := sha256.New()
		_, err := hash.Write([]byte(raw64)) //sha256.Sum256(raw64)
		if err == nil {
			b32 := strings.ToLower(strings.Replace(base32.StdEncoding.EncodeToString(hash.Sum(nil)), "=", "", -1))
			log.Printf(b32)
            return b32, err
		} else {
			return "", err
		}
	}
	return "", err
}