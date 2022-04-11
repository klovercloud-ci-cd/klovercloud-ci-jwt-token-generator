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
	block, rest := pem.Decode([]byte(config.PrivateKey))
	if rest != nil {
		log.Print("key:", string(rest))
	}
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return privateKeyImported
}

func (Jwt) GetPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(config.Publickey))
	publicKeyImported, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return publicKeyImported
}
