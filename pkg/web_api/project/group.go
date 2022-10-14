package project

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// 查单组成员

func ProjectGroup(c *gin.Context) {
	id := c.Param("id")
	project := c.Param("project")
	if id == ""{
		c.JSON(400, gin.H{
			"code": 1,
			"msg": "Project ID can not null",
			"data": "",
		})
		return
	}
	groups := []map[string]interface{}{}

	// 限定只允许读本项目用户组

	if err := db.MySQL.Model(&module.ProjectGroup{}).Select("id").Where("project = ? and id = ?", project, id).First(&groups).Error; err == gorm.ErrRecordNotFound{
		c.JSON(403, gin.H{
			"code": 1,
			"msg": "Dont have this group permission",
			"data": []string{},
		})
		return
	}
	users := []map[string]interface{}{}
	userID := db.MySQL.Model(&module.User2Group{}).Select("user_id").Where("group_id = ?", id)
	db.MySQL.Model(&module.User{}).Select("id, username, name, fullname").Where("id in (?)", userID).Find(&users)

	fmt.Println(groups)
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"data": users,
	})
}

// 查项目组列表

func GroupList(c *gin.Context){
	project := c.Param("project")
	groups := []map[string]interface{}{}
	db.MySQL.Model(&module.ProjectGroup{}).Where("project = ?", project).Select("id,name,code").Find(&groups)
	c.JSON(200,gin.H{
		"code": 0,
		"msg": "",
		"data": groups,
	})
}

type CreateGroupRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func CreateGroup(c *gin.Context){
	request := CreateProjectRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Request body error", "data": []string{}})
		return
	}
	projectID, _ := strconv.Atoi(c.Param("project"))
	project := module.Project{ID: projectID}
	db.MySQL.First(&project)
	groupName := strings.Split(request.Name, ".")
	if groupName[0] != project.Name{
		c.JSON(400, gin.H{"code": 1, "msg": "Group name error", "data": []string{}})
		return
	}
	db.MySQL.Create(&module.ProjectGroup{Name: request.Name, Code: request.Code})
	c.JSON(200, gin.H{"code": 0, "msg": "Add group success", "data": []string{}})
}

func GroupPermission(c *gin.Context){
	project := c.Param("project")
	group := c.Param("id")
	groups := []map[string]interface{}{}

	// 限定只允许读本项目用户组

	if err := db.MySQL.Model(&module.ProjectGroup{}).Select("id").Where("project = ? and id = ?", project, group).First(&groups).Error; err == gorm.ErrRecordNotFound{
		c.JSON(403, gin.H{
			"code": 1,
			"msg": "Dont have this group permission",
			"data": []string{},
		})
		return
	}
	permission := []module.GroupPermission{}
	db.MySQL.Table("project_group").Select("group_permission.permission_tag").
		Joins("RIGHT JOIN group_permission ON group_permission.group_id = project_group.id").
		Where("project = ? and project_group.id = ?", project, group).Find(&permission)
	permissionList := []string{}
	for _, p := range permission {
		permissionList = append(permissionList, p.PermissionTag)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"data": permissionList,
	})
}