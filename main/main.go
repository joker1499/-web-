package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gojo/docs"
	"gojo/liner"
	"gojo/middleware"
	"net/http"
	time2 "time"
)

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	r := gin.New()
	r.Use(middleware.Cors())

	r.GET("/swagger/*index", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/demo", SayHello)
		v1.POST("/file", upload)
	}
	_ = r.Run(":8080")
}

// SayHello
// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello sos"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /demo [get]
func SayHello(c *gin.Context) {
	name := c.Query("who")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"info": "who are you"})
		return
	}
	c.JSON(http.StatusOK, name)
}

// upload
// @Accept mpfd
// @Param file formData string true "随便"
// @Param other formData file true "文件"
// @Success 200 {string} string "{"msg": "hello sos"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /file [post]
func upload(c *gin.Context) {

	file := c.PostForm("file")

	c.JSON(200, file)

	f, _ := c.FormFile("other")

	time := time2.Now().Format("2006-01-02 15-04-11 ")
	fmt.Println(time + f.Filename)
	open, err := f.Open()
	if err != nil {
		return
	}
	liner.Extract(open)
	err = c.SaveUploadedFile(f, "./files/"+time+f.Filename)
	if err != nil {
		return
	}

}
