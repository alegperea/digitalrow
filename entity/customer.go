package entity

type Customer struct {
	CustomerID       string `json:"CustomerID"`
	NotificationType string `json:"NotificationType"`
	Phone            string `json:"Phone"`
	Diners           int    `json:"Diners"`
	CreatedAt        string `json:"CreatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
}
