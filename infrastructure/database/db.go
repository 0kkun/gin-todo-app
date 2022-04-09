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

// InsertTodo Todoを保存する
func InsertTodo(db *gorm.DB, todo models.Todo) {
	// &は型を参照
	db.Create(&todo)
}

// FetchTodo 任意の個数のint型のidを受け取り、Todoを取得して返す
func FetchTodo(db *gorm.DB, id ...int) []models.Todo {
	var todos []models.Todo
	// Findはwhere句と同じ
	db.Find(&todos)
	return todos
}

// getTodoByText はテキストでTodoを検索して返す
func getTodoByText(db *gorm.DB, text string) models.Todo {
	var todo models.Todo
	db.First(&todo, "text=?", text)
	return todo
}

// updateTodo はTodoをアップデートする
func updateTodo(db *gorm.DB, id int, text string, status models.Status, deadline int) models.Todo {
	var todo models.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
	return todo
}

// DeleteTodo idを指定してTodoを削除する
func DeleteTodoById(db *gorm.DB, id int) {
	var todo models.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}