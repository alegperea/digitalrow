module github.com/agp/push-notifications-api/models

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
	github.com/aws/aws-sdk-go v1.41.3
	github.com/guregu/dynamo v1.11.0
	github.com/agp/push-notifications-api/entity v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)
