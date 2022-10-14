package user

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

var MySecret = []byte("cmdb")

func Login(c *gin.Context)  {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Request body error", "data": []string{}})
		return
	}
	// 算法错误，加密是一次性的，只能做校验
	//hasedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	//if err != nil {
	//	c.JSON(500, gin.H{"code": 1, "msg": "bcrypt failed", "data": []string{}})
	//	panic(err)
	//	return
	//}
	user := module.User{UserName: request.UserName}
	if db.MySQL.Where(&user).First(&user).RowsAffected == 0 {
		c.JSON(401, gin.H{"code": 1, "msg": "username not exists", "data": []string{},})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    1,
			"message": "password was mistake",
			"data": []string{},
		})
		return
	}

	// 生成token加盐算法
	claims := &Claims{
		UserId: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: -1, //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(MySecret); err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Signing token failed", "data": []string{}})
		return
	} else {
		ExpiresTime := 2 * time.Hour
		result := db.Redis.TokenCache.Set(tokenString, "0", ExpiresTime)
		fmt.Println(result.Err())
		c.JSON(200, gin.H{"code": 0, "msg": "", "data": gin.H{"token": tokenString}})
		return
	}

}


func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	// 若token只是过期claims是有数据的，若token无法解析claims无数据
	return claims, err
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.JSON(401, gin.H{"code": 1, "msg": "token can not nil", "data": []string{}})
			context.Abort()
			return
		}
		authCache := db.Redis.TokenCache.Get(auth)

		if authCache.Err() != nil {
			context.JSON(403, gin.H{"code": 1, "msg": "token was exp", "data": []string{}})
			context.Abort()
			return
		}else {
			claims, _ := ParseToken(auth)
			context.Set("userid", claims.UserId)
			d, _ := db.Redis.TokenCache.TTL(auth).Result()
			if  d < 10 * time.Minute{
				ExpiresTime := 2 * time.Hour
				db.Redis.TokenCache.Set(auth, ExpiresTime, ExpiresTime)
			}
			context.Next()
		}

	}
}

func AuthSuperAdmin() gin.HandlerFunc{
	return func(context *gin.Context) {
		id, _ := context.Get("userid")
		user := module.User{}
		db.MySQL.Where("id = ?", id).Find(&user)

		result := []map[string]interface{}{}
		if result := db.MySQL.Table("user2group").Select("user.username, project_group.code").
			Joins("left join user on user.id = user2group.user_id ").
			Joins("left join project_group on project_group.id = user2group.group_id").
			Where("project_group.code = 'cmdb.sa' and user2group.user_id = ?", id).Find(&result); result.RowsAffected == 0 {
			context.JSON(403, gin.H{"code": 1, "msg": "user not super admin", "data": []string{}})
			context.Abort()
		}
		context.Next()

	}
}

func AuthSuperAdminOrOthers() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.Get("userid")
		fmt.Println(id)
		project := context.Param("project")
		user := module.User{}
		db.MySQL.Where("id = ?", id).Find(&user)
		//用户是否为超级管理员
		result := []map[string]interface{}{}
		if result := db.MySQL.Table("user2group").Select("user.username, project_group.code").
			Joins("left join user on user.id = user2group.user_id ").
			Joins("left join project_group on project_group.id = user2group.group_id").
			Where("project_group.code = 'cmdb.sa' and user2group.user_id = ?", id).Find(&result); result.RowsAffected != 0 {
			context.Next()
			return
		}
		// 不是超级管理员，判断相关项目的组是否有权限
		permissionURLS := []map[string]interface{}{}
		userGroup := db.MySQL.Table("user2group").Select("group_id").Where("user_id = ?", id)
		userProjectGroup := db.MySQL.Table("project_group").Select("id").Where("project = ? and id in (?)", project, userGroup)
		permissionTag := db.MySQL.Table("group_permission").Select("permission_tag").Where("group_id in (?)", userProjectGroup)
		db.MySQL.Table("api_permission_group").Select("api,method").Where("tag in (?)", permissionTag).Find(&permissionURLS)
		fmt.Println(permissionURLS)
		for _, permission := range permissionURLS {
			if permission["api"] == context.FullPath() && permission["method"] == context.Request.Method{
				context.Next()
				return
			}
		}
		context.JSON(403, gin.H{"code": 1, "msg": "Permission dened", "data": []string{}})
		context.Abort()
	}
}