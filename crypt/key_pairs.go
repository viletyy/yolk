/*
 * @Date: 2021-03-11 18:02:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-26 23:00:24
 * @FilePath: /yolk/crypt/key_pairs.go
 */
package crypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// ecdsaCmd represents the doc command
func KeyPairs(keyName string) {
	//elliptic.P256(),elliptic.P384(),elliptic.P521()

	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(fmt.Sprintf("Genera KeyPairs Error: %v", err))
	}
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	privateBs := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	privateFile, err := os.Create(keyName + ".private.pem")
	if err != nil {
		panic(fmt.Sprintf("Genera KeyPairs Error: %v", err))
	}
	_, err = privateFile.Write(privateBs)
	if err != nil {
		panic(fmt.Sprintf("Genera KeyPairs Error: %v", err))
	}
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	publicBs := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	publicKeyFile, err := os.Create(keyName + ".public.pem")
	if err != nil {
		panic(fmt.Sprintf("Genera KeyPairs Error: %v", err))
	}
	_, err = publicKeyFile.Write(publicBs)
	if err != nil {
		panic(fmt.Sprintf("Genera KeyPairs Error: %v", err))
	}
}
