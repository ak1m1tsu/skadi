package skadi

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CommandCode uint16

const (
	Message CommandCode = iota + 1
	Close
)

type Request struct {
	CommandCode CommandCode
	Body        []byte
}

func (r Request) String() string {
	return fmt.Sprintf("CommandCode: %d, Body: %s", r.CommandCode, r.Body)
}

type Response struct {
	Body []byte
}

func (r Response) String() string {
	return fmt.Sprintf("Body: %s", r.Body)
}

func Encode(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	switch t := data.(type) {
	case *Request:
		if err := binary.Write(&buffer, binary.BigEndian, uint16(t.CommandCode)); err != nil {
			return nil, err
		}
		if err := binary.Write(&buffer, binary.BigEndian, t.Body); err != nil {
			return nil, err
		}
	case *Response:
		if err := binary.Write(&buffer, binary.BigEndian, t.Body); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported data type: %T", data)
	}
	return buffer.Bytes(), nil
}

func Decode(data []byte, dest interface{}) error {
	reader := bytes.NewReader(data)
	switch t := dest.(type) {
	case *Request:
		if err := binary.Read(reader, binary.BigEndian, &t.CommandCode); err != nil {
			return err
		}
		t.Body = make([]byte, reader.Len())
		if err := binary.Read(reader, binary.BigEndian, t.Body); err != nil {
			return err
		}
	case *Response:
		t.Body = make([]byte, reader.Len())
		if err := binary.Read(reader, binary.BigEndian, t.Body); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported data type: %T", dest)
	}
	return nil
}
