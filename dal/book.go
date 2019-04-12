package dal

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookName string `gorm:"column:bookname"`
	PicUrl string `gorm:"column:book"`
	Author string `gorm:"column:author"`
	Category string `gorm:"column:category"`
	Price float32 `gorm:"column:price"`
	Score float32 `gorm:"column:score"`
	ReviewCount int `gorm:"column:review_count"`
	Review string `gorm:"column:reviews"`
	Publisher string `gorm:"column:publisher"`
	Content string `gorm:"column:content"`
	Extra string `gorm:"column:extra"`
}

