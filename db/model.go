package db

type Hello struct {
	Id         int64  `gorm:"autoIncrement; primaryKey; comment:主键id"`
	Name       string `gorm:"size:15; comment:名称"`
	CreateTime int64  `gorm:"column:createTime; comment:创建时间"`
}

func (Hello) TableName() string {
	return "hello"
}
