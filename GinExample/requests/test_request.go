package requests

type TestRequest struct {
	//测试请求结构体 该结构体定义了请求的参数和校验规则
	Username string `form:"username" binding:"required"`
}
