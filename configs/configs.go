package configs

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func LoadPrivateKey() *rsa.PrivateKey {

	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	priv, _ := ioutil.ReadFile(privKeyFile)
	privKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priv)
	return privKey
}

func LoadPublicKey() *rsa.PublicKey {

	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	pub, _ := ioutil.ReadFile(pubKeyFile)
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pub)
	return pubKey
}

func LoadTokenExp() int64 {
	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	idExp, _ := strconv.ParseInt(idTokenExp, 0, 64)
	return idExp
}
