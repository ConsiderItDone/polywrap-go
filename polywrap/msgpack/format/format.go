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
)
