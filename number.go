package number

import (
	"strconv"
	"strings"
)

type NType string

const (
	Int   NType = "Int"
	Float NType = "Float"
)

func New(in interface{}) *N {
	var n N

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

type N struct {
	t NType
	i int64
	f float64
}

func (n *N) Type() NType {
	return n.t
}

func (n *N) Int() int64 {
	return n.i
}

func (n *N) Float() float64 {
	return n.f
}

func (n *N) String() string {
	if n.t == Int {
		return strconv.Itoa(int(n.i))
	}

	if n.t == Float {
		return strconv.FormatFloat(n.f, 'g', -1, 64)
	}

	return "null"
}

func (n *N) MarshalJSON() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *N) UnmarshalJSON(data []byte) error {
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
