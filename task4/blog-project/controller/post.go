package controller

import (
	"blog-main/global"
	"blog-main/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestPost struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

func CreatePost(c *gin.Context) {
	var req requestPost
	userId, err := c.Get("userId")
	if !err {
		global.ReturnJson(c, http.StatusUnauthorized, -1, "请先登录", nil)
		return
	}
	if err := c.BindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "参数错误", nil)
		return
	}
	db := global.GetDB()
	post := model.Post{
		Content: req.Content,
		Title:   req.Title,
		UserID:  userId.(uint),
	}
	if err := db.Create(&post).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "创建失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "创建成功", nil)
}

func GetPosts(c *gin.Context) {
	db := global.GetDB()
	var posts []model.Post
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "查询失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "查询成功", posts)
}

func GetPost(c *gin.Context) {
	db := global.GetDB()
	var post model.Post
	id := c.Param("id")
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "查询失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "查询成功", post)
}

func DeletePost(c *gin.Context) {
	db := global.GetDB()
	var post model.Post
	id := c.Param("id")
	if err := db.First(&post, id).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "查询失败", nil)
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "删除失败", nil)
	}
}

func UpdatePost(c *gin.Context) {
	db := global.GetDB()
	var req requestPost
	var post model.Post
	id := c.Param("id")
	if err := db.First(&post, id).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "查询失败", nil)
		return
	}
	if err := c.BindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "参数错误", nil)
		return
	}
	post.Content = req.Content
	post.Title = req.Title
	if err := db.Save(&post).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "更新失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "更新成功", nil)
}
