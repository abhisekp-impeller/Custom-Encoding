package custom_encoding

type ICustomEncoding interface {
	Encode(interface{}) string
	Decode(string) string
	SetCharset(string) *CustomEncoding
}

type CustomEncoding struct {
	charset string
}

func (ce *CustomEncoding) SetCharset(charset string) *CustomEncoding {
	ce.charset = charset
	return ce
}

func (ce *CustomEncoding) Encode(value interface{}) string {
	encoded := ""

	switch value.(type) {
	case uint64:
		encoded = Convert(ce.charset, value.(uint64))
		break
	}

	return encoded
}
