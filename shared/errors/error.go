package errors

import "github.com/kiritokun07/go-zero-study/shared/errorx"

var (
	MyError          = errorx.NewCodeError(400000, "error")
	ShorturlNotExist = errorx.NewCodeError(400001, "shorturl not exist")
)
