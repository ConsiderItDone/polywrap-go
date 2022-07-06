package json

import (
	"github.com/valyala/fastjson"
)

type JSON struct {
	*fastjson.Value
	*fastjson.Arena
}

func NewJSON() *JSON {
	return &JSON{
		Value: nil,
		Arena: new(fastjson.Arena),
	}
}

func Decode(value string) (*JSON, error) {
	var p fastjson.Parser
	v, err := p.Parse(value)
	if err != nil {
		return nil, err
	}

	return &JSON{
		Value: v,
		Arena: nil,
	}, nil
}

//func Encode(data *map[string]interface{}) (string, error) {
//	out, err := json.Marshal(data)
//	if err != nil {
//		return "", err
//	}
//
//	return string(out), nil
//}
