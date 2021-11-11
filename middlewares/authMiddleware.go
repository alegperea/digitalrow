package middleware

import (
	"bufio"
	"fmt"
	"os"
	"time"

	entity "github.com/agp/push-notifications-api/entity"
	model "github.com/agp/push-notifications-api/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/gin-gonic/gin"
)

var identityID = "id"
var identityCUIT = "cuit"
var identityEmail = "email"
var identityPassword = "password"
var identityCompanyName = "companyName"
var identityFirstName = "firstName"
var identityLastName = "lastName"
var identityPhone = "phone"
var identitySubscription = "subscription"

func AuthMiddleware() (jwt.GinJWTMiddleware, error) {

	privateKeyFile, err := os.Open("./jwt/private.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(pembytes),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityEmail,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.Company); ok {
				return jwt.MapClaims{
					identityID:           v.UUID,
					identityCUIT:         v.CUIT,
					identityEmail:        v.Email,
					identityPassword:     v.Password,
					identityCompanyName:  v.CompanyName,
					identityFirstName:    v.FirstName,
					identityLastName:     v.LastName,
					identityPhone:        v.Phone,
					identitySubscription: v.Subscription,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println("claims IdentityHandler")
			return &entity.Company{
				CUIT:         claims[identityCUIT].(string),
				Email:        claims[identityEmail].(string),
				Password:     claims[identityPassword].(string),
				CompanyName:  claims[identityCompanyName].(string),
				FirstName:    claims[identityFirstName].(string),
				LastName:     claims[identityLastName].(string),
				Phone:        claims[identityPhone].(string),
				Subscription: claims[identitySubscription].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var loginVals entity.Login

			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			//email := loginVals.Email
			//password := loginVals.Password
			sess := c.MustGet("sess").(*session.Session)

			company, validate := model.ValidateCompany(sess, loginVals)

			fmt.Println("User  ", company)

			if validate == true {
				return &entity.Company{
					UUID:         company.UUID,
					CUIT:         company.CUIT,
					Email:        company.Email,
					Password:     company.Password,
					CompanyName:  company.CompanyName,
					FirstName:    company.FirstName,
					LastName:     company.LastName,
					Phone:        company.Phone,
					Subscription: company.Subscription,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Println("Authorizator")

			if _, ok := data.(*entity.Company); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return *authMiddleware, err
}
