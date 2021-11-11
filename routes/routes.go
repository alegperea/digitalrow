package routes

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/agp/push-notifications-api/controllers"
	"github.com/agp/push-notifications-api/docs"
	middleware "github.com/agp/push-notifications-api/middlewares"
	"github.com/agp/push-notifications-api/services"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Fila Digital API"
	docs.SwaggerInfo.Description = "Documentación de API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "api.filadigital.agp.com"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	port := os.Getenv("PORT")
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	sess := services.ConnectAws()

	r.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	if port == "" {
		port = "3001"
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/company", controllers.CreateCompanyHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/v1")
	// Refresh time can be longer than token timeout
	auth.GET("/v1/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// LOGIN ROUTES
		auth.POST("/logout", authMiddleware.LogoutHandler)

		// USERS ROUTES
		auth.GET("companies", controllers.CompaniesHandler)
		auth.GET("company/:id", controllers.CompanyByIDHandler)
		auth.PUT("company/:id", controllers.UpdateCompanyHandler)
		auth.DELETE("company/:id", controllers.DeleteCompanyHandler)

		// COMMERCES ROUTES
		auth.POST("commerce", controllers.CreateCommerceHandler)
		auth.GET("commerces", controllers.CommercesHandler)
		auth.GET("commerce/:id", controllers.CommerceByIDHandler)
		auth.PUT("commerce/:id", controllers.UpdateCommerceHandler)
		auth.DELETE("commerce/:id", controllers.DeleteCommerceHandler)

		// CUSTOMERS ROUTES
		auth.POST("customer/:commerceid", controllers.CreateCustomerHandler)
		auth.GET("customers", controllers.CustomersHandler)
		auth.GET("customers/:commerceid", controllers.CustomerByCommerceHandler)
		auth.PUT("customer/:id", controllers.UpdateCustomerHandler)
		auth.DELETE("customer/:id", controllers.DeleteCustomerHandler)
	}

	return r
}
