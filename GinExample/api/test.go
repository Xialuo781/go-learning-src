package api

import (
	"GinExample/requests"
	"GinExample/services/hello"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	// 实例化一个TestRequest结构体，用于接收参数
	testStruct := requests.TestRequest{}

	// 接收请求参数
	err := c.ShouldBind(&testStruct)

	// 判断参数校验是否通过，如果不通过，把错误返回给前端
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": requests.Translate(err)})
		return
	}

	// 调用HelloService
	var service hello.HelloContract

	// 这里使用的是接口定义了变量
	// service = &hello.HelloService{}
	service = &hello.GreeterService{}
	// 调用服务的方法处理业务
	// result := service.SayHello(testStruct.Username)
	result := service.SayHello(testStruct.Username)

	// 校验通过，返回请求参数
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// func GetHandler(c *gin.Context) {
// 	//获取name参数, 通过Query获取的参数值是String类型。
// 	name := c.Query("name")

// 	//获取name参数, 跟Query函数的区别是，可以通过第二个参数设置默认值。
// 	name = c.DefaultQuery("name", "tizi365")
// 	fmt.Println("GetHandler name:", name)

// 	//获取id参数, 通过GetQuery获取的参数值也是String类型,
// 	// 区别是GetQuery返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在。
// 	id, ok := c.GetQuery("id")
// 	if !ok {
// 		// 参数不存在
// 	}
// 	fmt.Println("GetHandler id:", id)
// }

// func PostHandler(c *gin.Context) {
// 	//获取name参数, 通过PostForm获取的参数值是String类型。
// 	name := c.PostForm("name")

// 	// 跟PostForm的区别是可以通过第二个参数设置参数默认值
// 	name = c.DefaultPostForm("name", "tizi365")
// 	fmt.Println("PostHandler name:", name)

// 	//获取id参数, 通过GetPostForm获取的参数值也是String类型,
// 	// 区别是GetPostForm返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在。
// 	id, ok := c.GetPostForm("id")
// 	if !ok {
// 		// 参数不存在
// 	}
// 	fmt.Println("PostHandler id:", id)
// }
