package i2pasta_convert ///convert

/*import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
)*/

func (i *I2pconv)I2p32to64(b32 string) (string, error) {
	var err error
	/*raw64, err := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~").DecodeString(b64) //.DecodeString(b64)
	if Error(err, "i2pdig.go Base64 Conversion", string(raw64)) {
		hash := sha256.New()
		_, err := hash.Write([]byte(raw64)) //sha256.Sum256(raw64)
		if Error(err, "i2pdig.go Base32 Conversion") {
			b32 := strings.ToLower(strings.Replace(base32.StdEncoding.EncodeToString(hash.Sum(nil)), "=", "", -1)) + ".b32.i2p"
			os.Stderr.WriteString("#i2p Address information for: " + hostname)
			return b32, err
		} else {
			return "", err
		}
	}*/
	return "", err
}
