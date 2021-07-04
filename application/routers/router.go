package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/togettoyou/blockchain-real-estate/application/routers/api/v1"
	"net/http"
	"strings"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	//swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/QueryUserTest", v1.QueryUserTest)
		apiV1.GET("/QueryUserTestGet", v1.QueryUserTestGet)
		apiV1.POST("/QueryOrderByIDTest", v1.QueryOrderByIDTest)

		apiV1.POST("/StoreOrder", v1.StoreOrder)
		apiV1.POST("/QueryOrder", v1.QueryOrder)
		apiV1.POST("/QueryOrderByID", v1.QueryOrderByID)
		apiV1.POST("/QueryOrderByBuyer", v1.QueryOrderByBuyer)
		apiV1.POST("/QueryOrderBySeller", v1.QueryOrderBySeller)
		apiV1.POST("/UpdateStatus", v1.UpdateStatus)
		apiV1.POST("/UpdateTransferStatus", v1.UpdateTransferStatus)
		apiV1.POST("/DeleteOrder", v1.DeleteOrder)

		apiV1.POST("/StoreProduct", v1.StoreProduct)
		apiV1.POST("/UpdateProductByAmount", v1.UpdateProductByAmount)
		apiV1.POST("/QueryProduct", v1.QueryProduct)
		apiV1.POST("/QueryProductByID", v1.QueryProductByID)
		apiV1.POST("/QueryProductByGoodType", v1.QueryProductByGoodType)
		apiV1.POST("/QueryProductByMaterial", v1.QueryProductByMaterial)
		apiV1.POST("/queryDonatingListByGrantee", v1.QueryProductByWorkmanship)
		apiV1.POST("/QueryProductBySupplier", v1.QueryProductBySupplier)
		apiV1.POST("/DeleteProduct", v1.DeleteProduct)

		apiV1.POST("/StoreFinance", v1.StoreFinance)
		apiV1.POST("/UpdateFinanceStatus", v1.UpdateFinanceStatus)
		apiV1.POST("/QueryFinance", v1.QueryFinance)
		apiV1.POST("/QueryFinanceBySupplier", v1.QueryFinanceBySupplier)
		apiV1.POST("/QueryFinanceByStatus", v1.QueryFinanceByStatus)
		apiV1.POST("/QueryFinanceByOrderID", v1.QueryFinanceByOrderID)

		apiV1.POST("/StoreUser", v1.StoreUser)
		apiV1.POST("/QueryUser", v1.QueryUser)
	}
	// 静态文件路由
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}

// Cors 允许跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
