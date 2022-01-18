package routes

import (
	v1 "Blog/api/v1"
	"Blog/middleware"
	"Blog/utils"
	"github.com/gin-gonic/gin"
)

/*
	函数名称大写就是公有方法，其他包也可以引用
	函数名称小写就是私有方法，只能在自己的包内使用
*/
func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// 路由组
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口


		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("cate/add",v1.AddCate)

		auth.PUT("cate/:id", v1.EditCate)
		auth.DELETE("cate/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.POST("article/add",v1.AddArticle)

		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("/api/v1")
	{
		router.POST("user/add",v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("cates", v1.GetCates)
		router.GET("articles", v1.GetArticle)
		router.GET("article/list/:cid", v1.GetCateAticle)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)

}
