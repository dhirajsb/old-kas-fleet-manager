package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	mocket "github.com/selvatico/go-mocket"
	"sync"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"
)

var singleton *ConnectionFactory
var once sync.Once

type ConnectionFactory struct {
	Config *config.DatabaseConfig
	DB     *gorm.DB
}

// NewConnectionFactory will initialize a singleton ConnectionFactory as needed and return the same instance.
// Go includes database connection pooling in the platform. Gorm uses the same and provides a method to
// clone a connection via New(), which is safe for use by concurrent Goroutines.
func NewConnectionFactory(config *config.DatabaseConfig) *ConnectionFactory {
	once.Do(func() {
		db, err := gorm.Open(config.Dialect, config.ConnectionString())
		if err != nil {
			panic(fmt.Sprintf(
				"failed to connect to %s database %s with connection string: %s\nError: %s",
				config.Dialect,
				config.Name,
				config.LogSafeConnectionString(),
				err.Error(),
			))
		}
		db.DB().SetMaxOpenConns(config.MaxOpenConnections)

		singleton = &ConnectionFactory{Config: config, DB: db}
	})
	return singleton
}

// NewMockConnectionFactory should only be used for defining mock database drivers
// This uses mocket under the hood, use the global mocket.Catcher to change how the database should respond to SQL
// queries
func NewMockConnectionFactory(dbConfig *config.DatabaseConfig) *ConnectionFactory {
	if dbConfig == nil {
		dbConfig = &config.DatabaseConfig{}
	}
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	mocketDB, err := gorm.Open(mocket.DriverName, "")
	if err != nil {
		panic(err)
	}
	connectionFactory := &ConnectionFactory{dbConfig, mocketDB}
	return connectionFactory
}

// New returns a new database connection
func (f *ConnectionFactory) New() *gorm.DB {
	if f.Config.Debug {
		return f.DB.New().Debug()
	}
	return f.DB.New()
}

// Checks to ensure a connection is present
func (f *ConnectionFactory) CheckConnection() error {
	return f.DB.Exec("SELECT 1").Error
}

// Close will close the connection to the database.
// THIS MUST **NOT** BE CALLED UNTIL THE SERVER/PROCESS IS EXITING!!
// This should only ever be called once for the entire duration of the application and only at the end.
func (f *ConnectionFactory) Close() error {
	return f.DB.Close()
}

// By default do no roll back transaction.
// only perform rollback if explicitly set by db.db.MarkForRollback(ctx, err)
const defaultRollbackPolicy = false

// TxFactory represents an sql transaction
type txFactory struct {
	rollbackFlag *rollbackFlag
	tx           *sql.Tx
	txid         int64
}

// rollbackFlag represents wether a transaction should be rolled back.
type rollbackFlag struct {
	val bool
}

// newTransaction constructs a new Transaction object.
func newTransaction() (*txFactory, error) {
	tx, err := singleton.DB.DB().Begin()
	if err != nil {
		return nil, err
	}

	var txid int64 = 0

	// current transaction ID set by postgres.  these are *not* distinct across time
	// and do get reset after postgres performs "vacuuming" to reclaim used IDs.
	row := tx.QueryRow("select txid_current()")
	if row != nil {
		err := row.Scan(&txid)
		if err != nil {
			return nil, err
		}
	}

	return &txFactory{
		tx:   tx,
		txid: txid,
		rollbackFlag: &rollbackFlag{
			val: defaultRollbackPolicy,
		},
	}, nil
}

// markedForRollback returns true if a transaction is flagged for rollback and false otherwise.
func (tx *txFactory) markedForRollback() bool {
	return tx.rollbackFlag.val
}
