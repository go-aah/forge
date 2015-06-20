package forge

type ValueType int

const (
	UNKNOWN ValueType = iota

	// Primative values
	BOOLEAN
	FLOAT
	INTEGER
	NULL
	STRING

	// Complex values
	REFERENCE
	SECTION
)

var valueTypes = [...]string{
	BOOLEAN: "BOOLEAN",
	FLOAT:   "FLOAT",
	INTEGER: "INTEGER",
	NULL:    "NULL",
	STRING:  "STRING",

	REFERENCE: "REFERENCE",
	SECTION:   "SECTION",
}

func (valueType ValueType) String() string {
	str := ""
	if 0 <= valueType && valueType < ValueType(len(valueTypes)) {
		str = valueTypes[valueType]
	}

	if str == "" {
		str = "UNKNOWN"
	}

	return str
}

type Value interface {
	GetType() ValueType
	GetValue() interface{}
	UpdateValue(interface{}) error
}