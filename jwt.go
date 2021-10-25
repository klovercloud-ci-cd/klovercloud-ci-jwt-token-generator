package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/klovercloud-ci/config"
	"log"
	"time"
)
type Jwt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (j Jwt) GenerateToken(duration int64, data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Duration(duration) * time.Hour,
		"iat": time.Now().Unix(),
		"data":data,
	}
	tokenString, err := token.SignedString(config.PrivateKey)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}
func(Jwt) GetPrivateKey() *rsa.PrivateKey {
	block,_ := pem.Decode([]byte(config.PrivateKey))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	return privateKeyImported
}

func(Jwt) GetPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(config.Publickey))
	publicKeyImported, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return publicKeyImported
}
