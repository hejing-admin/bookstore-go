package model

import "time"

type Book struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Price      int       `json:"price"`
	Discount   int       `json:"discount"`
	Type       string    `json:"type"`
	Stock      int       `json:"stock"`    // 库存
	Status     int       `json:"status"`   // 上架1 下架0
	Describe   string    `json:"describe"` // 描述
	CoverUrl   string    `json:"cover_url"`
	ISBN       string    `json:"isbn"`
	Publisher  string    `json:"publisher"` // 出版社
	Pages      int       `json:"pages"`     // 页数
	Language   string    `json:"language"`  // 语言
	Format     string    `json:"format"`
	CategoryID int       `json:"category_id"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
