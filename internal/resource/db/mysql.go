package db

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Engine struct {
	db  *gorm.DB
	cfg *conf.Data_Mysql
}

func NewClient(cfg *conf.Data_Mysql) *Engine {
	engine := &Engine{}
	engine.cfg = cfg
	err := engine.connection()
	if err != nil {
		return nil
	}
	engine.Migrate()
	return engine
}

func (e *Engine) connection() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		e.cfg.Username,
		e.cfg.Password,
		e.cfg.Host,
		e.cfg.Port,
		e.cfg.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.NewGormLogger(zap.InfoLevel, zap.InfoLevel, time.Millisecond*500),
	})
	if err != nil {
		panic("连接数据库失败 " + err.Error())
	}
	e.db = db
	return nil
}
func (e *Engine) Migrate() {
	// if err := e.db.AutoMigrate(model.Article{}); err != nil {
	//
	// }
}

func NewMySQLClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.NewGormLogger(zap.InfoLevel, zap.InfoLevel, time.Millisecond*500),
	})
	if err != nil {
		panic("连接数据库失败 " + err.Error())
	}
	return db
}
