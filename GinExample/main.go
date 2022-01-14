package main

import (
	"GinExample/pkg/config"
	"GinExample/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 将请求参数绑定到struct对象
type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	// // 设置一个get请求的路由，url为/hello, 处理函数（或者叫控制器函数）是一个闭包函数。
	// r.GET("/hello", func(c *gin.Context) {
	// 	// 通过请求上下文对象Context, 直接往客户端返回一个json
	// 	c.JSON(200, gin.H{"message": "Hello World!"})
	// })

	// r.GET("user/:id", func(c *gin.Context) {
	// 	// 初始化user struct
	// 	u := User{}
	// 	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	// 	if c.ShouldBind(&u) == nil {
	// 		fmt.Println(u.Name)
	// 		fmt.Println(u.Email)
	// 	}
	// 	// http 请求返回一个字符串
	// 	c.String(200, "Success")
	// })

	// r.GET("test/gethead", GetHeadHandler)

	//注册路由
	routers.Init(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))
}

// 获取请求头信息
func GetHeadHandler(c *gin.Context) {
	//获取请求头Host的值
	ct := c.GetHeader("Content-Type")
	//控制台输出host的值
	fmt.Println(ct)
	fmt.Println(c.Request.Header.Get("User-Agent"))
	c.String(200, "Success")
}
