package main 

import (
    "github.com/gin-gonic/gin"
	"github.com/0kkun/gin-todo-app/infrastructure/database"
)

func main() {

	db := database.DbInit()
	defer db.Close()

    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(c *gin.Context){
        c.HTML(200, "index.html", gin.H{})
    })

    router.Run(":3000")
}