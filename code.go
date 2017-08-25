package errors

type Code interface {
	GetValue() int64
	GetName() string
}

func ParseCode(err error) (Code, bool) {
	if err == nil {
		return nil, false
	}
	if err, ok := err.(Error); ok {
		if err.Code != nil {
			return err.Code, true
		}
	}
	if err, ok := err.(*Error); ok {
		if err.Code != nil {
			return err.Code, true
		}
	}
	if err, ok := err.(causer); ok {
		return ParseCode(err.Cause())
	}
	return nil, false
}

func HasCode(err error, code Code) bool {
	c, ok := ParseCode(err)
	if !ok {
		return false
	}
	return c == code
}
