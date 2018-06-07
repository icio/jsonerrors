package jsonerrors

import "bytes"

var buf = bytes.NewBuffer

type Package struct {
	Name     string
	Desc     string
	Decoders []*Decoder
}

type Decoder struct {
	Desc      string
	Unmarshal func([]byte, interface{}) error
}

func Packages() []*Package {
	return []*Package{
		stdlib,
		jsoniterator,
		koki,
	}
}
