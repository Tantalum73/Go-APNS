package goapns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tantalum73/Go-APNS"
)

func TestConnectionCertificateWrongPath(t *testing.T) {
	pathToCert := "example/nowhere"
	conn, err := goapns.NewConnection(pathToCert, "wrongPassword")
	assert.Error(t, err)
	assert.Nil(t, conn)
}
func TestConnectionCertificateWrongPassphrase(t *testing.T) {
	pathToCert := "example/certificate-valid-encrypted.p12"
	conn, err := goapns.NewConnection(pathToCert, "wrongPassword")
	assert.Error(t, err)
	assert.Nil(t, conn)
}
