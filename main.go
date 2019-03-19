package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/danceyoung/trycatchserver/constant"
	"github.com/danceyoung/trycatchserver/net"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(cors.Default())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		defer fmt.Println("defer println")
		defer fmt.Println("defer2 println")
		fmt.Println("normal println")
		var i = 1
		if i == 1 {
			panic("this is panic")
		}
		defer fmt.Println("println after panic")
		c.String(200, "pong")
	})

	user := r.Group("/user")
	user.POST("/signin", net.Signin)
	user.POST("/signinfrommobile", net.SigninFromMobile)
	user.POST("/profile", net.Profile)
	user.POST("/changepassword", net.ChangePassword)

	project := r.Group("/project")
	project.POST("/new", net.NewProject)
	project.POST("/save", net.SaveProject)
	project.POST("/delete", net.DeleteProject)
	project.POST("/detail", net.ProjectDetail)
	project.POST("/list", net.Projects)
	project.POST("/receivefromlist", net.ReceiveFromList)
	project.POST("bugs", net.Bugs)

	diaper := r.Group("/try")
	diaper.POST("/catchinfo", net.TryCatch)

	apns := r.Group("/apns")
	apns.POST("/devicetoken", net.DeviceToken)
	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := DB[user]
	// 	if ok {
	// 		c.JSON(200, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(200, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		// user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			// DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}

func setupLogger() {
	file, _ := os.Create("./log" + time.Now().Format("2006-01-02 15-04-05"))
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	if constant.DEBUG == true {
		fmt.Println("DEBUG Model")
	}
	// gin.SetMode(gin.ReleaseMode)
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
	// r.RunTLS(":8000", "/Users/young/young/Biz/TryCatch/keycer/server.pem", "/Users/young/young/Biz/TryCatch/keycer/server.key")
}
