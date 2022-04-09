package main 

import (
    "github.com/gin-gonic/gin"
	"github.com/0kkun/gin-todo-app/infrastructure/database"
	"fmt"
	"os"
)

func main() {

	db := database.DbInit()
	fmt.Println(db)

    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(c *gin.Context){
        c.HTML(200, "index.html", gin.H{})
    })
	
    router.Run(":" + os.Getenv("PORT"))
}