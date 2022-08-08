package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

func serializeCustomType(customType *CustomType) []byte {
	ctx := msgpack.NewContext("Serializing (encoding) env-type: CustomType")
	encoder := msgpack.NewWriteEncoder(ctx)
	writeCustomType(encoder, customType)
	return encoder.Buffer()
}

func writeCustomType(writer msgpack.Write, customType *CustomType) {
	writer.WriteMapLength(41)

	writer.Context().Push("str", "string", "writing property")
	writer.WriteString("str")
	writer.WriteString(customType.Str)
	writer.Context().Pop()

	writer.Context().Push("optStr", "string | null", "writing property")
	writer.WriteString("optStr")
	writer.WriteOptionalString(customType.OptStr)
	writer.Context().Pop()

	writer.Context().Push("u", "u32", "writing property")
	writer.WriteString("u")
	writer.WriteU32(customType.U)
	writer.Context().Pop()

	writer.Context().Push("optU", "Option<u32>", "writing property")
	writer.WriteString("optU")
	writer.WriteOptionalU32(customType.OptU)
	writer.Context().Pop()

	writer.Context().Push("u8", "u8", "writing property")
	writer.WriteString("u8")
	writer.WriteU8(customType.U8)
	writer.Context().Pop()

	writer.Context().Push("u16", "u16", "writing property")
	writer.WriteString("u16")
	writer.WriteU16(customType.U16)
	writer.Context().Pop()

	writer.Context().Push("u32", "u32", "writing property")
	writer.WriteString("u32")
	writer.WriteU32(customType.U32)
	writer.Context().Pop()

	writer.Context().Push("i", "i32", "writing property")
	writer.WriteString("i")
	writer.WriteI32(customType.I)
	writer.Context().Pop()

	writer.Context().Push("i8", "i8", "writing property")
	writer.WriteString("i8")
	writer.WriteI8(customType.I8)
	writer.Context().Pop()

	writer.Context().Push("i16", "i16", "writing property")
	writer.WriteString("i16")
	writer.WriteI16(customType.I16)
	writer.Context().Pop()

	writer.Context().Push("i32", "i32", "writing property")
	writer.WriteString("i32")
	writer.WriteI32(customType.I32)
	writer.Context().Pop()

	writer.Context().Push("bigint", "BigInt", "writing property")
	writer.WriteString("bigint")
	writer.WriteBigInt(customType.Bigint)
	writer.Context().Pop()

	writer.Context().Push("optBigint", "BigInt | null", "writing property")
	writer.WriteString("optBigint")
	writer.WriteOptionalBigInt(customType.OptBigint)
	writer.Context().Pop()

	writer.Context().Push("bignumber", "BigNumber", "writing property")
	writer.WriteString("bignumber")
	writer.WriteBigInt(customType.Bignumber)
	writer.Context().Pop()

	writer.Context().Push("optBignumber", "BigNumber | null", "writing property")
	writer.WriteString("optBignumber")
	writer.WriteOptionalBigInt(customType.OptBignumber)
	writer.Context().Pop()

	writer.Context().Push("json", "JSON.Value", "writing property")
	writer.WriteString("json")
	writer.WriteJson(customType.Json)
	writer.Context().Pop()

	writer.Context().Push("optJson", "JSON.Value | null", "writing property")
	writer.WriteString("optJson")
	writer.WriteOptionalJson(customType.OptJson)
	writer.Context().Pop()

	writer.Context().Push("bytes", "ArrayBuffer", "writing property")
	writer.WriteString("bytes")
	writer.WriteBytes(customType.Bytes)
	writer.Context().Pop()

	writer.Context().Push("optBytes", "ArrayBuffer | null", "writing property")
	writer.WriteString("optBytes")
	writer.WriteOptionalBytes(customType.OptBytes)
	writer.Context().Pop()

	writer.Context().Push("boolean", "bool", "writing property")
	writer.WriteString("boolean")
	writer.WriteBool(customType.Boolean)
	writer.Context().Pop()

	writer.Context().Push("optBoolean", "Option<bool>", "writing property")
	writer.WriteString("optBoolean")
	writer.WriteOptionalBool(customType.OptBoolean)
	writer.Context().Pop()

	writer.Context().Push("uArray", "Array<u32>", "writing property")
	writer.WriteString("uArray")
	writer.WriteArray([]interface{}(customType.UArray), func(encoder msgpack.Write, item interface{}) {
		writer.WriteU32(item.(uint32))
	})
	writer.Context().Pop()

	writer.Context().Push("uOptArray", "Array<u32> | null", "writing property")
	writer.WriteString("uOptArray")
	writer.WriteOptionalArray(customType.UOptArray, func(encoder msgpack.Write, item interface{}) {
		writer.WriteU32(item.(uint32))
	})
	writer.Context().Pop()

	writer.Context().Push("optUOptArray", "Array<Option<u32>> | null", "writing property")
	writer.WriteString("optUOptArray")
	writer.WriteOptionalArray(customType.OptUOptArray, func(encoder msgpack.Write, item interface{}) {
		writer.WriteOptionalU32(item.(container.Option))
	})
	writer.Context().Pop()

	writer.Context().Push("optStrOptArray", "Array<string | null> | null", "writing property")
	writer.WriteString("optStrOptArray")
	writer.WriteOptionalArray(customType.OptStrOptArray, func(encoder msgpack.Write, item interface{}) {
		writer.WriteOptionalString(item.(container.Option))
	})
	writer.Context().Pop()

	writer.Context().Push("uArrayArray", "Array<Array<u32>>", "writing property")
	writer.WriteString("uArrayArray")
	writer.WriteArray([]interface{}(customType.UArrayArray), func(encoder msgpack.Write, item interface{}) {
		encoder.WriteArray(item.([]interface{}), func(encoder msgpack.Write, item interface{}) {
			encoder.WriteU32(item.(uint32))
		})
	})
	writer.Context().Pop()

	writer.Context().Push("uOptArrayOptArray", "Array<Array<Option<u32>> | null>", "writing property")
	writer.WriteString("uOptArrayOptArray")
	writer.WriteArray([]interface{}(customType.UOptArrayOptArray), func(encoder msgpack.Write, item interface{}) {
		encoder.WriteOptionalArray(item.(container.Option), func(encoder msgpack.Write, item interface{}) {
			writer.WriteOptionalU32(item.(container.Option))
		})
	})
	writer.Context().Pop()

	writer.Context().Push("uArrayOptArrayArray", "Array<Array<Array<u32>> | null>", "writing property")
	writer.WriteString("uArrayOptArrayArray")
	writer.WriteArray([]interface{}(customType.UArrayOptArrayArray), func(encoder msgpack.Write, item interface{}) {
		encoder.WriteOptionalArray(item.(container.Option), func(encoder msgpack.Write, item interface{}) {
			encoder.WriteArray(item.([]interface{}), func(encoder msgpack.Write, item interface{}) {
				encoder.WriteU32(item.(uint32))
			})
		})
	})
	writer.Context().Pop()

	writer.Context().Push("crazyArray", "Array<Array<Array<Array<u32> | null>> | null> | null", "writing property")
	writer.WriteString("crazyArray")
	writer.WriteOptionalArray(customType.CrazyArray, func(encoder msgpack.Write, item interface{}) {
		encoder.WriteOptionalArray(item.(container.Option), func(encoder msgpack.Write, item interface{}) {
			encoder.WriteArray(item.([]interface{}), func(encoder msgpack.Write, item interface{}) {
				encoder.WriteOptionalArray(item.(container.Option), func(encoder msgpack.Write, item interface{}) {
					encoder.WriteU32(item.(uint32))
				})
			})
		})
	})
	writer.Context().Pop()

	writer.Context().Push("object", "Types.AnotherType", "writing property")
	writer.WriteString("object")
	AnotherTypeWrite(writer, &customType.Object)
	writer.Context().Pop()

	writer.Context().Push("optObject", "Types.AnotherType | null", "writing property")
	writer.WriteString("optObject")
	if customType.OptObject.IsSome() {
		v := customType.OptObject.MustGet().(AnotherType)
		AnotherTypeWrite(writer, &v)
	} else {
		writer.WriteNil()
	}
	writer.Context().Pop()

	writer.Context().Push("ObjectArray", "Array<Types.AnotherType>", "writing property")
	writer.WriteString("objectArray")
	writer.WriteArray([]interface{}(customType.ObjectArray), func(encoder msgpack.Write, item interface{}) {
		v := item.(AnotherType)
		AnotherTypeWrite(encoder, &v)
	})
	writer.Context().Pop()

	writer.Context().Push("optObjectArray", "Array<Types.AnotherType | null> | null", "writing property")
	writer.WriteString("optObjectArray")
	writer.WriteOptionalArray(customType.OptObjectArray, func(encoder msgpack.Write, item interface{}) {
		c := item.(container.Option)
		if c.IsSome() {
			v := c.MustGet().(AnotherType)
			AnotherTypeWrite(encoder, &v)
		} else {
			encoder.WriteNil()
		}
	})
	writer.Context().Pop()

	writer.Context().Push("en", "Types.CustomEnum", "writing property")
	writer.WriteString("en")
	writer.WriteI32(int32(customType.En))
	writer.Context().Pop()

	writer.Context().Push("optEnum", "Option<Types.CustomEnum>", "writing property")
	writer.WriteString("optEnum")
	writer.WriteOptionalI32(customType.OptEnum)
	writer.Context().Pop()

	writer.Context().Push("enumArray", "Array<Types.CustomEnum>", "writing property")
	writer.WriteString("enumArray")
	writer.WriteArray([]interface{}(customType.EnumArray), func(encoder msgpack.Write, item interface{}) {
		encoder.WriteI32(item.(int32))
	})
	writer.Context().Pop()

	writer.Context().Push("optEnumArray", "Array<Option<Types.CustomEnum>> | null", "writing property")
	writer.WriteString("optEnumArray")
	writer.WriteOptionalArray(customType.OptEnumArray, func(encoder msgpack.Write, item interface{}) {
		encoder.WriteOptionalI32(item.(container.Option))
	})
	writer.Context().Pop()

	writer.Context().Push("map", "Map<string, i32>", "writing property")
	writer.WriteString("map")
	writer.WriteMap(map[interface{}]interface{}(customType.Map), func(encoder msgpack.Write, key, value interface{}) {
		encoder.WriteString(key.(string))
		encoder.WriteI32(value.(int32))
	})
	writer.Context().Pop()

	writer.Context().Push("mapOfArr", "Map<string, Array<i32>>", "writing property")
	writer.WriteString("mapOfArr")
	writer.WriteMap(map[interface{}]interface{}(customType.MapOfArr), func(encoder msgpack.Write, key, value interface{}) {
		encoder.WriteString(key.(string))
		encoder.WriteArray(value.([]interface{}), func(encoder msgpack.Write, item interface{}) {
			encoder.WriteI32(item.(int32))
		})
	})
	writer.Context().Pop()

	writer.Context().Push("mapOfObj", "Map<string, Types.AnotherType>", "writing property")
	writer.WriteString("mapOfObj")
	writer.WriteMap(map[interface{}]interface{}(customType.MapOfObj), func(encoder msgpack.Write, key, value interface{}) {
		encoder.WriteString(key.(string))
		v := value.(AnotherType)
		AnotherTypeWrite(encoder, &v)
	})
	writer.Context().Pop()

	writer.Context().Push("mapOfArrOfObj", "Map<string, Array<Types.AnotherType>>", "writing property")
	writer.WriteString("mapOfArrOfObj")
	writer.WriteMap(map[interface{}]interface{}(customType.MapOfArrOfObj), func(encoder msgpack.Write, key, value interface{}) {
		encoder.WriteString(key.(string))
		encoder.WriteArray(value.([]interface{}), func(encoder msgpack.Write, item interface{}) {
			v := value.(AnotherType)
			AnotherTypeWrite(encoder, &v)
		})
	})
	writer.Context().Pop()
}

