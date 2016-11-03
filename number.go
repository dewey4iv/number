package number

import (
	"strconv"
	"strings"
)

type NumberType string

const (
	Int   NumberType = "Int"
	Float NumberType = "Float"
)

func NewNumber(in interface{}) *Number {
	var n Number

	switch in.(type) {
	case int64:
		n.t = Int
		n.i = in.(int64)
		n.f = float64(n.i)
	case float64:
		n.t = Float
		n.f = in.(float64)
		n.i = int64(n.f)
	case int:
		n.t = Int
		n.i = int64(in.(int))
		n.f = float64(n.i)
	case float32:
		n.t = Float
		n.f = float64(in.(float32))
		n.i = int64(n.f)
	default:
		return nil
	}

	return &n
}

type Number struct {
	t NumberType
	i int64
	f float64
}

func (n *Number) Type() NumberType {
	return n.t
}

func (n *Number) Int() int64 {
	return n.i
}

func (n *Number) Float() float64 {
	return n.f
}

func (n *Number) String() string {
	if n.t == Int {
		return strconv.Itoa(int(n.i))
	}

	if n.t == Float {
		return strconv.FormatFloat(n.f, 'g', -1, 64)
	}

	return "null"
}

func (n *Number) MarshalJSON() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *Number) UnmarshalJSON(data []byte) error {
	var err error

	if n.i, err = strconv.ParseInt(string(data), 10, 64); err != nil {
		return err
	}

	if n.f, err = strconv.ParseFloat(string(data), 64); err != nil {
		return err
	}

	if strings.Contains(string(data), ".") {
		n.t = Float
	} else {
		n.t = Int
	}

	return nil
}
