// person
package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/errgo.v1"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"

	"github.com/vsdutka/ps/models"
	//	"github.com/vsdutka/ps/shared/passhash"
	"github.com/vsdutka/ps/shared/database"
	"github.com/vsdutka/ps/utils"
)

func PersonController(c *gin.Context) {
	action := strings.ToLower(c.Params.ByName("action"))
	fmt.Println("action: ", action)
	var form models.Person
	if action != "/" {
		if !c.Bind(&form) {
			utils.JsonError(c.Writer, errgo.New("Отсутствуют данные"))
			return
		}
	}
	switch action {
	case "/":
		var ol []models.Person
		if err := database.List(&ol); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, ol)
	case "/get":
		if err := database.Sel(&form, form.ID); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, form)
	case "/add":
		if err := database.Ins(&form); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, struct {
			Status string `json:"status"`
			ID     uint   `json:"id"`
		}{
			Status: "OK",
			ID:     form.ID,
		})
	case "/upd":
		fmt.Println(form)
		if err := database.Upd(&form); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		fmt.Println(form)
		utils.JsonOK(c.Writer)
	case "/del":
		if err := database.Del(&form); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonOK(c.Writer)
	default:
		fmt.Println("default")
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("404 Page Not Found"))
	}
}
