package builtin_errors

type QueryErr struct {
	Msg    string
	Detail error
}

func (e *QueryErr) Error() string {
	return e.Msg
}

// Unwrap 实现 Wrapper 接口
func (e *QueryErr) Unwrap() error {
	return e.Detail
}

func NewQueryErr(msg string, detail error) *QueryErr {
	return &QueryErr{Msg: "query " + msg + " failed", Detail: detail}
}
