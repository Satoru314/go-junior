package services

import "database/sql"

// MyAppService はサービス層の構造体
type MyAppService struct {
	db *sql.DB
}

// 外部からMyAppServiceを作成
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
