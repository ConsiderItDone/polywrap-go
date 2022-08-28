package format

type Format uint8

const (
	ERROR                       Format = 0
	FOUR_LEAST_SIG_BITS_IN_BYTE Format = 0x0f
	FOUR_SIG_BITS_IN_BYTE       Format = 0xf0
	POSITIVE_FIXINT             Format = 0x00
	FIXMAP                      Format = 0x80
	FIXARRAY                    Format = 0x90
	FIXSTR                      Format = 0xa0
	NIL                         Format = 0xc0
	RESERVED                    Format = 0xc1
	FALSE                       Format = 0xc2
	TRUE                        Format = 0xc3
	BIN8                        Format = 0xc4
	BIN16                       Format = 0xc5
	BIN32                       Format = 0xc6
	EXT8                        Format = 0xc7
	EXT16                       Format = 0xc8
	EXT32                       Format = 0xc9
	FLOAT32                     Format = 0xca
	FLOAT64                     Format = 0xcb
	UINT8                       Format = 0xcc
	UINT16                      Format = 0xcd
	UINT32                      Format = 0xce
	UINT64                      Format = 0xcf
	INT8                        Format = 0xd0
	INT16                       Format = 0xd1
	INT32                       Format = 0xd2
	INT64                       Format = 0xd3
	FIXEXT1                     Format = 0xd4
	FIXEXT2                     Format = 0xd5
	FIXEXT4                     Format = 0xd6
	FIXEXT8                     Format = 0xd7
	FIXEXT16                    Format = 0xd8
	STR8                        Format = 0xd9
	STR16                       Format = 0xda
	STR32                       Format = 0xdb
	ARRAY16                     Format = 0xdc
	ARRAY32                     Format = 0xdd
	MAP16                       Format = 0xde
	MAP32                       Format = 0xdf
	NEGATIVE_FIXINT             Format = 0xe0
)

func ToString(f Format) string {
	switch f {
	case NIL:
		return "nil"
	case RESERVED:
		return "reserved"
	case TRUE, FALSE:
		return "bool"
	case BIN8:
		return "BIN8"
	case BIN16:
		return "BIN16"
	case BIN32:
		return "BIN32"
	case EXT8:
		return "EXT8"
	case EXT16:
		return "EXT16"
	case EXT32:
		return "EXT32"
	case FLOAT32:
		return "float32"
	case FLOAT64:
		return "float64"
	case UINT8:
		return "uint8"
	case UINT16:
		return "uint16"
	case UINT32:
		return "uint32"
	case UINT64:
		return "uint64"
	case INT8:
		return "int8"
	case INT16:
		return "int16"
	case INT32:
		return "int32"
	case INT64:
		return "int64"
	case FIXEXT1:
		return "FIXEXT1"
	case FIXEXT2:
		return "FIXEXT2"
	case FIXEXT4:
		return "FIXEXT4"
	case FIXEXT8:
		return "FIXEXT8"
	case FIXEXT16:
		return "FIXEXT16"
	case STR8, STR16, STR32:
		return "string"
	case ARRAY16, ARRAY32:
		return "array"
	case MAP16, MAP32:
		return "map"
	default:
		return "unknown"
	}
}
