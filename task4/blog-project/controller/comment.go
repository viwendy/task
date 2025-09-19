package controller

import (
	"blog-main/global"
	"blog-main/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type requestComment struct {
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}

func CreateComment(c *gin.Context) {
	var req requestComment
	userId, exists := c.Get("userId")
	if !exists {
		global.ReturnJson(c, http.StatusBadRequest, -1, "请登录", nil)
		return
	}
	if err := c.BindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "参数错误", nil)
		return
	}
	db := global.GetDB()
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "参数错误", nil)
		return
	}
	var comment = model.Comment{
		Content: req.Content,
		PostID:  uint(postID),
		UserID:  userId.(uint),
	}
	if err := db.Create(&comment).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "创建失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "创建成功", nil)
}

func DeleteComment(c *gin.Context) {
	var comment model.Comment
	var db = global.GetDB()
	if err := db.First(&comment, c.Param("id")).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "删除失败", nil)
		return
	}
	if err := db.Delete(&comment).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "删除失败", nil)
		return
	}
}

func UpdateComment(c *gin.Context) {
	var comment model.Comment
	var req requestComment
	var db = global.GetDB()
	if err := db.First(&comment, c.Param("id")).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "更新失败", nil)
		return
	}
	if err := c.BindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "参数错误", nil)
	}
}

func GetCommentsByPost(c *gin.Context) {
	var comments []model.Comment
	var db = global.GetDB()
	if err := db.Where("post_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "获取失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "获取成功", comments)
}

func GetComments(c *gin.Context) {
	var comments []model.Comment
	var db = global.GetDB()
	if err := db.Find(&comments).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "获取失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "获取成功", comments)
}

func GetComment(c *gin.Context) {
	var comment model.Comment
	var db = global.GetDB()
	if err := db.First(&comment, c.Param("id")).Error; err != nil {
		global.ReturnJson(c, http.StatusOK, -1, "获取失败", nil)
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "获取成功", comment)
}
