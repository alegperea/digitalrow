package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/agp/push-notifications-api/entity"
	"github.com/agp/push-notifications-api/httputil"
	"github.com/agp/push-notifications-api/services"
	"github.com/aws/aws-sdk-go/aws/session"

	model "github.com/agp/push-notifications-api/models"

	"github.com/gin-gonic/gin"
)

// CompanyHandler godoc
// @Summary List Companies
// @Description add by json dataset
// @Tags Company
// @Accept  json
// @Produce  json
// @Param body body entity.Company true "List companies"
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/companies [get]
func CompaniesHandler(c *gin.Context) {

	var company []entity.Company
	errGet := model.GetAllCompanies(&company)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}
	c.JSON(http.StatusOK, company)

}

// CreateCompanyHandler godoc
// @Summary Create Company
// @Description add by json company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param body body entity.Company true "Add company"
// @Success 200 {object} entity.Company
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/company [post]
func CreateCompanyHandler(c *gin.Context) {
	var company entity.Company
	if errJson := c.ShouldBindJSON(&company); errJson != nil {
		httputil.NewError(c, http.StatusBadRequest, errJson)
		return
	}

	sess := c.MustGet("sess").(*session.Session)
	errCre := model.CreateCompany(sess, &company)
	if errCre != nil {
		httputil.NewError(c, http.StatusInternalServerError, errCre)
		return
	}

	itemString, errJson := json.Marshal(company)
	if errJson != nil {
		httputil.NewError(c, http.StatusNotFound, errJson)
		return
	}
	_, errQueue := services.SetQueue(string(itemString), os.Getenv("sqs_email_name"))
	if errQueue != nil {
		httputil.NewError(c, http.StatusNotFound, errJson)
		return
	}

	c.JSON(http.StatusOK, company)
	return
}

// CompanyByIDHandler godoc
// @Summary Company by ID
// @Description get string by ID
// @Tags Company
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Company
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/company/:id [get]
func CompanyByIDHandler(c *gin.Context) {
	//id := c.Params.ByName("id")
	var company entity.Company

	sess := c.MustGet("sess").(*session.Session)

	errGet := model.GetCompanyByID(sess, &company)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}
	c.JSON(http.StatusOK, company)
	return
}

// UpdateCompanyHandler godoc
// @Summary Update Company
// @Description get string by ID
// @Tags Company
// @Accept  json
// @Produce  json
// @Param body body entity.Company true "Update company"
// @Param  id path int true "company ID"
// @Success 200 {object} entity.Company
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/company/:id [put]
func UpdateCompanyHandler(c *gin.Context) {
	var company entity.Company
	id := c.Params.ByName("id")
	company.UUID = id

	sess := c.MustGet("sess").(*session.Session)
	errGet := model.GetCompanyByID(sess, &company)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return

	}
	c.BindJSON(&company)
	errUpd := model.UpdateCompany(&company, id)
	if errUpd != nil {
		httputil.NewError(c, http.StatusInternalServerError, errUpd)
		return
	}
	c.JSON(http.StatusOK, company)
	return
}

// DeleteCompanyHandler godoc
// @Summary Delete Company
// @Description get string by ID
// @Tags Company
// @Accept  json
// @Produce  json
// @Success 204 {object} entity.Company
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/company/:id [delete]
func DeleteCompanyHandler(c *gin.Context) {
	var company entity.Company
	id := c.Params.ByName("id")
	errDel := model.DeleteCompany(&company, id)
	if errDel != nil {
		httputil.NewError(c, http.StatusInternalServerError, errDel)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
	return
}
