package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type CustomType struct {
	Str                 string
	OptStr              container.Option // string
	U                   uint32
	OptU                container.Option // uint32
	U8                  uint8
	U16                 uint16
	U32                 uint32
	I                   int32
	I8                  int8
	I16                 int16
	I32                 int32
	Bigint              *big.Int
	OptBigint           container.Option // *big.Int
	Bignumber           *big.Int
	OptBignumber        container.Option // *big.Int
	Json                *fastjson.Value
	OptJson             container.Option // *fastjson.Value
	Bytes               []byte
	OptBytes            container.Option // []byte
	Boolean             bool
	OptBoolean          container.Option // bool
	UArray              []uint32
	UOptArray           container.Option // []uint32
	OptUOptArray        container.Option // Array<Option<u32>> | null;
	OptStrOptArray      container.Option // Array<string | null> | null;
	UArrayArray         [][]uint32
	UOptArrayOptArray   []container.Option // Array<Array<Option<u32>> | null>;
	UArrayOptArrayArray []container.Option // Array<Array<Array<u32>> | null>;
	CrazyArray          container.Option   // Array<Array<Array<Array<u32> | null>> | null> | null;
	Object              AnotherType
	OptObject           container.Option // AnotherType | null;
	ObjectArray         []AnotherType
	OptObjectArray      container.Option // Array<Types.AnotherType | null> | null;
	En                  CustomEnum
	OptEnum             container.Option // Option<Types.CustomEnum>;
	EnumArray           []CustomEnum
	OptEnumArray        container.Option // Array<Option<Types.CustomEnum>> | null;
	Map                 map[string]int32
	MapOfArr            map[string][]int32
	MapOfObj            map[string]AnotherType
	MapOfArrOfObj       map[string][]AnotherType
}

func CustomTypeToBuffer(env *CustomType) []byte {
	return serializeCustomType(env)
}

func CustomTypeFromBuffer(data []byte) *CustomType {
	return deserializeCustomType(data)
}

func CustomTypeWrite(writer msgpack.Write, env *CustomType) {
	writeCustomType(writer, env)
}

func CustomTypeRead(reader msgpack.Read) *CustomType {
	return readCustomType(reader)
}
