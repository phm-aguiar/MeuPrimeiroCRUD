package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phm-aguiar/MeuPrimeiroCRUD/src/configuration/validation"
	"github.com/phm-aguiar/MeuPrimeiroCRUD/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser Controller")
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to bind JSON: %s\n", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(UserRequest)

}
