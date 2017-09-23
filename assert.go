package errors

type AssertError struct {
	err error
}

func AssertNil(err error, msg ...string) {
	if err == nil {
		return
	}
	if len(msg) > 0 {
		err = Wrap(err, msg[0])
	}
	panic(AssertError{
		err: err,
	})
}

func Recover(errVar *error) {
	if r := recover(); r != nil {
		if err, ok := r.(AssertError); ok {
			*errVar = err.err
			return
		}
		panic(r)
	}
}
