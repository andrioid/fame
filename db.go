package fame

import (
	"database/sql"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	// Sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	db, err := gorm.Open("sqlite3", "fame.db")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&User{}, &Role{}, &Permission{}, &Auth{}, &Media{}, &Tag{})
}

// User explains itself
type User struct {
	gorm.Model
	Name       string
	Email      string
	Photo      string
	Auths      []Auth
	Roles      []Role `gorm:"many2many:user_roles"`
	ProfilePic string
}

// Role is used by Users
type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

// Permission is used by Roles
type Permission struct {
	gorm.Model
	Name string
}

// Auth stores OAUTH details for users
type Auth struct {
	gorm.Model
	UserID   int64
	External string
	Network  string
}

// MediaType is a constant
type MediaType int

const (
	// Photo is a media type
	Photo MediaType = iota
	// Video is a media type
	Video
)

// Media is used for all media
type Media struct {
	gorm.Model
	Title       string
	Description string
	Location    string
	Type        MediaType
	Tags        []Tag `gorm:"many2many:media_tags"`
	People      []User
	Owner       User
}

// Tag helps you find media
type Tag struct {
	gorm.Model
	Name string
}
