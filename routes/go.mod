module github.com/agp/push-notifications-api/routes

go 1.16

replace (
	github.com/agp/push-notifications-api/controllers => ../controllers
	github.com/agp/push-notifications-api/docs => ../docs
	github.com/agp/push-notifications-api/entity => ../entity
	github.com/agp/push-notifications-api/httputil => ../httputil
	github.com/agp/push-notifications-api/middlewares => ../middlewares
	github.com/agp/push-notifications-api/models => ../models
	github.com/agp/push-notifications-api/routes => ../routes
	github.com/agp/push-notifications-api/services => ../services

)

require (
	github.com/appleboy/gin-jwt/v2 v2.7.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/google/uuid v1.3.0 // indirect
	github.com/agp/push-notifications-api/controllers v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/docs v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/middlewares v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/services v0.0.0-00010101000000-000000000000
	github.com/swaggo/gin-swagger v1.3.2
)
