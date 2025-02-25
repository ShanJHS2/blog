package middlewares

import (
	"backend/config"
	"backend/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Auth 中间件，用于验证 JWT 令牌和权限等级
func Auth(requiredLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey, _ = config.GetJWT() //获取JWT密钥
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ") //（注意Bearer后面有个空格）用于处理字符串的函数，从目标字符串处去掉特定前缀，返回处理完成后的字符串；若没有该前缀则返回原字符串。需要import "strings"
		if tokenString == authHeader {                           //在HTTP请求头中，Authorization字段的值一般是 Authorization： Bearer token（此处的token即是JWT令牌字符串）
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		//以上函数是验证HTTP请求的格式是否正确，下面是验证JWT令牌的有效性

		//jwt.Parse 是 JWT-go 库中用于解析和验证 JSON Web Token (JWT) 的核心函数。它的主要功能是解析一个 JWT 字符串，将其解码为一个 Token 对象，并验证其签名是否有效。两个参数分别为：JWT 字符串和一个回调函数（该回调函数返回SecretKey）。
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //这个回调函数的作用是根据Token的头部信息返回对应的密钥，以供后续验证签名;为了应对多个密钥和动态密钥的情况，回调函数的返回值是一个接口类型，而不是一个具体的密钥。
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //确保token的签名方法是HMAC，比如HS256/RS256等
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		}) //返回一个Token对象和一个错误对象，Token对象包含了解析后的JWT令牌信息（包括头Header\载荷payload（此处叫做claims）\验证信息valid），错误对象包含了解析过程中的错误信息。如果解析成功，错误对象为nil；否则，Token对象为nil。

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims) //payload是JWT的第二部分，claims指payload的内容
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
				return
			}
		}

		// 获取用户的 name
		nameValue, nameExists := claims["name"]
		if !nameExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token does not contain user name"})
			c.Abort()
			return
		}

		name, ok := nameValue.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user name format"})
			c.Abort()
			return
		}

		// 从数据库中查找用户
		var user models.User
		if err := config.DB.Where("Username = ?", name).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 判断用户权限是否符合要求
		if user.Level < requiredLevel {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// 将用户信息存入 Context 供后续使用
		c.Set("user", user)

		// 继续请求处理
		c.Next()

	}
}
