package initialize

import (
	"fmt"
	"ginDemo/global"
	"ginDemo/model/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL() {
	// 配置
	var addr, port, user, password, dbname string
	addr = global.VP.GetString("db.addr")
	port = global.VP.GetString("db.port")
	user = global.VP.GetString("db.user")
	password = global.VP.GetString("db.password")
	dbname = global.VP.GetString("db.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, port, dbname)

	// 连接
	var err error
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("数据库出问题啦: %s\n", err))
	}

	// 迁移
	global.DB.AutoMigrate(
		&database.User{},
	)
}

func Close() {
	err := global.DB.Close()
	if err != nil {
		panic(fmt.Errorf("数据库出问题啦: %s\n", err))
	}
}
