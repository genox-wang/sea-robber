package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// User is user model
type User struct {
	ID           int64     `json:"id"`
	UUID         string    `json:"uuid" gorm:"unique;not null"`
	DisplayName  string    `json:"display_name"`
	BattleScore  int64     `json:"battle_score" gorm:"not null"`
	BattleShipID int       `json:"battle_ship_id" gorm:"not null"`
	Friends      []*User   `json:"friends" gorm:"many2many:user_friends;association_jointable_foreignkey:friend_id;ASSOCIATION_SAVE_REFERENCE:true;ASSOCIATION_AUTOUPDATE:false;"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	DeletedAt    time.Time `json:"-" gorm:"default:null"`
}

// BeforeCreate add uuid before user created
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	UUID := fmt.Sprintf("%s", uuid.NewV1())
	scope.SetColumn("UUID", UUID)
	logrus.Warnf("[ User Create ] => uuid = %s", UUID)
	return nil
}

// Create create a new user
func (u *User) Create() error {
	return DB.Create(u).Error
}

//Update update user
func (u *User) Update() error {
	if err := DB.Model(u).Update(u).Error; err != nil {
		return err
	}
	return nil
}

// GetAll get all users
func (*User) GetAll() ([]*User, error) {
	users := make([]*User, 0)
	err := DB.Find(&users).Error
	return users, err
}

// GetAllFriends get all friends
func (u *User) GetAllFriends() ([]*User, error) {
	err := DB.Preload("Servers").First(u).Error
	return u.Friends, err
}
