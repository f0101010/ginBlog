package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"ginBlog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUserExist(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 { // 默认查询第一页
		pageNum = 1
	}
	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUpUserExist(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.ResetPassword(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
