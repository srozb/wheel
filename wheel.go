package wheel

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

//Token token struct
type Token struct {
	timeTolerance uint
	timeOffset    uint
	vendorEpoch   [4]byte
	secret        []byte
	hashToken     []byte
}

//NewToken Returns a pointer to a new instance of token struct
func NewToken(secret string) (*Token, error) {
	t := new(Token)
	t.timeTolerance = 60
	t.timeOffset = 3600
	var err error
	t.secret, err = hex.DecodeString(secret)
	if err != nil {
		log.Printf("Error decoding token secret: %s", err)
	}
	return t, err
}

//SetTime updates vendorEpoch field with a value based on current time. This assures
//token will not be already expired.
func (t *Token) SetTime(epoch int64) {
	vendorEpoch := int64(math.Round((float64(epoch) + float64(t.timeOffset)) / float64(t.timeTolerance)))
	t.vendorEpoch = [4]byte{
		byte(vendorEpoch),
		byte(vendorEpoch >> 8),
		byte(vendorEpoch >> 16),
		byte(vendorEpoch >> 24),
	}
}

//Generate generate/regenerate token value
func (t *Token) Generate() error {
	if t.vendorEpoch[0] == 0 && t.vendorEpoch[1] == 0 && t.vendorEpoch[2] == 0 && t.vendorEpoch[3] == 0 {
		return errors.New("time not set, UpdateTime() hasn't been called")
	}
	data := append(t.secret, t.vendorEpoch[:]...)
	data = append(data, t.secret...)
	h := sha256.New()
	h.Write(data)
	t.hashToken = h.Sum(nil)
	return nil
}

//GetTokenString returns previously generated token
func (t *Token) GetTokenString() string {
	hh1 := binary.BigEndian.Uint64(append(make([]byte, 1), t.hashToken[:8]...))
	hh2 := binary.BigEndian.Uint64(append(make([]byte, 2), t.hashToken[7:13]...))
	bin1 := strconv.Itoa(int(hh1))
	bin2 := strconv.Itoa(int(hh2))
	return fmt.Sprintf("%s%s", string(bin1)[len(bin1)-4:], string(bin2)[len(bin2)-4:])
}
