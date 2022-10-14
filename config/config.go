package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type Gin struct {
	AppMode     string
	Port        string
	BackendPort string
	JwtKey      string
	ExpireTime  int
	MDPath      string
	MDUrlPath   string
}

type Mysql struct {
	Db       string
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

type GORM struct {
	LogMode                                  int
	TablePrefix                              string
	SingularTable                            bool
	SkipDefaultTransaction                   bool
	DisableForeignKeyConstraintWhenMigrating bool
	MaxIdleConns                             int
	MaxOpenConns                             int
	SetConnMaxLifetime                       time.Duration
}

type TencentCloud struct {
	Url         string
	SecretID    string
	SecretKey   string
	ArticlePath string
	UserPath    string
	TalkPath    string
	AlbumPath   string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
	MaxIdle  int
	Network  string
}

type Session struct {
	Name   string
	Salt   string
	MaxAge int
}

type EMail struct {
	Sender     string
	AuthCode   string
	Title      string
	BodyType   string
	SMTPAddr   string
	SMTPPort   int
	ExpireTime int
}

var (
	configPath  = "./config/config.toml"
	Conf        = &Config{}
	GinConf     = &Gin{}
	MysqlConf   = &Mysql{}
	GORMConf    = &GORM{}
	TcConf      = &TencentCloud{}
	RedisConf   = &Redis{}
	SessionConf = &Session{}
	EMailConf   = &EMail{}
)

type Config struct {
	Gin          *Gin
	Mysql        *Mysql
	GORM         *GORM
	TencentCloud *TencentCloud
	Redis        *Redis
	Session      *Session
	EMail        *EMail
}

func init() {
	if _, err := toml.DecodeFile(configPath, &Conf); err != nil {
		log.Fatal(err)
	}
	GinConf = Conf.Gin
	MysqlConf = Conf.Mysql
	GORMConf = Conf.GORM
	TcConf = Conf.TencentCloud
	RedisConf = Conf.Redis
	SessionConf = Conf.Session
	EMailConf = Conf.EMail
}
