package router

import (
	"net/http"

	_ "lithum/docs" // docs is generated by Swag CLI, you have to import it.
	"lithum/handler/role"
	"lithum/handler/sd"
	"lithum/handler/user"

	"lithum/router/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/v1/login", user.Login)

	g.POST("/v1/logout", user.Logout, middleware.AuthMiddleware())
	g.GET("/v1/me", user.Me, middleware.AuthMiddleware())

	// The user handlers, requiring authentication
	userGroup := g.Group("/v1/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("", user.Create)
		userGroup.DELETE("/:id", user.Delete)
		userGroup.PUT("/:id", user.Update)
		userGroup.GET("", user.List)
		userGroup.GET("/:username", user.Get)
	}

	roleGroup := g.Group("/v1/role")
	roleGroup.Use(middleware.AuthMiddleware())
	{
		roleGroup.POST("", role.Create)
		roleGroup.DELETE("/:id", role.Delete)
		roleGroup.PUT("/:id", role.Update)
		roleGroup.GET("", role.List)
		roleGroup.GET("/:username", role.Get)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
