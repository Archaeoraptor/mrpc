package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service Method"
	Seq           uint64 // sequence number choosen by client
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const GobType Type = "application/gob"
const JsonType Type = "application/json"

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
