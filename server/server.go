package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/weasel/schema"
	"io/ioutil"
)

type root struct{}

func (_ *root) Hello() string { return "Hello, world!" }

func Run() {
	// 设置 gin 的模式（调试模式：DebugMode, 发行模式：ReleaseMode）
	gin.SetMode(gin.DebugMode)
	// 创建一个不包含中间件的路由器
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/graphql", graphiql)

	//通过 graphql 字符串 & query 对象构建出schema对象
	s1 := graphql.MustParseSchema(schema.String(), &root{})

	r.POST("/api", func(c *gin.Context) {
		var opts RequestOptions
		body, err := ioutil.ReadAll(c.Request.Body)

		err = json.Unmarshal(body, &opts)
		if err != nil {
			// Probably `variables` was sent as a string instead of an object.
			// So, we try to be polite and try to parse that as a JSON string
			var optsCompatible requestOptionsCompatibility
			json.Unmarshal(body, &optsCompatible)
			json.Unmarshal([]byte(optsCompatible.Variables), &opts.Variables)
		}

		fmt.Println(opts)
		fmt.Println(s1)
		result := s1.Exec(c, opts.Query, opts.OperationName, opts.Variables)
		c.JSON(200, result)
	})
	// 在8081 端口，启动http服务
	//log.Fatal(http.ListenAndServe(":8081", nil))

	r.Run()

}
