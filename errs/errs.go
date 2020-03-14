package errs

import (
	"fmt"
	"io"
)

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	code  int         // external code
	msg   string      // external message
	data  interface{} // external data
	stack *stack      // external stack
}

// Error error interface function
func (f *fundamental) Error() string {
	return f.msg
}

// Format error interface function
func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, f.msg)
			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, f.msg)
	case 'q':
		fmt.Fprintf(s, "%q", f.msg)
	}
}

// New returns an error with the supplied message
func New(msg string) error {
	return &fundamental{
		msg:   msg,
		stack: callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
func Errorf(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// NewWithOption returns an error with some option
func NewWithOption(opt ...Option) error {
	res := &fundamental{
		stack: callers(),
	}
	for _, o := range opt {
		o(res)
	}
	return res
}

// Wrap return an error annotating err with a stack trace
func Wrap(err error, msg string) error {
	return &fundamental{
		code:  GetCode(err),
		msg:   fmt.Sprintf("%s: %s", msg, err.Error()),
		data:  GetData(err),
		stack: callers(),
	}
}

// Option error option
type Option func(*fundamental)

// OptCode set error code
func OptCode(code int) Option {
	return func(o *fundamental) {
		o.code = code
	}
}

// GetCode get error option code
func GetCode(e error) int {
	if f, ok := e.(*fundamental); ok {
		return f.code
	}
	return 0
}

// OptMsg set error msg
func OptMsg(msg string) Option {
	return func(o *fundamental) {
		o.msg = msg
	}
}

// OptData set error data
func OptData(data interface{}) Option {
	return func(o *fundamental) {
		o.data = data
	}
}

// GetData get error option data
func GetData(e error) interface{} {
	if f, ok := e.(*fundamental); ok {
		return f.data
	}
	return nil
}

// GetStack get error stack
func GetStack(e error) string {
	if f, ok := e.(*fundamental); ok {
		return f.stack.Output()
	}
	return ""
}
