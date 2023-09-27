package controller

import (
	"cli-server/internal/cli-server/store"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func CreateTemplate(c *gin.Context) {
	var template store.Template
	if err := c.BindJSON(&template); err != nil {
		c.String(http.StatusBadRequest, "The template params is error")
		return
	}

	id, err := store.CreateTemplate(c, &template)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Creating template failed")
		return
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "OK",
		Data:    id,
	})
}

func TemplateList(c *gin.Context) {
	list, count, err := store.TemplateList(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Getting template list failed")
		return
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "OK",
		Data: map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

func UpdateTemplate(c *gin.Context) {
	var template store.Template
	if err := c.BindJSON(&template); err != nil {
		c.String(http.StatusBadRequest, "The template params is error")
		return
	}
	update, err := store.UpdateTemplate(c, &template)
	if err != nil {
		c.String(http.StatusInternalServerError, "Updating the template failed")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "OK",
		Data:    update,
	})
}

func DeleteTemplate(c *gin.Context) {
	value := c.Param("value")
	count, err := store.DeleteTemplate(c, value)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Deleting the template failed")
		return
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "OK",
		Data:    count,
	})
}
