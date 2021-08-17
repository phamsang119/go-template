package postgres

import (
	"fmt"
	"game-api/config"
	"game-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User UserRepository
	db   *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	driver := config.Env().GetPostgresDBDriver()
	host := config.Env().GetPostgresDBHost()
	password := config.Env().GetPostgresDBPassword()
	user := config.Env().GetPostgresDBUser()
	dbname := config.Env().GetPostgresDBName()
	port := config.Env().GetPostgresDBPort()
	sslMode := config.Env().GetPostgresDBSSLMode()
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", host, port, user, dbname, sslMode, password)
	db, err := gorm.Open(driver, dbUrl)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// Close closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// AutoMigrate This migrates all tables
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.User{}).Error
}
