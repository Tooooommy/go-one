package jsonx

const (
	Failure = iota - 1
	Success
)

// Response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// RawJSON: return raw Response
func RawJSON(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSON: return success Response
func JSON(data interface{}) *Response {
	return RawJSON(Success, "success", data)
}

// Assert: assert bool and panic Response
func Assert(bo bool, code int, msg string, ds ...interface{}) {
	if !bo {
		var data interface{}
		if len(ds) > 0 {
			data = ds[0]
		}
		panic(RawJSON(code, msg, data))
	}
}

// CheckErr: check err and panic Response
func CheckErr(err error, cs ...int) {
	if err != nil {
		var code = Failure
		if len(cs) > 0 {
			code = cs[0]
		}
		Assert(false, code, err.Error())
	}
}
