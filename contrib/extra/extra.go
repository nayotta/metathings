package extra

import (
	"errors"
	"regexp"
)

const (
	OPERATOR_REGEX = `^\$(?P<opcode>[\w\d]+)([\.](?P<operand>.+))?$`
)

var (
	ErrUnsupportedOperator = func(op string) error {
		return errors.New("unsupported operator: " + op)
	}
)

type Extra map[string]string

func New() Extra {
	return make(Extra)
}

func FromRaw(raw map[string]string) Extra {
	return Extra(raw)
}

func (e Extra) Get(key string) (val string, ok bool) {
	val, ok = e[key]
	return
}

func (e Extra) Set(key, val string) Extra {
	e[key] = val
	return e
}

func (e Extra) Unset(key string) Extra {
	delete(e, key)
	return e
}

func (e Extra) Raw() map[string]string {
	return map[string]string(e)
}

func (e Extra) Copy() (y Extra) {
	y = New()

	for k, v := range e.Raw() {
		y.Set(k, v)
	}

	return y
}

func (e Extra) Apply(cfg map[string]string) (y Extra, err error) {
	y = e.Copy()

	for key, val := range cfg {
		opcode, operand := parse_operator(key)
		switch opcode {
		case "set":
			y.Set(operand, val)
		case "unset":
			y.Unset(operand)
		case "clear":
			return New(), nil
		default:
			return nil, ErrUnsupportedOperator(key)
		}
	}

	return y, nil
}

func parse_operator(x string) (string, string) {
	var opcode, operand string
	re := regexp.MustCompile(OPERATOR_REGEX)
	matchs := re.FindStringSubmatch(x)
	names := re.SubexpNames()

	if len(matchs) == 0 {
		return "", ""
	}

	for i, name := range names {
		switch name {
		case "opcode":
			opcode = matchs[i]
		case "operand":
			operand = matchs[i]
		}
	}

	return opcode, operand
}
