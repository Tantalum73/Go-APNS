package goapns

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

func CertificateFromP12(filePath string, key string) (tls.Certificate, error) {
	p12Data, err := ioutil.ReadFile(filePath)
	fmt.Printf("Read Data %v error: %v\n", p12Data, err)
	if err != nil {
		return tls.Certificate{}, err
	}

	privateKey, crt, err := pkcs12.Decode(p12Data, key)
	fmt.Printf("Private key %v crt %v, error %v \n", privateKey, crt, err)

	return tls.Certificate{}, nil
}
