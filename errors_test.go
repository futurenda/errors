package errors_test

import (
	"github.com/futurenda/errors"

	goErrors "errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Code uint32

func (c Code) GetValue() int64 {
	return int64(c)
}

func (c Code) GetName() string {
	return fmt.Sprintf("%d", c)
}

var (
	NoRows Code = 1
)

var _ = Describe("Errors", func() {
	It("Should be able to new error", func() {
		err := errors.New(NoRows, "No Rows")
		Expect(err.Error()).To(Equal("No Rows"))
		Expect(err.Code).To(Equal(NoRows))
		Expect(err.Cause()).To(BeNil())
	})
	It("Should be able to Wrap() and Cause()", func() {
		err0 := errors.New(NoRows, "No Rows")
		Expect(err0.Cause()).To(BeNil())
		err1 := errors.Wrap(err0, "wrap 1")
		Expect(err1.Cause()).To(Equal(err0))
		Expect(err1.Error()).To(Equal("wrap 1\n\tNo Rows"))
		err2 := errors.Wrap(err1, "wrap 2")
		Expect(err2.Cause()).To(Equal(err1))
		Expect(err2.Error()).To(Equal("wrap 2\n\twrap 1\n\tNo Rows"))
	})
	It("Should be able to parse code", func() {
		err0 := errors.New(NoRows, "No Rows")
		err1 := errors.Wrap(err0, "wrap 1")
		err2 := errors.Wrap(err1, "wrap 2")
		code, ok := errors.ParseCode(err2)
		Expect(ok).To(BeTrue())
		Expect(code).To(Equal(NoRows))
		code, ok = errors.ParseCode(err1)
		Expect(ok).To(BeTrue())
		Expect(code).To(Equal(NoRows))
		code, ok = errors.ParseCode(err0)
		Expect(ok).To(BeTrue())
		Expect(code).To(Equal(NoRows))
		err4 := goErrors.New("test")
		code, ok = errors.ParseCode(err4)
		Expect(ok).To(BeFalse())
		Expect(code).To(BeNil())
	})
	It("Should be able to assertNil and recover", func() {
		e1 := errors.New(nil, "E1")
		e2 := errors.New(nil, "E2")
		copy := func(i int, err error) (int, error) {
			return i, err
		}
		i, err := func() (i int, err error) {
			defer errors.Recover(&err)
			i, err = copy(1, nil)
			errors.AssertNil(err)
			i, err = copy(2, nil)
			errors.AssertNil(err, "msg")
			i, err = copy(3, e1)
			errors.AssertNil(err)
			return 4, e2
		}()
		Expect(err).To(Equal(e1))
		Expect(i).To(Equal(3))
	})
	It("Should be able to assertNil with wrap msg", func() {
		e1 := errors.New(nil, "E1")
		e2 := errors.New(nil, "E2")
		exec := false
		err := func() (err error) {
			defer errors.Recover(&err)
			errors.AssertNil(e1, "msg3")
			exec = true
			return e2
		}()
		Expect(err.(errors.Error).Cause()).To(Equal(e1))
		Expect(err.(errors.Error).Message).To(Equal("msg3"))
		Expect(exec).To(Equal(false))
	})
})
