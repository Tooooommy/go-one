package server

const (
	Failure = iota
	Success
)

// JSONResponse
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// RawJSON: return raw JSONResponse
func RawJSON(code int, msg string, data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSON: return success JSONResponse
func JSON(data interface{}) *JSONResponse {
	return RawJSON(Success, "success", data)
}

// Assert: assert bool and panic
func Assert(bo bool, code int, msg string, ds ...interface{}) {
	if !bo {
		var data interface{}
		if len(ds) > 0 {
			data = ds[0]
		}
		panic(RawJSON(code, msg, data))
	}
}

// CheckErr: check err and panic
func CheckErr(err error, cs ...int) {
	if err != nil {
		var code = Failure
		if len(cs) > 0 {
			code = cs[0]
		}
		Assert(false, code, err.Error())
	}
}
