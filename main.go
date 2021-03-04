package mysqlpool

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"runtime"
	"time"
)

var (
	pool *sql.DB
)

func New(
	dsn string,
	maxopencon int,
	lifetime time.Duration,
	logger *log.Logger,
) error {
	cfg, errCfg := mysql.ParseDSN(dsn)
	if errCfg != nil {
		return errCfg
	}
	cn, errCn := mysql.NewConnector(cfg)
	if errCn != nil {
		return errCn
	}
	pool = sql.OpenDB(cn)
	pool.SetMaxIdleConns(runtime.NumCPU())
	pool.SetMaxOpenConns(maxopencon)
	//set to 0 for reuse connections not recommend
	//because Pool crashed in long time work (in server for example)
	pool.SetConnMaxLifetime(lifetime)
	// redirect mysql errLog to logger
	if err := mysql.SetLogger(logger); err != nil {
		return fmt.Errorf("[ERROR] mysql logger set error: %s", err.Error())
	}
	return nil
}
