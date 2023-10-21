package uniqueIdentifier

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const uuidLen = 36

func TestUUID(t *testing.T) {
	v1, _ := uuid.NewUUID()
	t.Log("v1 uuid:", v1.String())
	assert.Len(t, v1.String(), uuidLen)

	// uuid.NewDCEPerson() == uuid.NewDCESecurity(uuid.Person, uint32(os.Getuid()))
	// uuid.NewDCEGroup() == uuid.NewDCESecurity(uuid.Group, uint32(os.Getgid()))
	v2, _ := uuid.NewDCESecurity(uuid.Person, 10)
	t.Log("v2 uuid:", v2.String())
	assert.Len(t, v2.String(), uuidLen)

	v3 := uuid.NewMD5(uuid.New(), []byte("abcdefg"))
	t.Log("v3 uuid:", v3.String())
	assert.Len(t, v3.String(), uuidLen)

	v4, _ := uuid.NewRandom()
	t.Log("v4 uuid:", v4.String())
	assert.Len(t, v4.String(), uuidLen)

	v5 := uuid.NewSHA1(uuid.New(), []byte("abcdefg"))
	t.Log("v5 uuid:", v5.String())
	assert.Len(t, v5.String(), uuidLen)
}
