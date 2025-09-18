package main

//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
