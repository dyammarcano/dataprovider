package dataprovider

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// MySQLProvider defines the auth provider for MySQL/MariaDB database
type MySQLProvider struct {
	dbHandle *sqlx.DB
}

func (m *MySQLProvider) Disconnect() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) GetConnection() *sqlx.DB {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) CheckAvailability() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) ReconnectDatabase() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) InitializeDatabase() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) migrateDatabase() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) RevertDatabase(targetVersion int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLProvider) ResetDatabase() error {
	//TODO implement me
	panic("implement me")
}

func newMySQLProvider(ctx context.Context) error {
	ctxValue := ctx.Value("config").(*ConfigModule)
	if ctxValue == nil {
		return fmt.Errorf("config not found in context")
	}

	dbHandle, err := sqlx.Connect("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		return err
	}

	if err = dbHandle.PingContext(ctx); err != nil {
		return err
	}

	provider = &MySQLProvider{dbHandle: dbHandle}

	return nil
}
