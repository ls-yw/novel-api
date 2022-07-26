package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"novel/woodlsy"
	"novel/woodlsy/log"
	"time"
)

//type Mysql struct {
//	db *gorm.DB
//}

func connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		woodlsy.Configs.Databases.UserName,
		woodlsy.Configs.Databases.Password,
		woodlsy.Configs.Databases.Host,
		woodlsy.Configs.Databases.Port,
		woodlsy.Configs.Databases.Dbname,
		woodlsy.Configs.Databases.Charset,
	)

	dbLogger := log.NewDbLogger(ormLogger.Info, log.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,       // 禁用彩色打印
	})

	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   woodlsy.Configs.Databases.Prefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                             // use singular table name, table for `User` would be `user` with this option enabled
			//NoLowerCase:   true,                              // skip the snake_casing of names
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: dbLogger,
	})

	if err != nil {
		log.Logger.Error("数据库连接失败", err, dsn)
		panic("数据库链接失败")
	}
	fmt.Println("数据库 连接成功")
	return con
}
