module github.com/agp/push-notifications-api/entity

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
