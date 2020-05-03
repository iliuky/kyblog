package model

// Article 文章表
type Article struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT;"`
	Title       string `gorm:"type:varchar(128);not null;default:'';"`
	PinYin      string `gorm:"type:varchar(256);not null;default:'';"`
	Content     string `gorm:"type:text;not null;"`
	ContentType string `gorm:"type:varchar(15);not null;default:'';comment:'内容类型:md,html';"`
	ArticleType int32  `gorm:"type:int;not null;default:0;"`
	Status      int32  `gorm:"type:int;not null;default:0;comment:'状态: 0 正常, 1 删除 2 草稿箱'"`
	CreateTime  int32  `gorm:"type:int;not null;default:0;column:ctime"`
	ModifyTime  int32  `gorm:"type:int;not null;default:0;"`
}

// ArticleType 文章分类表
type ArticleType struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT;"`
	Name       string `gorm:"type:varchar(16);not null;default:'';"`
	Status     int32  `gorm:"type:int;not null;default:0;comment:'状态: 0 正常, 1 删除'"`
	CreateTime int32  `gorm:"type:int;not null;default:0;column:ctime"`
	ModifyTime int32  `gorm:"type:int;not null;default:0;"`
}
