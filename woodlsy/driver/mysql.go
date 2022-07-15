package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"novel/utils/log"
	"novel/woodlsy"
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
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   woodlsy.Configs.Databases.Prefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                             // use singular table name, table for `User` would be `user` with this option enabled
			//NoLowerCase:   true,                              // skip the snake_casing of names
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})

	if err != nil {
		log.Logger.Error("数据库连接失败", err, dsn)
		panic("数据库链接失败")
	}
	fmt.Println("链接数据库成功")
	return con
}
