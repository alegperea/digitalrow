module github.com/agp/push-notifications-api/controllers

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
	github.com/aws/aws-sdk-go v1.41.2 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/agp/push-notifications-api/entity v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/httputil v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/models v0.0.0-00010101000000-000000000000
	github.com/agp/push-notifications-api/services v0.0.0-00010101000000-000000000000
)
