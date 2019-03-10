package wheel

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tk := NewToken("aa")
	tk.vendorEpoch = [4]byte{163, 193, 138, 1}
	tk.Generate()
	if !reflect.DeepEqual(tk.hashToken, []byte{73, 166, 126, 28, 190, 65, 221, 71,
		13, 33, 192, 172, 201, 202, 54, 88, 90, 230, 55, 7, 172, 142, 162, 103, 201,
		179, 104, 49, 62, 160, 217, 158}) {
		t.Errorf("hashToken value mismatch")
	}
}

func TestGetTokenString(t *testing.T) {
	tk := NewToken("aa")
	tk.vendorEpoch = [4]byte{163, 193, 138, 1}
	tk.Generate()
	tokenString := tk.GetTokenString()
	if tokenString != "87492217" {
		t.Errorf("token value mismatch (%s vs %s)", "87492217", tokenString)
	}
}
