package v1

import (
	"Blog/model"
	"Blog/utils/errmsg"
	"Blog/utils/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 查询用户是否存在
func UserExist(c *gin.Context){

}

// 添加用户
func AddUser(c *gin.Context)  {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	// 数据验证
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context)  {
	fmt.Println("123")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	fmt.Println(pageSize, pageNum)
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(code),
	})

}
// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ :=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()    // 阻止执行
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code  = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}