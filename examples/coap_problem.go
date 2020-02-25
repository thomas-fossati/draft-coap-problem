package coap_problem

import (
	"reflect"

	"github.com/ugorji/go/codec"
)

const UriCBORTag = 32

type Uri string

// URI encoder
func (u Uri) ConvertExt(v interface{}) interface{} {
	return string(v.(Uri))
}

func (Uri) UpdateExt(dest interface{}, v interface{}) {}

// CoapProblem provides the (fixed, for now) structure of a CoAP Problem
// Details.  The struct tags allow encoding to as well as decoding from CBOR.
type CoapProblem struct {
	_struct      bool   `codec:",int"`
	Namespace    int    `codec:"0"`            //   0 -> int
	Type         uint64 `codec:"1"`            //   1 -> uint
	Title        string `codec:"2,omitempty"`  // ? 2 -> text
	ResponseCode uint   `codec:"3,omitempty"`  // ? 3 -> uint .size 1,
	Detail       string `codec:"4,omitempty"`  // ? 4 -> text
	Instance     Uri    `codec:"5,omitempty"`  // ? 5 -> uri -- i.e., #6.32(tstr)
	AnExtension  []uint `codec:"-6,omitempty"` // ? -6 -> []`
}

type Ext struct{}

func GetCoapProblemHandle() (h *codec.CborHandle) {
	h = new(codec.CborHandle)
	h.Canonical = true // sort map keys

	var u Uri
	h.SetInterfaceExt(reflect.TypeOf(Uri("")), UriCBORTag, u)

	return h
}

func (p CoapProblem) EncodeCBOR() []byte {
	b := make([]byte, 0, 128)

	h := GetCoapProblemHandle()

	codec.NewEncoderBytes(&b, h).MustEncode(p)

	return b
}
