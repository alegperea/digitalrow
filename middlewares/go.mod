module github.com/agp/push-notifications-api/middlewares

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
	github.com/gin-gonic/gin v1.7.4
	github.com/agp/push-notifications-api/entity v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/models v0.0.0-00010101000000-000000000000
)
