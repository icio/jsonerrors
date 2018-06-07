package jsonerrors

import (
	"github.com/koki/json"
	"github.com/koki/json/jsonutil"
)

var koki = &Package{
	Name: "github.com/koki/json",
	Decoders: []*Decoder{
		{
			Desc:      "json.Unmarshal",
			Unmarshal: json.Unmarshal,
		},
		{
			Desc: "json.UnmarshalMap",
			Unmarshal: func(b []byte, v interface{}) error {
				vMap := make(map[string]interface{})
				if err := json.Unmarshal(b, &vMap); err != nil {
					return err
				}
				if err := jsonutil.UnmarshalMap(vMap, v); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Desc: "json.UnmarshalMap + jsonutil.ExtraneousFieldPaths",
			Unmarshal: func(b []byte, v interface{}) error {
				vMap := make(map[string]interface{})
				if err := json.Unmarshal(b, &vMap); err != nil {
					return err
				}
				if err := jsonutil.UnmarshalMap(vMap, v); err != nil {
					return err
				}
				if extraPaths, err := jsonutil.ExtraneousFieldPaths(vMap, v); err != nil {
					return err
				} else if len(extraPaths) > 0 {
					return &jsonutil.ExtraneousFieldsError{Paths: extraPaths}
				}
				return nil
			},
		},
	},
}
