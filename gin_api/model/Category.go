package model

import (
	"Blog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increament" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 检查分类是否存在
func CheckCate(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	fmt.Println("category.ID===>", category.ID)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

// 增加分类
func AddCate(data *Category) int {
	// 结构体在函数入参时，采用指针的方式传入
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCates(pagesize, pagenum int) ([]Category, int) {
	var cates []Category
	var total int   // 分类列表总数
	err := db.Limit(pagesize).Offset((pagenum - 1)*pagesize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, 0
	}
	fmt.Println("cates===>", cates)
	return cates, total
}

// 删除分类列表
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑分类
func EditCate(id int, data *Category) int {
	var maps = make(map[string]string)
	maps["name"] = data.Name
	err := db.Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}