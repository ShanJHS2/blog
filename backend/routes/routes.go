package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 自定义 CORS 配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3002"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // 允许的请求头，包括 Authorization
		AllowCredentials: true, // 允许携带凭证（如 cookies 或授权头）
	}))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})	

	// 公共路由，不需要验证
	r.POST("/auth/login", controllers.Login)       // 登录
	r.POST("/auth/register", controllers.Register) // 注册

	// 需要验证的路由
	// 用户更新自己的信息，权限: 1 (普通用户及以上)
	r.POST("/auth/updateUser", middlewares.Auth(1), controllers.UpdateUser)

	// 文章相关路由，权限: 0 (所有用户)
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/count", controllers.GetArticleCount)
	r.GET("/articles/:id", controllers.GetArticleById)
	r.POST("/articles", middlewares.Auth(3), controllers.CreateArticle)   // 创建文章需要管理员权限
	r.DELETE("/articles", middlewares.Auth(3), controllers.DeleteArticle) // 删除文章需要管理员权限

	// 评论相关路由，权限: 0 (所有用户)
	r.GET("/comments/:blogID", controllers.GetComments)
	r.POST("/comments", controllers.CreateComment)
	r.DELETE("/comments/:blogID", middlewares.Auth(3), controllers.DeleteComment) // 删除评论需要管理员权限

	// 媒体相关路由，权限: 3 (管理员)
	r.GET("/media", controllers.GetMedia)
	r.POST("/media", middlewares.Auth(3), controllers.CreateMedia)
	r.PUT("/media/:mediaId", middlewares.Auth(3), controllers.UpdateMedia)
	r.DELETE("/media/:mediaId", middlewares.Auth(3), controllers.DeleteMedia) // 删除媒体需要管理员权限

	// 问题相关路由，权限: 0 (所有用户)
	r.GET("/questions", controllers.GetQuestions)
	r.POST("/questions", controllers.CreateQuestion)
	r.POST("/questions/:questionId/answer", middlewares.Auth(3), controllers.AnswerQuestion) // 回答问题需要编辑者权限

	return r
}
