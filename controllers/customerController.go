package controllers

import (
	"net/http"

	"github.com/agp/push-notifications-api/entity"
	"github.com/agp/push-notifications-api/httputil"
	model "github.com/agp/push-notifications-api/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/gin-gonic/gin"
)

// CompanyHandler godoc
// @Summary List Customers
// @Description add by json dataset
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param body body entity.Customer true "List customers"
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/customers [get]
func CustomersHandler(c *gin.Context) {

	var customers []entity.Customer
	errGet := model.GetAllCustomers(&customers)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}
	c.JSON(http.StatusOK, customers)

}

// CreateCompanyHandler godoc
// @Summary Create Customer
// @Description add by json customer
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param body body entity.Customer true "Add customer"
// @Success 200 {object} entity.Customer
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/customer/:commerceid [post]
func CreateCustomerHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)

	var customer entity.Customer

	if errJson := c.ShouldBindJSON(&customer); errJson != nil {
		httputil.NewError(c, http.StatusBadRequest, errJson)
		return
	}

	sess := c.MustGet("sess").(*session.Session)
	commerceID := c.Params.ByName("commerceid")

	commerce, errGet := model.GetCommerceByID(sess, uuID, commerceID)
	if errGet != nil {
		httputil.NewError(c, http.StatusInternalServerError, errGet)
		return
	}
	errCre := model.AddCustomer(sess, uuID, &commerce, &customer)
	if errCre != nil {
		httputil.NewError(c, http.StatusInternalServerError, errCre)
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}

// CompanyByIDHandler godoc
// @Summary Customer by ID
// @Description get string by ID
// @Tags Customer
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Customer
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/customer/:commerceid [get]
func CustomerByCommerceHandler(c *gin.Context) {
	//id := c.Params.ByName("id")
	commerceID := c.Params.ByName("commerceid")

	sess := c.MustGet("sess").(*session.Session)
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)
	customers, errGet := model.GetCustomerByCommerce(sess, uuID, commerceID)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}
	c.JSON(http.StatusOK, customers)
	return
}

// UpdateCompanyHandler godoc
// @Summary Update Customer
// @Description get string by ID
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param body body entity.Customer true "Update customer"
// @Param  id path int true "customer ID"
// @Success 200 {object} entity.Customer
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/customer/:id [put]
func UpdateCustomerHandler(c *gin.Context) {
	var customer entity.Customer
	customerID := c.Params.ByName("id")
	customer.CustomerID = customerID

	sess := c.MustGet("sess").(*session.Session)
	errGet := model.GetCustomerByID(sess, customerID, &customer)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return

	}
	c.BindJSON(&customer)
	errUpd := model.UpdateCustomer(&customer, customerID)
	if errUpd != nil {
		httputil.NewError(c, http.StatusInternalServerError, errUpd)
		return
	}
	c.JSON(http.StatusOK, customer)
	return
}

// DeleteCompanyHandler godoc
// @Summary Delete Customer
// @Description get string by ID
// @Tags Customer
// @Accept  json
// @Produce  json
// @Success 204 {object} entity.Customer
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/customer/:id [delete]
func DeleteCustomerHandler(c *gin.Context) {
	var customer entity.Customer
	id := c.Params.ByName("id")
	errDel := model.DeleteCustomer(&customer, id)
	if errDel != nil {
		httputil.NewError(c, http.StatusInternalServerError, errDel)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
	return
}
