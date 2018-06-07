package jsonerrors

import "github.com/json-iterator/go"

var jsoniterator = &Package{
	Name: "github.com/json-iterator/go",
	Decoders: []*Decoder{
		{
			Desc:      "jsoniter.ConfigCompatibleWithStandardLibrary",
			Unmarshal: jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
		},
	},
}
