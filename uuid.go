package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"strings"

	"github.com/trimark-jp/errors"
)

const (
	uuidByteLength       = 16
	errMessageReadLength = "rand read length is invalid"
	errMessage           = "generating UUID failed"
)

// V4 is UUID version 4.
type V4 struct {
	Raw []byte
}

func newV4() *V4 {
	return &V4{
		Raw: make([]byte, uuidByteLength),
	}
}

// GenV4 returns an UUID version 4.
func GenV4() (*V4, error) {
	result := newV4()
	n, e := rand.Read(result.Raw)

	if e != nil {
		return result, errors.WrapBySourceMsg(e, errMessage)
	}

	if n != uuidByteLength {
		return result, errors.WrapBySourceMsg(errors.New(errMessageReadLength), errMessage)
	}

	// Set the two most significant bits (bits 6 and 7) of the
	// clock_seq_hi_and_reserved to zero and one, respectively.
	result.Raw[8] = 0x80 | (result.Raw[8] & 0x3F)

	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number from
	result.Raw[6] = 0x40 | (result.Raw[6] & 0x0F)

	return result, nil
}

func (uuid *V4) String() string {
	hexText := hex.EncodeToString(uuid.Raw)
	log.Println(hexText)
	return strings.Join([]string{
		hexText[0:8],
		hexText[8:12],
		hexText[12:16],
		hexText[16:20],
		hexText[20:32],
	}, "-")
}
