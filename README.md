# mysqlpool

`mysqlpool` is a Go package that provides a simple and efficient way to manage a pool of MySQL database connections. It leverages the `database/sql` package and the `go-sql-driver/mysql` driver to create and manage connections, ensuring optimal performance and resource utilization.

## Features

- **Connection Pooling**: Efficiently manage a pool of MySQL connections to handle multiple database requests.
- **Configurable Settings**: Customize the maximum number of open connections, connection lifetime, and maximum allowed packet size.
- **Error Logging**: Redirect MySQL error logs to a custom logger for better monitoring and debugging.
- **Automatic Connection Management**: Automatically handle connection creation, reuse, and cleanup.

## Installation

To install the `mysqlpool` package, use the following command:

```sh
go get github.com/yourusername/mysqlpool
```

## Usage

```go

package main

import (
	"github.com/oleg578/mysqlpool"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "mysqlpool: ", log.LstdFlags)
	err := mysqlpool.New("user:password@tcp(127.0.0.1:3306)/dbname", 10, time.Minute, mysqlpool.MAX_ALLOWED_PACKET, logger)
	if err != nil {
		log.Fatalf("Failed to create MySQL pool: %v", err)
	}

	// Use mysqlpool.Pool to interact with the database
}

```

## Configuration
DSN: Data Source Name for MySQL connection.
maxOpenCon: Maximum number of open connections to the database.
lifetime: Maximum amount of time a connection may be reused.
maxAllowedPacket: Maximum allowed packet size for MySQL.
logger: Custom logger for MySQL error logs.

## License
This project is licensed under the MIT License.