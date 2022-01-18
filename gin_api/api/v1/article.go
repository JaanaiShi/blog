package v1

import (
	"Blog/model"
	"Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
var articleCode int

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)
	articleCode = model.AddArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"data": data,
		"message": errmsg.GetErrMsg(articleCode),
	})
}
// 查询分类下的所有文章
func GetCateAticle(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, articleCode, total := model.GetCateArticle(pageSize, pageNum, cid)
	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(articleCode),
	})
}

// todo 查询单个文章
func GetArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data, articleCode  := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"data": data,
		"message": errmsg.GetErrMsg(articleCode),
	})
}

// todo 查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, articleCode, total := model.GetArticle(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"data": data,
		"total": total,
		"message": errmsg.GetErrMsg(articleCode),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	articleCode = model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"message": errmsg.GetErrMsg(articleCode),
	})
}
// 删除文章

func DeleteArticle(c *gin.Context) {
	id, _ :=strconv.Atoi(c.Param("id"))
	articleCode = model.DeleteArtic(id)

	c.JSON(http.StatusOK, gin.H{
		"status": articleCode,
		"message": errmsg.GetErrMsg(articleCode),
	})
}
