package model

import (
	"Blog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string    `gorm:"type:varchar(100);not null" json:"title"`
	Cid int         `gorm:"type:int;not null" json:"cid"`
	Desc string     `gorm:"type:varchar(200)" json:"desc"`
	Content string  `gorm:"type:longtext" json:"content"`
	Img string      `gorm:"type:varchar(100)" json:"img"`
}

// 添加文章
func AddArticle(data *Article) int {
	// 结构体在函数入参时，采用指针的方式传入
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// todo 查询分类下的所有文章
func GetCateArticle(pagesize, pagenum, id int) ([]Article, int, int){
	var cateArtList []Article
	var total int
	err := db.Preload("Category").Limit(pagesize).Offset((pagenum - 1)*pagesize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error

	if err != nil{
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// todo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS


}

// todo 查询文章列表
func GetArticle(pagesize, pagenum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Preload("Category").Limit(pagesize).Offset((pagenum - 1)*pagesize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 删除文章
func DeleteArtic(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	fmt.Println("maps==>", maps)
	err := db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
