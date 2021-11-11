package entity

type Company struct {
	UUID         string `json:"UUID" index:"UUID-index,hash" swaggerignore:"true"`
	UserData     string `json:"UserData" dynamo:",range" swaggerignore:"true"`
	CUIT         string `json:"CUIT"`
	Email        string `json:"Email" index:"Email-index,hash"`
	Password     string `json:"Password"`
	CompanyName  string `json:"CompanyName"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Phone        string `json:"Phone"`
	Logo         string `json:"Logo"`
	Subscription string `json:"Subscription"`
	CreatedAt    string `json:"CreatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
	UpdatedAt    string `json:"UpdatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
}

//Login is a simple type to contain username and password
type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
