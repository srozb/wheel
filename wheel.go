package wheel

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

//Token token struct
type Token struct {
	timeTolerance uint
	timeOffset    uint
	vendorEpoch   [4]byte
	secret        []byte
	hashToken     []byte
	token         string
}

//NewToken Return new instance of token struct
func NewToken(secret string) *Token {
	t := new(Token)
	t.timeTolerance = 60
	t.timeOffset = 3600
	binsecret, err := hex.DecodeString(secret)
	if err != nil {
		log.Printf("err: %s", err)
	}
	t.secret = binsecret
	t.vendorEpoch = generateVendorEpoch(t.timeOffset, t.timeTolerance)
	return t
}

func generateVendorEpoch(timeOffset, timeTolerance uint) [4]byte {
	vendorEpoch := int64(math.Round((float64(time.Now().Unix()) + float64(timeOffset)) / float64(timeTolerance)))
	return [4]byte{
		byte(vendorEpoch),
		byte(vendorEpoch >> 8),
		byte(vendorEpoch >> 16),
		byte(vendorEpoch >> 24),
	}
}

func (t *Token) hashSecretAndTime() {
	data := append(t.secret, t.vendorEpoch[:]...)
	data = append(data, t.secret...)
	h := sha256.New()
	h.Write(data)
	t.hashToken = h.Sum(nil)
}

func (t *Token) createTokenString() {
	hh1 := binary.BigEndian.Uint64(append(make([]byte, 1), t.hashToken[:8]...))
	hh2 := binary.BigEndian.Uint64(append(make([]byte, 2), t.hashToken[7:13]...))
	bin1 := strconv.Itoa(int(hh1))
	bin2 := strconv.Itoa(int(hh2))
	t.token = fmt.Sprintf("%s%s", string(bin1)[len(bin1)-4:], string(bin2)[len(bin2)-4:])
}

//Generate generate/regenerate token value
func (t *Token) Generate() {
	t.hashSecretAndTime()
}

//GetTokenString returns previously generated token
func (t *Token) GetTokenString() string {
	t.createTokenString()
	return t.token
}
