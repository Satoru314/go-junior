package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"` // エラーの詳細をJSONに含めない
}

func (e *MyAppError) Error() string {
	return e.Err.Error()
}

func (e *MyAppError) Unwrap() error {
	return e.Err
}
