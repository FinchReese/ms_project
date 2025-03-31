package common

type BusinessCode int

type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string `json:"msg"`
	Data any `json:"data"`
}

func (result *Result) Success(data any) *Result {
	result.Code = 200
	result.Msg = "success"
	result.Data = data
	return result
}

func (result *Result) Fail(code BusinessCode, msg string) *Result {
	result.Code = code
	result.Msg = msg
	return result
}
