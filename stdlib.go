package jsonerrors

import "encoding/json"

var stdlib = &Package{
	Name: "encoding/json",
	Desc: "stdlib",
	Decoders: []*Decoder{
		{
			Desc:      "json.Unmarshal",
			Unmarshal: json.Unmarshal,
		},
		{
			Desc: "json.Decoder.DisallowUnknownFields",
			Unmarshal: func(b []byte, v interface{}) error {
				dec := json.NewDecoder(buf(b))
				dec.DisallowUnknownFields()
				return dec.Decode(v)
			},
		},
	},
}
