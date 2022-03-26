package database

import (
	"fmt"
	"virtual-white-board-service/internal/flags"
	"virtual-white-board-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//New inits the database
func New(flags *flags.Flags) (*gorm.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", flags.PostgresUser, flags.PostgresPassword, flags.PostgresHost, flags.PostgresPort, flags.PostgresDbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = AutoMigrate(db)

	return db, nil
}

//AutoMigrate for creating tables in correct order
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Comment{})
	if err != nil {
		return err
	}
	return nil
}

//DropTables for dropping tables in database
func DropTables(db *gorm.DB) error {
	err := db.Exec("drop table comments;").Error
	if err != nil {
		return err
	}
	err = db.Exec("drop table posts;").Error
	if err != nil {
		return err
	}
	err = db.Exec("drop table users;").Error
	if err != nil {
		return err
	}
	return nil
}

//CreateTestUsers creates test users
func CreateTestUsers(db *gorm.DB) error {
	DropTables(db)
	AutoMigrate(db)
	user := &models.User{Username: "username", Password: "password"}

	err := db.Create(user).Error
	if err != nil {
		return err
	}

	post := &models.Post{Message: "message", UserID: user.ID}

	err = db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}
