module github.com/agp/push-notifications-api

go 1.16

replace (
	github.com/agp/push-notifications-api/controllers => ./controllers
	github.com/agp/push-notifications-api/docs => ./docs
	github.com/agp/push-notifications-api/entity => ./entity
	github.com/agp/push-notifications-api/httputil => ./httputil
	github.com/agp/push-notifications-api/middlewares => ./middlewares
	github.com/agp/push-notifications-api/models => ./models
	github.com/agp/push-notifications-api/routes => ./routes

	github.com/agp/push-notifications-api/services => ./services

)

require (
	github.com/apex/gateway v1.1.2
	github.com/aws/aws-lambda-go v1.17.0
	github.com/joho/godotenv v1.4.0
	github.com/agp/push-notifications-api/routes v0.0.0-00010101000000-000000000000
	github.com/xhit/go-simple-mail/v2 v2.10.0 // indirect
)
