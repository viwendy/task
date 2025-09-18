package main

//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

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

func (u *User) AfterCreate(db *gorm.DB) error {
	// 更新用户的文章数量统计字段
	err := db.Model(&User{}).Where("id = ?", u.ID).Update("post_count", gorm.Expr("post_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) AfterDelete(db *gorm.DB) error {
	// 检查文章的评论数量
	var commentCount int64
	err := db.Model(&Comment{}).Where("post_id = ?", p.ID).Count(&commentCount).Error
	if err != nil {
		return err
	}
	// 如果评论数量为 0，则更新文章的评论状态为 "无评论"
	if commentCount == 0 {
		err := db.Model(&Post{}).Where("id = ?", p.ID).Update("comment_status", "无评论").Error
		if err != nil {
			return err
		}
	}
	return nil
}
