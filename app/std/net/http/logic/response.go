package logic

type Response struct {
	code       int
	msg        string
	statusCode *int
	isString   bool
	data       interface{}
}

type response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	StatusCode *int        `json:"-"`
	IsString   bool        `json:"-"`
	Data       interface{} `json:"data"`
}

// OptionsFunc 是一个函数类型，用于设置选项
type OptionsFunc func(*Options)

// Options 包含函数的选项
type Options struct {
	StatusCode *int
}

// WithErrorCode 是一个选项函数，用于设置错误代码
func WithErrorCode(statusCode int) OptionsFunc {
	return func(opts *Options) {
		opts.StatusCode = &statusCode
	}
}

// 响应成功

func (r *Response) Success(msg string, data interface{}, optsFuncs ...OptionsFunc) {
	opts := &Options{}
	for _, f := range optsFuncs {
		f(opts)
	}

	r.code = 1
	r.msg = msg
	r.data = data

	if opts.StatusCode != nil {
		r.statusCode = opts.StatusCode
	}
}

// 响应失败

func (r *Response) Error(msg string, data interface{}, optsFuncs ...OptionsFunc) {
	opts := &Options{}
	for _, f := range optsFuncs {
		f(opts)
	}
	r.code = 1
	r.msg = msg
	r.data = data
	if opts.StatusCode != nil {
		r.statusCode = opts.StatusCode
	}
}

func (r *Response) String(msg string, optsFuncs ...OptionsFunc) {
	opts := &Options{}
	for _, f := range optsFuncs {
		f(opts)
	}
	r.code = 1
	r.msg = msg
	if opts.StatusCode != nil {
		r.statusCode = opts.StatusCode
	}
}

func (r *Response) Response() response {
	return response{
		Code:       r.code,
		Msg:        r.msg,
		StatusCode: r.statusCode,
		IsString:   r.isString,
		Data:       r.data,
	}
}
