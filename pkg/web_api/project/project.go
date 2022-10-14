package project

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func ProjectList(c *gin.Context)  {
	//var projects []module.Project
	projects := []map[string]interface{}{}
	db.MySQL.Model(&module.Project{}).Select("id, name, code").Find(&projects)
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": projects})
}

type CreateProjectRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func CreateProject(c *gin.Context){
	defer func() {
		if err := recover(); err != nil{
			c.JSON(500, gin.H{"code": 1, "msg": "Create project failed", "data": []string{}})
			return
		}
	}()
	request := CreateProjectRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Request body error", "data": []string{}})
		return
	}
	//groups := []module.ProjectGroup{module.ProjectGroup{Code: fmt.Sprintf("%s.sa", request.Code)}}
	project := module.Project{Name: request.Name, Code: request.Code, }
	result := db.MySQL.Model(&module.Project{}).Create(&project)
	var mysqlErr *mysql.MySQLError
	if  errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		c.JSON(500, gin.H{"code": 1, "msg": "Project was exist", "data": []string{}})
		return
	}
	group := module.ProjectGroup{Name: fmt.Sprintf("%s项目管理员", project.Name), Code: fmt.Sprintf("%s.sa", project.Code),
		Project: project.ID}
	db.MySQL.Create(&group)
	groupPermission := []module.GroupPermission{module.GroupPermission{GroupID: group.ID, PermissionTag: "group.w"},
		module.GroupPermission{GroupID: group.ID, PermissionTag: "group.r"}}
	db.MySQL.Create(&groupPermission)
	c.JSON(200, gin.H{
		"code": 0,
		"msg": fmt.Sprintf("Create project success"),
		"data": gin.H{
			"id": project.ID,
		} ,
	})
	return
}

func ProjectInfo(c *gin.Context){
	project := c.Param("project")
	projectInfo := map[string]interface{}{}
	db.MySQL.Model(&module.Project{}).Where("id = ?", project).Select("id, name, code").Find(&projectInfo)
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": projectInfo})
}
