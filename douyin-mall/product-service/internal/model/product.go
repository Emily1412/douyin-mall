package model

import (
	"database/sql/driver"
	"encoding/json"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现sql.Scanner接口，Scan将value扫描至Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

// 商品表结构的定义
type Product struct {
	ID          int32    `gorm:"primary_key;type:int"`
	Name        string   `gorm:"type:varchar(50);not null"`
	Description string   `gorm:"type:varchar(100);not null"`
	Picture     string   `gorm:"type:varchar(200);not null"`
	Price       float32  `gorm:"not null"`
	Categories  GormList `gorm:"type:varchar(100);not null"`
}
