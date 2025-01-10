package helper

type ResponseValidator struct {
	StatusCode int
	Response   string
}

type ResponseResult struct {
	Success   bool
	SQLErrors bool
}
