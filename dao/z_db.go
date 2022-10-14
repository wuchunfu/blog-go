package dao

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"myblog/config"
	"time"
)

var (
	db        *gorm.DB
	dbError   error
	dbConf    = config.MysqlConf
	gormConf  = config.GORMConf
	e         *casbin.Enforcer
	rdb       *redis.Client
	redisConf = config.RedisConf
)

// func InitDb() {
func init() {
	if db, dbError = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname,
		),
		DefaultStringSize: 256,
	}), &gorm.Config{
		SkipDefaultTransaction:                   gormConf.SkipDefaultTransaction,
		DisableForeignKeyConstraintWhenMigrating: gormConf.DisableForeignKeyConstraintWhenMigrating,
		Logger:                                   logger.Default.LogMode(logger.LogLevel(gormConf.LogMode)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: gormConf.SingularTable,
		},
	}); dbError != nil {
		log.Fatal("数据库连接失败，请检查参数 ", dbError)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("数据池错误，请检查", err)
	}
	sqlDB.SetMaxIdleConns(gormConf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(gormConf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(gormConf.SetConnMaxLifetime * time.Hour)
}

// func InitCasbinRule() {
func init() {
	var adapter, _ = gormadapter.NewAdapterByDB(db)
	e, _ = casbin.NewEnforcer("config/casbin.conf", adapter)
	/*e.AddPolicy("admin", "/admin/user/menus", "GET")
	e.AddPolicy("admin", "/admin/menus", "GET")
	e.AddPolicy("admin", "/admin/menus", "POST")
	e.AddPolicy("admin", "/admin/menus/:menuId", "DELETE")
	e.AddPolicy("admin", "/admin/resources", "GET")
	e.AddPolicy("admin", "/admin/resources", "POST")
	e.AddPolicy("admin", "/admin/resources/:resourcesId", "DELETE")
	e.AddPolicy("admin", "/admin/role/resources", "GET")
	e.AddPolicy("admin", "/admin/role/menus", "GET")
	e.AddPolicy("admin", "/admin/role", "POST")
	e.AddPolicy("admin", "/admin/roles", "GET")
	e.AddPolicy("admin", "/admin/roles", "DELETE")*/
	e.EnableAutoSave(true) // 开启自动持久化到数据库，不开启的话进行增删查改只是在内存上的改动，不会持久化到数据库
}

// func InitRedis() {
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
}

func GetRDB() *redis.Client {
	return rdb
}

func AddPolicy(params ...interface{}) {
	e.AddPolicy(params...)
}

func AddPolicies(rules [][]string) {
	e.AddPolicies(rules)
}

func RemovePolicy(params ...interface{}) {
	e.RemovePolicy(params...)
	e.LoadPolicy() // **重新加载一下策略才能让删除生效**
}

func ReturnEnforcer() *casbin.Enforcer {
	return e
}

// 下面这些只针对数据库中存在的表进行CRUD而不能对DTO、VO等进行使用，后面可以继续做出修改，使他们更完善

func Create[T any](data *T) {
	db.Create(&data)
}

// GetOne [通用]使用条件查找一条数据
func GetOne[T any](data T, query string, args ...any) T {
	db.Where(query, args...).First(&data)
	return data
}

// Updates [通用]单行更新: 传入对应结构体[传递主键用]和带有对应更新字段值的结构体,0和空字段不会更新
func Updates[T any](data *T, slt ...string) {
	if len(slt) > 0 {
		db.Model(&data).Select(slt).Updates(&data)
		return
	}
	db.Model(&data).Updates(&data)
}

func List[T any](data T, slt string, order string, query string, args ...any) T {
	tx := db.Select(slt).Order(order)
	if query != "" {
		tx = tx.Where(query, args...)
	}
	tx.Find(&data)
	return data
}

func Delete[T any](data T, query string, args ...any) {
	db.Where(query, args...).Delete(&data)
}

func Count[T any](data T, query string, args ...any) int {
	var count int64
	tx := db.Model(data)
	if query != "" {
		tx = tx.Where(query, args...)
	}
	tx.Count(&count)
	return int(count)
}
