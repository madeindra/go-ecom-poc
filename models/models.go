package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type (
	UserModelImpl interface {
		FindByID(id int) Users
		FindAll() []Users
	}

	Users struct {
		ID       uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
		Email    string `json:"email" gorm:"type:varchar(255);"`
		Username string `json:"username" gorm:"type:varchar(255);"`
		Name     string `json:"name" gorm:"type:varchar(255);"`
		Address  string `json:"address" gorm:"type:varchar(255);"`
		Phone    string `json:"phone" gorm:"type:varchar(255);"`
		Password string `json:"password" gorm:"type:varchar(255);"`
		Image    string `json:"image" gorm:"type:varchar(255);"`
	}

	UserModel struct {
		db *gorm.DB
	}
)
type Carts struct {
	ID        uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Status    string `json:"status" gorm:"type:varchar(255);"`
	Price     uint   `json:"price"`
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	Qty       uint   `json:"qty"`
}

type Categories struct {
	ID   uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"type:varchar(255);"`
}

type Invoices struct {
	ID      uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Total   uint   `json:"total"`
	UserID  uint   `json:"user_id"`
	Address string `json:"address" gorm:"type:varchar(255);"`
	Status  string `json:"status" gorm:"type:varchar(255);"`
	PromoID uint   `json:"promo_id"`
	Method  string `json:"method" gorm:"type:varchar(255);"`
}

type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type Login struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type Products struct {
	ID          uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"type:varchar(255);"`
	Description string `json:"description" gorm:"type:varchar(255);"`
	Price       string `json:"price" gorm:"type:varchar(255);"`
	StoreID     uint   `json:"store_id"`
	CategoryID  uint   `json:"category_id"`
}

type Promos struct {
	ID        uint       `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Value     float64    `json:"value"`
	BeginDate *time.Time `json:"begin_date"`
	EndDate   *time.Time `json:"end_date"`
	PromoCode string     `json:"promo_code" gorm:"type:varchar(255);"`
}

type Stores struct {
	ID          uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserID      uint   `json:"user_id"`
	Name        uint   `json:"name"`
	Description uint   `json:"description"`
	Image       string `json:"image"`
}

// type Users struct {
// 	ID       uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
// 	Email    string `json:"email" gorm:"type:varchar(255);"`
// 	Username string `json:"username" gorm:"type:varchar(255);"`
// 	Name     string `json:"name" gorm:"type:varchar(255);"`
// 	Address  string `json:"address" gorm:"type:varchar(255);"`
// 	Phone    string `json:"phone" gorm:"type:varchar(255);"`
// 	Password string `json:"password" gorm:"type:varchar(255);"`
// 	Image    string `json:"image" gorm:"type:varchar(255);"`
// }

type Wishilists struct {
	ID        uint `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db}
}

func (u *UserModel) FindByID(id int) Users {
	users := Users{}
	u.db.Find(&users, id)

	return users
}

func (u *UserModel) FindAll() []Users {
	users := []Users{}
	u.db.Find(&users)

	return users
}
