package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//go:embed all:database/foo.db
var database embed.FS

// App struct
type App struct {
	ctx context.Context
}

type Web struct {
	Website string `json:"website"`
	Name    string `json:"name"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) WebList() []Web {
	// 存储查询结果
	var webs []Web
	db := getDb()
	// 执行查询
	rows, err := db.Query("SELECT website, name FROM weblist")
	if err != nil {
		log.Fatalf("Failed to query database: %v", err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	for rows.Next() {
		var web Web
		if err := rows.Scan(&web.Website, &web.Name); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		webs = append(webs, web) // 将结果添加到切片中
	}
	// 检查迭代是否出现错误
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}
	return webs
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func getDb() *sql.DB {
	// 创建临时文件
	tempFile, err := os.CreateTemp("", "foo.db")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name()) // 确保临时文件在退出时被删除

	// 将嵌入的数据库复制到临时文件
	dbContent, err := database.ReadFile("database/foo.db")
	if err != nil {
		log.Fatalf("Failed to read embedded database: %v", err)
	}

	if _, err := tempFile.Write(dbContent); err != nil {
		log.Fatalf("Failed to write to temp file: %v", err)
	}

	// 确保在写入后关闭文件
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Failed to close temp file: %v", err)
	}
	// 打开数据库
	db, err := sql.Open("sqlite3", tempFile.Name())
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	return db
}
