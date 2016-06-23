// user
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
	"github.com/vsdutka/ps/utils"
)

func UserController(c *gin.Context) {
	action := strings.ToLower(c.Params.ByName("action"))
	fmt.Println("action: ", action)
	var form models.User
	if action != "/" {
		if !c.Bind(&form) {
			utils.JsonError(c.Writer, errgo.New("Отсутствуют данные"))
			return
		}
	}
	fmt.Println(form)
	switch action {
	case "/":
		ul, err := models.UserList()
		if err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, ul)
	case "/register":
		if err := form.UserCreate(); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		if err := utils.SmsSend(form.Phone, fmt.Sprintf("Pin code : %s", form.PinCode)); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}

		utils.JsonData(c.Writer, struct {
			Status string `json:"status"`
			ID     uint   `json:"id"`
			Pin    string `json:"pin"`
		}{
			Status: "OK",
			ID:     form.ID,
			Pin:    form.PinCode,
		})
	case "/confirm":
		if err := form.UserUpdate(); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, struct {
			Status    string `json:"status"`
			SecretKey string `json:"secret_key"`
		}{
			Status:    "OK",
			SecretKey: form.SecretKey,
		})
	case "/get":
		if err := form.UserGet(); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonData(c.Writer, form)
	case "/upd":
		if err := form.UserUpdate(); err != nil {
			utils.JsonError(c.Writer, err)
			return
		}
		utils.JsonOK(c.Writer)
	case "/del":
		if err := form.UserDel(); err != nil {
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
