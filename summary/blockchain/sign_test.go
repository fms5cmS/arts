package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	signatureHex := "0xe40757f4c08789dae3c407ebeecc79724f9b650b73c9034ee25109b125549b070de56aa53a65b2301c08001a104b2a3d99fdc9cdbf7c7f6f8934df12cafdbc641b"
	message := "123"
	verified, err := VerifySignature(signatureHex, message, "0xAEcf7c588eAe107D98373Fd039D8FE2De26Bcf57")
	assert.Nil(t, err)
	assert.True(t, verified)
}
