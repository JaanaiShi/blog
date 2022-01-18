package v1

import (
	"Blog/model"
	"Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var cate_code int
// 查询分类是否存在

// 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	cate_code = model.CheckCate(data.Name)
	if cate_code == errmsg.SUCCESS {
		model.AddCate(&data)
	}
	if cate_code == errmsg.ERROR_CATEGORY_USED {
		cate_code = errmsg.ERROR_CATEGORY_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code": cate_code,
		"data": data,
		"message": errmsg.GetErrMsg(cate_code),
	})
}
// todo 查询分类下的所有文章


// 查询分类列表
func GetCates(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetCates(pageSize, pageNum)
	cate_code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": cate_code,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(cate_code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	cate_code = model.CheckCate(data.Name)
	if cate_code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if cate_code == errmsg.ERROR_CATEGORY_USED {
		cate_code = errmsg.ERROR_CATEGORY_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status": cate_code,
		"message": errmsg.GetErrMsg(cate_code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	cate_code = model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status": cate_code,
		"message": errmsg.GetErrMsg(cate_code),
	})
}
