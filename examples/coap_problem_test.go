package coap_problem

import (
	"encoding/hex"
	"testing"

	"github.com/go-ocf/go-coap"
)

var Examples = map[string]string{
	"private_ns_minimal":  "a2003982ae0105",
	"private_ns_full":     "a6003982ae0105026e756e6b6e6f776e206b65792069640318840478254b657920776974682069642030783031303230333034206e6f74207265676973746572656405d820782468747470733a2f2f707269766174652d6170692e6578616d706c652f6572726f72732f35",
	"private_ns_full_ext": "a725831a010203001a010203011a01020302003982ae0105026e756e6b6e6f776e206b65792069640318840478254b657920776974682069642030783031303230333034206e6f74207265676973746572656405d820782468747470733a2f2f707269766174652d6170692e6578616d706c652f6572726f72732f35",
}

func rigmarole(t *testing.T, p CoapProblem, exName string) {
	exp := Examples[exName]

	got := hex.EncodeToString(p.EncodeCBOR())
	if got != exp {
		t.Errorf("got: %s; want %s", got, exp)
	} else {
		t.Logf("HexEncoded # %s\n", got)
	}
}

func TestCoapProblem_EncodeCBOR_private_ns_minimal(t *testing.T) {
	p := CoapProblem{
		Namespace: -33455,
		Type:      5,
	}

	rigmarole(t, p, "private_ns_minimal")
}

func TestCoapProblem_EncodeCBOR_private_ns_full(t *testing.T) {
	p := CoapProblem{
		Namespace:    -33455,
		Type:         5,
		Title:        "unknown key id",
		ResponseCode: uint(coap.NotFound),
		Detail:       "Key with id 0x01020304 not registered",
		Instance:     "https://private-api.example/errors/5",
	}

	rigmarole(t, p, "private_ns_full")
}

//func TestCoapProblem_EncodeCBOR_private_ns_full_ext(t *testing.T) {
//	p := CoapProblem{
//		Namespace:    -33455,
//		Type:         5,
//		Title:        "unknown key id",
//		ResponseCode: uint(coap.NotFound),
//		Detail:       "Key with id 0x01020304 not registered",
//		Instance:     "https://private-api.example/errors/5",
//		AnExtension:  []uint{0x01020300, 0x01020301, 0x01020302},
//	}
//
//	rigmarole(t, p, "private_ns_full_ext")
//}
