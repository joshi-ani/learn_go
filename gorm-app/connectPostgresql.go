package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

// equals
type PII struct {
	ID           int64
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	Upvotes      int32
}

// gorm.Model definition columns
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Setting default Values
type User struct {
	ID   string `gorm:"default:uuid_generate_v3()"`
	Name string `gorm:"default:Jacob"`
	Age  int64  `gorm:"default:18"`
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=postgres password=root dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("connected to database")
		db.AutoMigrate(&Product{})
		db.AutoMigrate(&User{})
		db.Create(&Product{Code: "D44", Price: 100})
		println("Updated Product.")
		println("Created user.")
	}

}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	// if !u.IsValid() {
	// 	err = errors.New("can't save invalid data")
	// }
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == "1" {
		tx.Model(u).Update("role", "admin")
	}
	// if !u.IsValid() {
	// 	return errors.New("rollback invalid user")
	// }
	return
}
