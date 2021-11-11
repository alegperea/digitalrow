package entity

type Commerce struct {
	UUID      string     `json:"UUID" index:"UUID-index,hash" swaggerignore:"true"`
	UserData  string     `json:"UserData" dynamo:",range" swaggerignore:"true"`
	Name      string     `json:"Name"`
	Address   string     `json:"Address"`
	Email     string     `json:"Email" index:"Email-index,hash"`
	Phone     string     `json:"Phone"`
	Capacity  int        `json:"Capacity"`
	QRCode    string     `json:"QRCode"`
	Customers []Customer `json:"Customers"`
	CreatedAt string     `json:"CreatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
	UpdatedAt string     `json:"UpdatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key

}
