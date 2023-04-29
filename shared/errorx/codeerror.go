package errorx

type (
	CodeError interface {
		error
		Status() int
		Code() int
		Desc() string
	}

	codeError struct {
		status int    //http-status-code
		code   int    //biz-code
		desc   string //biz-desc
	}
)

func (c *codeError) Error() string {
	return c.desc
}

func (c *codeError) Status() int {
	return c.status
}

func (c *codeError) Code() int {
	return c.code
}

func (c *codeError) Desc() string {
	return c.desc
}

func NewCodeError(code int, desc string) CodeError {
	return newError(400, code, desc)
}

func newError(status, code int, desc string) CodeError {
	return &codeError{
		status: status,
		code:   code,
		desc:   desc,
	}
}
