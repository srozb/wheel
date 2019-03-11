package wheel

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tk, _ := NewToken("aa")
	tk.vendorEpoch = [4]byte{163, 193, 138, 1}
	_ = tk.Generate()
	if !reflect.DeepEqual(tk.hashToken, []byte{73, 166, 126, 28, 190, 65, 221, 71,
		13, 33, 192, 172, 201, 202, 54, 88, 90, 230, 55, 7, 172, 142, 162, 103, 201,
		179, 104, 49, 62, 160, 217, 158}) {
		t.Errorf("hashToken value mismatch")
	}
	tk.vendorEpoch = [4]byte{}
	if tk.Generate() == nil {
		t.Errorf("No error on Generate() with zeroed time")
	}
}

func TestGetTokenString(t *testing.T) {
	tk, err := NewToken("aa")
	if err != nil {
		t.Errorf("Error during token creation: %s", err)
	}
	tk.vendorEpoch = [4]byte{163, 193, 138, 1}
	_ = tk.Generate()
	tokenString := tk.GetTokenString()
	if tokenString != "87492217" {
		t.Errorf("token value mismatch (%s vs %s)", "87492217", tokenString)
	}
}

func TestNewToken(t *testing.T) {
	_, err := NewToken("zz")
	if err == nil {
		t.Errorf("no error on invalid hexstring")
	}
}

func TestSetTime(t *testing.T) {
	tk, _ := NewToken("aa")
	tk.SetTime(1552308361)
	if !reflect.DeepEqual(tk.vendorEpoch, [4]byte{250, 197, 138, 1}) {
		t.Errorf("vendorEpoch value mismatch: %v should be %v", tk.vendorEpoch, []byte{250, 197, 138, 1})
	}
}
