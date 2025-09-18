package main

import "fmt"

//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

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

func SelectUserPosts(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func SelectMostCommentedPost(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Preload("Comments").Order("count(comments.id) desc").Limit(1).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	posts, err := SelectUserPosts(db, 1)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Printf("Post: %+v\n", post)
		for _, comment := range post.Comments {
			fmt.Printf("Comment: %+v\n", comment)
		}
	}
	post, err := SelectMostCommentedPost(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Most commented post: %+v\n", post)
}
