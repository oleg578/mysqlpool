package mysqlpool

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	Pool *sql.DB
)

// New Pool
func New(dsn string, maxopencon int, lifetime time.Duration, maxAllowedPacket int, logger *log.Logger) error {
	cfg, errCfg := mysql.ParseDSN(dsn)
	cfg.MaxAllowedPacket = maxAllowedPacket
	if errCfg != nil {
		return errCfg
	}
	cn, errCn := mysql.NewConnector(cfg)
	if errCn != nil {
		return errCn
	}
	Pool = sql.OpenDB(cn)
	Pool.SetMaxOpenConns(maxopencon)
	Pool.SetMaxIdleConns(runtime.NumCPU())
	//set to 0 for reuse connections not recommend
	//because Pool crashed in long time work (in server, for example)
	Pool.SetConnMaxLifetime(lifetime)
	// redirect mysql errLog to logger
	if err := mysql.SetLogger(logger); err != nil {
		return fmt.Errorf("[ERROR] mysql logger set error: %s", err.Error())
	}
	//check Pool created successfully
	if err := Pool.Ping(); err != nil {
		return fmt.Errorf("[ERROR] new mysql pool check error: %s", err.Error())
	}
	return nil
}
