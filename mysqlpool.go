package mysqlpool

import (
	"database/sql"
	"errors"
	"log"
	"runtime"
	"time"

	"github.com/go-sql-driver/mysql"
)

const MAX_ALLOWED_PACKET = 1073741824 // 1GB

var (
	Pool *sql.DB
)

// New Pool
func New(dsn string, maxOpenCon int, lifetime time.Duration, maxAllowedPacket int, logger *log.Logger) error {
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
	Pool.SetMaxOpenConns(maxOpenCon)
	Pool.SetMaxIdleConns(runtime.NumCPU())
	//set to 0 for reuse connections not recommend
	//because Pool crashed in long time work (in server, for example)
	Pool.SetConnMaxLifetime(lifetime)
	// redirect mysql errLog to logger
	if err := mysql.SetLogger(logger); err != nil {
		return errors.New("[ERROR] mysql logger set error: " + err.Error())
	}
	//check Pool created successfully
	if err := Pool.Ping(); err != nil {
		return errors.New("[ERROR] new mysql pool check error: " + err.Error())
	}
	return nil
}
