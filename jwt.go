package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci/config"
	"log"
	"time"
)

var RsaKeys *Jwt = nil

type Jwt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (j Jwt) GetRsaKeys() *Jwt {
	if RsaKeys == nil {
		RsaKeys = &Jwt{
			PrivateKey: j.GetPrivateKey(),
			PublicKey:  j.GetPublicKey(),
		}
	}
	return RsaKeys
}
func (j Jwt) GenerateToken(duration int64, data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp":  time.Now().UTC().Add(time.Duration(duration) * time.Hour).Unix(),
		"iat":  time.Now().UTC().Unix(),
		"data": data,
	}
	tokenString, err := token.SignedString(j.GetRsaKeys().PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (Jwt) GetPrivateKey() *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(config.PrivateKey))
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil

	} else {
		return rsaKey
	}
}

func (Jwt) GetPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(config.Publickey))
	pkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	rsaKey, ok := pkey.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("got unexpected key type: %T", pkey)
	}
	return rsaKey
}
