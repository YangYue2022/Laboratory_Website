package routers

import (
	_ "UISwebsite_backend/docs"
	"UISwebsite_backend/middleware/jwt"
	"UISwebsite_backend/pkg/setting"
	"UISwebsite_backend/pkg/upload"
	"UISwebsite_backend/routers/api"
	"UISwebsite_backend/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //  "*" 允许所有域
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// 是否使用swagger
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiPublic := r.Group("/api")
	{
		apiPublic.GET("/auth", api.GetAuth)
		apiPublic.POST("/upload", api.UploadImage)

		//获取成果列表
		apiPublic.GET("/works", v1.GetWorks)
		//获取奖励列表
		apiPublic.GET("/rewards", v1.GetRewards)
		//获取QA列表
		apiPublic.GET("/qas", v1.GetQas)
		//获取成员列表
		apiPublic.GET("/members", v1.GetMembers)
		//获取记录列表
		apiPublic.GET("/records", v1.GetRecords)
	}

	apiAdmin := r.Group("/api/admin")
	apiAdmin.Use(jwt.JWT())
	{
		//新建成果
		apiAdmin.POST("/work", v1.AddWork)
		//更新成果标签
		apiAdmin.PUT("/work/:id", v1.UpdateWork)
		//删除成果标签
		apiAdmin.DELETE("/work/:id", v1.DeleteWork)

		//新建奖励
		apiAdmin.POST("/reward", v1.AddReward)
		//更新奖励
		apiAdmin.PUT("/reward/:id", v1.UpdateReward)
		//删除奖励
		apiAdmin.DELETE("/reward/:id", v1.DeleteReward)

		//新建QA
		apiAdmin.POST("/qa", v1.AddQa)
		//更新QA
		apiAdmin.PUT("/qa/:id", v1.EditQa)
		//删除QA
		apiAdmin.DELETE("/qa/:id", v1.DeleteQa)

		apiAdmin.POST("/member", v1.AddMember)
		apiAdmin.PUT("/member/:id", v1.UpdateMember)
		apiAdmin.DELETE("/member/:id", v1.DeleteMember)

		apiAdmin.POST("/record", v1.AddRecord)
		apiAdmin.PUT("/record/visibility", v1.UpdateRecordVisibility)
		apiAdmin.PUT("/record/:id", v1.UpdateRecord)
		apiAdmin.DELETE("/record/:id", v1.DeleteRecord)
	}
	return r
}
