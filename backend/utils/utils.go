package utils

import (
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"backend/config"
	"strconv"//Go 的标准库包，专门用于字符串与其他数据类型之间的转换
)

func GenerateJWT(username string, userlevel int)(string, error) {
	SecretKey, exp := config.GetJWT();
	expireTime, err := strconv.Atoi(exp);//把string类型的exp转换为int类型
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err;
	}
	log.Println("SecretKey:", SecretKey) // 检查 SecretKey
    log.Println("exp:", exp)             // 检查 exp

	expirationTime := time.Now().Add(time.Duration(expireTime) * time.Second);//time.Duration的作用是将int类型的时间转换为time.Duration类型（后面可以乘上time.Hour/Second等等）

	claim := jwt.MapClaims{
		"name":   username, 
        "level":  userlevel,
        "exp":    expirationTime.Unix(),
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim);//生成一个token

	tokenString, err := token.SignedString([]byte(SecretKey));//生成一个签名字符串
	if err != nil {
		//日志操作
		log.Printf("Error signing token: %v", err)
	}
	log.Println("Generated Token:", tokenString) // 检查生成的 Token

	return tokenString, nil;
}