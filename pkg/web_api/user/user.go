package user

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AddUserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	FullName string `json:"fullname"`
}

func CreateUser(c *gin.Context)  {
	request := AddUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Request body error", "data": []string{}})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	result := db.MySQL.Create(&module.User{UserName: request.UserName, Password: string(hasedPassword), Name: request.Name, FullName: request.FullName})
	var mysqlErr *mysql.MySQLError
	if  errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		c.JSON(500, gin.H{"code": 1, "msg": "User was exist", "data": []string{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Create user success", "data": []string{}})
}
