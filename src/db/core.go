package core

import (
	"gorm.io/driver/mysql" //GORM 提供的 MySQL 驱动程序包。它实现了与 MySQL 数据库的交互，包括连接管理、查询执行、事务处理等。
	"gorm.io/gorm"         //GORM 提供的核心包。它提供了数据库模型定义、查询构建、数据操作等功能。通过 GORM 包，你可以更方便地进行数据库操作，如创建、查询、更新和删除数据等。
	"gorm.io/gorm/logger"  //提供了用于日志记录的功能。你可以使用该包中的日志记录器来记录数据库操作的日志，包括 SQL 语句、参数等信息。
	"gorm.io/gorm/schema"  //提供了数据库模型的命名策略设置。你可以使用这个包中的 NamingStrategy 结构体配置数据库表名和列名的命名规则，如是否使用复数形式等。
)

// 定义db全局变量
var Db *gorm.DB //声明一个全局变量 Db ，类型为 *gorm.DB ，表示这是一个 GORM 数据库连接对象。

func init() {
	var err error
	dsn := "root:asdasd@tcp(127.0.0.1:3306)/go1_preoject?charset=utf8mb4&parseTime=True&loc=Local"
	//使用 gorm.Open() 方法连接到 MySQL 数据库，并返回一个 gorm.DB 类型的对象。这里将连接对象赋值给全局变量 Db ，方便在其他地方使用。
	//gorm.Open() 的第二个参数是一个 gorm.Config 类型的实例，用于配置数据库的行为。在这个配置中，设置了一些选项，如是否跳过默认事务处理、定义命名策略、启用 SQL 日志等。
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, //设置为 false，表示不跳过默认的事务处理，即在 GORM 的数据库操作中默认启用事务。
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表明加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), //打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约
	})
	if err != nil {
		//如果数据库连接出现错误，将会触发一个 panic，并打印出带有连接失败提示信息和具体错误原因的错误消息。
		//这有助于在连接数据库失败时，立即中止程序的执行，以避免继续执行可能会导致更严重问题的代码。
		panic("Connecting database failed: " + err.Error())
	}
}

// GetDB
func GetDB() *gorm.DB {
	return Db
}
