package types

type CustomEnum int

const (
	CustomEnumSTRING CustomEnum = iota
	CustomEnumBYTES
	customEnumMAX
)

func SanitizeCustomEnumValue(value int) {
	if !(value >= 0 && value < int(customEnumMAX)) {
		panic("Invalid value for enum 'CustomEnum'")
	}
}

func GetCustomEnumValue(key string) CustomEnum {
	switch key {
	case "STRING":
		return CustomEnumSTRING
	case "BYTES":
		return CustomEnumBYTES
	default:
		panic("Invalid key for enum 'CustomEnum'")
	}
}

func GetCustomEnumKey(value CustomEnum) string {
	SanitizeCustomEnumValue(int(value))
	switch value {
	case CustomEnumSTRING:
		return "STRING"
	case CustomEnumBYTES:
		return "BYTES"
	default:
		panic("Invalid value for enum 'CustomEnum'")
	}
}
