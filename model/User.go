package model

import (
	"blog/utils/errmsg"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) (users []User) {
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) (code int) {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&User{}).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) (code int) {
	err = db.Where("id = ?",id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

// 密码加密
func ScryptPw(password string) (cryptPassword string) {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 23, 6, 15, 20}

	HashPassword, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	cryptPassword = base64.StdEncoding.EncodeToString(HashPassword)
	return
}
