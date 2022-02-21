package mysqlpool

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	Pool                   *sql.DB
	MaxAllowedPacketLength uint64
)

func getMaxAllowedPacketLength() (maxpack uint64) {
	const (
		MaxAllowedPacketMysqlDrv = 4194304
	)
	q := fmt.Sprint("SHOW VARIABLES LIKE 'max_allowed_packet'")
	mpname := ""
	maxpack = 1024
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, errCon := Pool.Conn(ctx)
	if errCon != nil {
		return
	}
	defer conn.Close()
	err := conn.QueryRowContext(ctx, q).Scan(&mpname, &maxpack)
	if err != nil {
		maxpack = 1024
	}
	if maxpack > MaxAllowedPacketMysqlDrv {
		maxpack = MaxAllowedPacketMysqlDrv
	}
	return
}

//New Pool
func New(dsn string, maxopencon int, lifetime time.Duration, logger *log.Logger) error {
	cfg, errCfg := mysql.ParseDSN(dsn)
	if errCfg != nil {
		return errCfg
	}
	cn, errCn := mysql.NewConnector(cfg)
	if errCn != nil {
		return errCn
	}
	Pool = sql.OpenDB(cn)
	Pool.SetMaxIdleConns(runtime.NumCPU())
	Pool.SetMaxOpenConns(maxopencon)
	//set to 0 for reuse connections not recommend
	//because Pool crashed in long time work (in server for example)
	Pool.SetConnMaxLifetime(lifetime)
	// redirect mysql errLog to logger
	if err := mysql.SetLogger(logger); err != nil {
		return fmt.Errorf("[ERROR] mysql logger set error: %s", err.Error())
	}
	MaxAllowedPacketLength = getMaxAllowedPacketLength()
	return nil
}
