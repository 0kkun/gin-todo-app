package database

import (
	"os"
	"log"
	"github.com/0kkun/gin-todo-app/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DbInit はデータベースのDBを接続してマイグレーションを行いDBを取得する
func DbInit() *gorm.DB {
	mysqlConnection := mysql.Open(os.Getenv("DSN"))
	db, err := gorm.Open(mysqlConnection, &gorm.Config{})

	if err != nil {
		log.Fatalf("DB Connectin ERROR = %v \n", err)
		panic(err)
	}

	log.Println("DB Connection Success")
	db.AutoMigrate(&models.Todo{})
	return db
}


// func CloseDB(db *gorm.DB) {
// 	if err := db.Close(); err != nil {
// 		log.Fatal("CloseDB failed:", err)
// 	}
// }

// Todoのマイグレート
// func DbInit() *gorm.DB {
// 	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/gin_todo_app?parseTime=true")
// 	if err != nil {
// 		fmt.Errorf("could not open database")
// 	}
// 	fmt.Println("DB Connection Success")
// 	db.AutoMigrate(&models.Todo{})
// 	return db
// }

// // DBの作成
// func DbCreate(todo models.Todo) {
// 	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
// 	if err != nil {
// 		fmt.Errorf("could not open database")
// 	}
// 	db.Create(&todo)
// }

// func DbRead(id ...int) []models.Todo {
// 	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
// 	if err != nil {
// 		fmt.Errorf("could not open database")
// 	}
// 	var todos []models.Todo
// 	db.Find(&todos)
// 	return todos
// }

// func DbUpdate(id int, text string, status models.Status, deadline int) models.Todo {
// 	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
// 	if err != nil {
// 		fmt.Errorf("could not open database")
// 	}
// 	var todo models.Todo
// 	db.First(&todo, id)
// 	todo.Text = text
// 	todo.Status = status
// 	todo.Deadline = deadline
// 	db.Save(&todo)
// 	return todo
// }

// func DbDelete(id int) {
// 	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
// 	if err != nil {
// 		fmt.Errorf("could not open database")
// 	}
//     var todo models.Todo
//     db.First(&todo, id)
//     db.Delete(&todo)
// }