func deserializeCustomType(data []byte) *CustomType {
	ctx := msgpack.NewContext("Deserializing env-type CustomType")
	reader := msgpack.NewReadDecoder(ctx, data)
	return readCustomType(reader)
}

func readCustomType(reader msgpack.Read) *CustomType {
	var _str string
	var _strSet bool
	var _optStr container.Option
	var _u uint32
	var _uSet bool
	var _optU container.Option
	var _u8 uint8
	var _u8Set bool
	var _u16 uint16
	var _u16Set bool
	var _u32 uint32
	var _u32Set bool
	var _i int32
	var _iSet bool
	var _i8 int8
	var _i8Set bool
	var _i16 int16
	var _i16Set bool
	var _i32 int32
	var _i32Set bool
	var _bigint *big.Int
	var _bigintSet bool
	var _optBigint container.Option
	var _bignumber *big.Int
	var _bignumberSet bool
	var _optBignumber container.Option
	var _json *fastjson.Value
	var _jsonSet bool
	var _optJson container.Option
	var _bytes []byte
	var _bytesSet bool
	var _optBytes container.Option
	var _boolean bool
	var _booleanSet bool
	var _optBoolean container.Option
	var _uArray []uint32
	var _uArraySet bool
	var _uOptArray container.Option
	var _optUOptArray container.Option
	var _optStrOptArray container.Option
	var _uArrayArray [][]uint32
	var _uArrayArraySet bool
	var _uOptArrayOptArray []container.Option
	var _uOptArrayOptArraySet bool
	var _uArrayOptArrayArray []container.Option
	var _uArrayOptArrayArraySet bool
	var _crazyArray container.Option
	var _object AnotherType
	var _objectSet bool
	var _optObject container.Option
	var _objectArray []AnotherType
	var _objectArraySet bool
	var _optObjectArray container.Option
	var _en CustomEnum
	var _enSet bool
	var _optEnum container.Option
	var _enumArray []CustomEnum
	var _enumArraySet bool
	var _optEnumArray container.Option
	var _map map[string]int32
	var _mapSet bool
	var _mapOfArr map[string][]int32
	var _mapOfArrSet bool
	var _mapOfObj map[string]AnotherType
	var _mapOfObjSet bool
	var _mapOfArrOfObj map[string][]AnotherType
	var _mapOfArrOfObjSet bool

	for i := int32(reader.ReadMapLength()); i > 0; i-- {
		field := reader.ReadString()

		reader.Context().Push(field, "unknown", "searching for property type")
		if field == "str" {
			reader.Context().Push(field, "string", "type found, reading property")
			_str = reader.ReadString()
			_strSet = true
			reader.Context().Pop()
		} else if field == "optStr" {
			reader.Context().Push(field, "string | null", "type found, reading property")
			_optStr = reader.ReadOptionalString()
			reader.Context().Pop()
		} else if field == "u" {
			reader.Context().Push(field, "u32", "type found, reading property")
			_u = reader.ReadU32()
			_uSet = true
			reader.Context().Pop()
		} else if field == "optU" {
			reader.Context().Push(field, "Option<u32>", "type found, reading property")
			_optU = reader.ReadOptionalU32()
			reader.Context().Pop()
		} else if field == "u8" {
			reader.Context().Push(field, "u8", "type found, reading property")
			_u8 = reader.ReadU8()
			_u8Set = true
			reader.Context().Pop()
		} else if field == "u16" {
			reader.Context().Push(field, "u16", "type found, reading property")
			_u16 = reader.ReadU16()
			_u16Set = true
			reader.Context().Pop()
		} else if field == "u32" {
			reader.Context().Push(field, "u32", "type found, reading property")
			_u32 = reader.ReadU32()
			_u32Set = true
			reader.Context().Pop()
		} else if field == "i" {
			reader.Context().Push(field, "i32", "type found, reading property")
			_i = reader.ReadI32()
			_iSet = true
			reader.Context().Pop()
		} else if field == "i8" {
			reader.Context().Push(field, "i8", "type found, reading property")
			_i8 = reader.ReadI8()
			_i8Set = true
			reader.Context().Pop()
		} else if field == "i16" {
			reader.Context().Push(field, "i16", "type found, reading property")
			_i16 = reader.ReadI16()
			_i16Set = true
			reader.Context().Pop()
		} else if field == "i32" {
			reader.Context().Push(field, "i32", "type found, reading property")
			_i32 = reader.ReadI32()
			_i32Set = true
			reader.Context().Pop()
		} else if field == "bigint" {
			reader.Context().Push(field, "BigInt", "type found, reading property")
			_bigint = reader.ReadBigInt()
			_bigintSet = true
			reader.Context().Pop()
		} else if field == "optBigint" {
			reader.Context().Push(field, "BigInt | null", "type found, reading property")
			_optBigint = reader.ReadOptionalBigInt()
			reader.Context().Pop()
		} else if field == "bignumber" {
			reader.Context().Push(field, "BigNumber", "type found, reading property")
			_bignumber = reader.ReadBigInt()
			_bignumberSet = true
			reader.Context().Pop()
		} else if field == "optBignumber" {
			reader.Context().Push(field, "BigNumber | null", "type found, reading property")
			_optBignumber = reader.ReadOptionalBigInt()
			reader.Context().Pop()
		} else if field == "json" {
			reader.Context().Push(field, "JSON.Value", "type found, reading property")
			_json = reader.ReadJson()
			_jsonSet = true
			reader.Context().Pop()
		} else if field == "optJson" {
			reader.Context().Push(field, "JSON.Value | null", "type found, reading property")
			_optJson = reader.ReadOptionalJson()
			reader.Context().Pop()
		} else if field == "bytes" {
			reader.Context().Push(field, "ArrayBuffer", "type found, reading property")
			_bytes = reader.ReadBytes()
			_bytesSet = true
			reader.Context().Pop()
		} else if field == "optBytes" {
			reader.Context().Push(field, "ArrayBuffer | null", "type found, reading property")
			_optBytes = reader.ReadOptionalBytes()
			reader.Context().Pop()
		} else if field == "boolean" {
			reader.Context().Push(field, "bool", "type found, reading property")
			_boolean = reader.ReadBool()
			_booleanSet = true
			reader.Context().Pop()
		} else if field == "optBoolean" {
			reader.Context().Push(field, "Option<bool>", "type found, reading property")
			_optBoolean = reader.ReadOptionalBool()
			reader.Context().Pop()
		} else if field == "uArray" {
			reader.Context().Push(field, "Array<u32>", "type found, reading property")
			reader.ReadArray(func(reader msgpack.Read) interface{} {
				v := reader.ReadU32()
				_uArray = append(_uArray, v)
				return v
			})
			_uArraySet = true
			reader.Context().Pop()
		} else if field == "uOptArray" {
			reader.Context().Push(field, "Array<u32> | null", "type found, reading property")
			_uOptArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				return reader.ReadU32()
			})
			reader.Context().Pop()
		} else if field == "optUOptArray" {
			reader.Context().Push(field, "Array<Option<u32>> | null", "type found, reading property")
			_optUOptArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				return reader.ReadOptionalU32()
			})
			reader.Context().Pop()
		} else if field == "optStrOptArray" {
			reader.Context().Push(field, "Array<string | null> | null", "type found, reading property")
			_optStrOptArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				return reader.ReadOptionalString()
			})
			reader.Context().Pop()
		} else if field == "uArrayArray" {
			reader.Context().Push(field, "Array<Array<u32>>", "type found, reading property")
			_uArrayArray = reader.ReadArray(func(reader msgpack.Read) interface{} {
				return reader.ReadArray(func(reader msgpack.Read) interface{} {
					return reader.ReadU32()
				})
			})
			_uArrayArraySet = true
			reader.Context().Pop()
		} else if field == "uOptArrayOptArray" {
			reader.Context().Push(field, "Array<Array<Option<u32>> | null>", "type found, reading property")
			_uOptArrayOptArray = reader.ReadArray(func(reader msgpack.Read) interface{} {
				return reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
					return reader.ReadOptionalU32()
				})
			})
			_uOptArrayOptArraySet = true
			reader.Context().Pop()
		} else if field == "uArrayOptArrayArray" {
			reader.Context().Push(field, "Array<Array<Array<u32>> | null>", "type found, reading property")
			_uArrayOptArrayArray = reader.ReadArray(func(reader msgpack.Read) interface{} {
				return reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
					return reader.ReadArray(func(reader msgpack.Read) interface{} {
						return reader.ReadU32()
					})
				})
			})
			_uArrayOptArrayArraySet = true
			reader.Context().Pop()
		} else if field == "crazyArray" {
			reader.Context().Push(field, "Array<Array<Array<Array<u32> | null>> | null> | null", "type found, reading property")
			_crazyArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				return reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
					return reader.ReadArray(func(reader msgpack.Read) interface{} {
						return reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
							return reader.ReadU32()
						})
					})
				})
			})
			reader.Context().Pop()
		} else if field == "object" {
			reader.Context().Push(field, "Types.AnotherType", "type found, reading property")
			object := AnotherTypeRead(reader)
			_object = *object
			_objectSet = true
			reader.Context().Pop()
		} else if field == "optObject" {
			reader.Context().Push(field, "Types.AnotherType | null", "type found, reading property")
			if reader.IsNil() {
				_optObject = container.None()
			} else {
				object := AnotherTypeRead(reader)
				_optObject = container.Some(*object)
			}
			reader.Context().Pop()
		} else if field == "objectArray" {
			reader.Context().Push(field, "Array<Types.AnotherType>", "type found, reading property")
			_objectArray = reader.ReadArray(func(reader msgpack.Read) interface{} {
				object := AnotherTypeRead(reader)
				return *object
			})
			_objectArraySet = true
			reader.Context().Pop()
		} else if field == "optObjectArray" {
			reader.Context().Push(field, "Array<Types.AnotherType | null> | null", "type found, reading property")
			_optObjectArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				if reader.IsNil() {
					return container.None()
				}
				object := AnotherTypeRead(reader)
				return *object
			})
			reader.Context().Pop()
		} else if field == "en" {
			reader.Context().Push(field, "Types.CustomEnum", "type found, reading property")
			v := reader.ReadI32()
			SanitizeCustomEnumValue(int(v))
			_en = CustomEnum(v)
			_enSet = true
			reader.Context().Pop()
		} else if field == "optEnum" {
			reader.Context().Push(field, "Option<Types.CustomEnum>", "type found, reading property")
			if reader.IsNil() {
				_optEnum = container.None()
			} else {
				v := reader.ReadI32()
				SanitizeCustomEnumValue(int(v))
				_optEnum = container.Some(CustomEnum(v))
			}
			reader.Context().Pop()
		} else if field == "enumArray" {
			reader.Context().Push(field, "Array<Types.CustomEnum>", "type found, reading property")
			_enumArray = reader.ReadArray(func(reader msgpack.Read) interface{} {
				v := reader.ReadI32()
				SanitizeCustomEnumValue(int(v))
				return CustomEnum(v)
			})
			_enumArraySet = true
			reader.Context().Pop()
		} else if field == "optEnumArray" {
			reader.Context().Push(field, "Array<Option<Types.CustomEnum>> | null", "type found, reading property")
			_optEnumArray = reader.ReadOptionalArray(func(reader msgpack.Read) interface{} {
				if reader.IsNil() {
					return container.None()
				}
				v := reader.ReadI32()
				SanitizeCustomEnumValue(int(v))
				return container.Some(CustomEnum(v))
			})
			reader.Context().Pop()
		} else if field == "map" {
			reader.Context().Push(field, "Map<string, i32>", "type found, reading property")
			_map = reader.ReadMap(func(reader msgpack.Read) (interface{}, interface{}) {
				return reader.ReadString(), reader.ReadI32()
			})
			_mapSet = true
			reader.Context().Pop()
		} else if field == "mapOfArr" {
			reader.Context().Push(field, "Map<string, Array<i32>>", "type found, reading property")
			_mapOfArr = reader.ReadMap(func(reader msgpack.Read) (interface{}, interface{}) {
				return reader.ReadString(), reader.ReadArray(func(reader msgpack.Read) interface{} {
					return reader.ReadI32()
				})
			})
			_mapOfArrSet = true
			reader.Context().Pop()
		} else if field == "mapOfObj" {
			reader.Context().Push(field, "Map<string, Types.AnotherType>", "type found, reading property")
			_mapOfObj = reader.ReadMap(func(reader msgpack.Read) (interface{}, interface{}) {
				key := reader.ReadString()
				obj := AnotherTypeRead(reader)
				return key, *obj
			})
			_mapOfObjSet = true
			reader.Context().Pop()
		} else if field == "mapOfArrOfObj" {
			reader.Context().Push(field, "Map<string, Array<Types.AnotherType>>", "type found, reading property")
			_mapOfArrOfObj = reader.ReadMap(func(reader msgpack.Read) (interface{}, interface{}) {
				return reader.ReadString(), reader.ReadArray(func(reader msgpack.Read) interface{} {
					return AnotherTypeRead(reader)
				})
			})
			reader.Context().Pop()
		}
		reader.Context().Pop()
	}
	if !_strSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'str: String'"))
	}
	if !_uSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'u: UInt'"))
	}
	if !_u8Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'u8: UInt8'"))
	}
	if !_u16Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'u16: UInt16'"))
	}
	if !_u32Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'u32: UInt32'"))
	}
	if !_iSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'i: Int'"))
	}
	if !_i8Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'i8: Int8'"))
	}
	if !_i16Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'i16: Int16'"))
	}
	if !_i32Set {
		panic(reader.Context().PrintWithContext("Missing required property: 'i32: Int32'"))
	}
	if !_bigintSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'bigint: BigInt'"))
	}
	if !_bignumberSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'bignumber: BigNumber'"))
	}
	if !_jsonSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'json: JSON'"))
	}
	if !_bytesSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'bytes: Bytes'"))
	}
	if !_booleanSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'boolean: Boolean'"))
	}
	if !_uArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'uArray: [UInt]'"))
	}
	if !_uArrayArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'uArrayArray: [[UInt]]'"))
	}
	if !_uOptArrayOptArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'uOptArrayOptArray: [[UInt32]]'"))
	}
	if !_uArrayOptArrayArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'uArrayOptArrayArray: [[[UInt32]]]'"))
	}
	if !_objectSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'object: AnotherType'"))
	}
	if !_objectArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'objectArray: [AnotherType]'"))
	}
	if !_enSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'en: CustomEnum'"))
	}
	if !_enumArraySet {
		panic(reader.Context().PrintWithContext("Missing required property: 'enumArray: [CustomEnum]'"))
	}
	if !_mapSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'map: Map<String, Int>'"))
	}
	if !_mapOfArrSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'mapOfArr: Map<String, [Int]>'"))
	}
	if !_mapOfObjSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'mapOfObj: Map<String, AnotherType>'"))
	}
	if !_mapOfArrOfObjSet {
		panic(reader.Context().PrintWithContext("Missing required property: 'mapOfArrOfObj: Map<String, [AnotherType]>'"))
	}
	return &CustomType{
		Str:                 _str,
		OptStr:              _optStr,
		U:                   _u,
		OptU:                _optU,
		U8:                  _u8,
		U16:                 _u16,
		U32:                 _u32,
		I:                   _i,
		I8:                  _i8,
		I16:                 _i16,
		I32:                 _i32,
		Bigint:              _bigint,
		OptBigint:           _optBigint,
		Bignumber:           _bignumber,
		OptBignumber:        _optBignumber,
		Json:                _json,
		OptJson:             _optJson,
		Bytes:               _bytes,
		OptBytes:            _optBytes,
		Boolean:             _boolean,
		OptBoolean:          _optBoolean,
		UArray:              _uArray,
		UOptArray:           _uOptArray,
		OptUOptArray:        _optUOptArray,
		OptStrOptArray:      _optStrOptArray,
		UArrayArray:         _uArrayArray,
		UOptArrayOptArray:   _uOptArrayOptArray,
		UArrayOptArrayArray: _uArrayOptArrayArray,
		CrazyArray:          _crazyArray,
		Object:              _object,
		OptObject:           _optObject,
		ObjectArray:         _objectArray,
		OptObjectArray:      _optObjectArray,
		En:                  _en,
		OptEnum:             _optEnum,
		EnumArray:           _enumArray,
		OptEnumArray:        _optEnumArray,
		Map:                 _map,
		MapOfArr:            _mapOfArr,
		MapOfObj:            _mapOfObj,
		MapOfArrOfObj:       _mapOfArrOfObj,
	}

}
