package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"backend/models"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Username     string
		Password string
		Name     string
	}
	Jwt struct {
		SecretKey string
		Expiration string
	}
}

var AppConfig *Config//包级变量，初始时直接实例化（根据导出规则，首字母大写的包级变量可以直接在其他包中加上包名使用）
var DB *gorm.DB//数据库连接的全局变量

func InitConfig() {//只有首字母大写的函数才能导出
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)//日志相关操作暂不熟悉，暂且这么写
	}

	AppConfig = &Config{}//appconfig不会是nil，可以填充数据
	if err:= viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct,%v",err)//日志相关操作暂不熟悉，暂且这么写
	}

	if AppConfig == nil {
        log.Fatalf("fuck!AppConfig is not initialized")
    }
} 

func InitDB() error {// 初始化全局连接池DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Database.Username,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name,
	)

	var err error;
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// 自动迁移数据库模型
    if err := DB.AutoMigrate(
        &models.BlogArticle{},
        &models.ResearchArticle{},
        &models.ProjectArticle{},
        &models.Comment{},
        &models.Media{},
        &models.Question{},
        &models.User{},
    ); err != nil {
        return fmt.Errorf("error migrating database: %v", err)
    }

    return nil	
}

func GetJWT() (string, string) {
	if AppConfig == nil {
        log.Fatalf("shit!AppConfig is not initialized")
    }
	return AppConfig.Jwt.SecretKey, AppConfig.Jwt.Expiration;
}