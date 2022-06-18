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
