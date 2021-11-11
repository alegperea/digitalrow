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

// CommercesHandler godoc
// @Summary List Commerce
// @Description add by json commerce
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Param body body entity.Commerce true "List commerces"
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/commerces [get]
func CommercesHandler(c *gin.Context) {

	var commerces []entity.Commerce

	sess := c.MustGet("sess").(*session.Session)
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)

	commerces, errGet := model.GetAllCommerces(sess, uuID)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}

	c.JSON(http.StatusOK, commerces)

}

// CreateCommerceHandler godoc
// @Summary Create Commerce
// @Description add by json Commerce
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Param body body entity.Commerce true "Add Commerce"
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/commerce [post]
func CreateCommerceHandler(c *gin.Context) {

	var commerce entity.Commerce
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)
	companyName := claims["companyName"].(string)

	if errJson := c.ShouldBindJSON(&commerce); errJson != nil {
		httputil.NewError(c, http.StatusBadRequest, errJson)
		return
	}
	sess := c.MustGet("sess").(*session.Session)
	errCre := model.CreateCommerce(sess, companyName, &commerce, uuID)
	if errCre != nil {
		httputil.NewError(c, http.StatusInternalServerError, errCre)
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}

// CommerceByIDHandler godoc
// @Summary Commerce by ID
// @Description get string by ID
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/commerce/:id [get]
func CommerceByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)

	sess := c.MustGet("sess").(*session.Session)

	commerce, errGet := model.GetCommerceByID(sess, uuID, id)
	if errGet != nil {
		httputil.NewError(c, http.StatusNotFound, errGet)
		return
	}
	c.JSON(http.StatusOK, commerce)
	return
}

// UpdateCommerceHandler godoc
// @Summary Update Commerce
// @Description get string by ID
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Param body body entity.Commerce true "Update Commerce"
// @Param  id path int true "Commerce ID"
// @Success 200 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/commerce/:id [put]
func UpdateCommerceHandler(c *gin.Context) {
	var commerce entity.Commerce
	id := c.Params.ByName("id")
	claims := jwt.ExtractClaims(c)
	uuID := claims["id"].(string)

	sess := c.MustGet("sess").(*session.Session)

	c.BindJSON(&commerce)
	errUpd := model.UpdateCommerce(sess, &commerce, uuID, id)
	if errUpd != nil {
		httputil.NewError(c, http.StatusInternalServerError, errUpd)
		return
	}
	c.JSON(http.StatusOK, commerce)
	return
}

// DeleteCommerceHandler godoc
// @Summary Delete Commerce
// @Description get string by ID
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Success 204 {object} entity.Commerce
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Header 200 {string} Authorization "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
// @Header 200 {string} Content-Type "application/json"
// @Router /v1/commerce/:id [delete]
func DeleteCommerceHandler(c *gin.Context) {
	var commerce entity.Commerce
	id := c.Params.ByName("id")
	claims := jwt.ExtractClaims(c)

	uuID := claims["id"].(string)

	sess := c.MustGet("sess").(*session.Session)
	errDel := model.DeleteCommerce(sess, &commerce, uuID, id)
	if errDel != nil {
		httputil.NewError(c, http.StatusInternalServerError, errDel)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
	return
}
