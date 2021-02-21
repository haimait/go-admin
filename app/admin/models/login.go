package models

import (
	"errors"
	orm "go-admin/common/global"
	"go-admin/tools"
	"time"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, e error) {

	e = orm.Eloquent.Table("sys_user").
		Joins("Casusers").
		Where("sys_user.username = ? ", u.Username).
		Find(&user).Error
	if e != nil {
		return
	}
	_, e = tools.CompareHashAndPassword(user.Password, u.Password)
	if e != nil {
		return
	}
	//判断企业用户是否过期
	if user.RoleId != 1 && (user.Casusers.ID==0 ||
		user.Casusers.Isvip != "1" ||
		user.Casusers.VipLevel != "3" ||
		user.Casusers.Duration.Before(time.Now())){
		e = errors.New("vip time overtime")
		return
	}
	e = orm.Eloquent.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if e != nil {
		return
	}
	return
}

//func (u *Login) GetUser() (user SysUser, role SysRole, e error) {
//
//	e = orm.Eloquent.Table("sys_user").Where("username = ?  and status = 2", u.Username).Find(&user).Error
//	if e != nil {
//		return
//	}
//	_, e = tools.CompareHashAndPassword(user.Password, u.Password)
//	if e != nil {
//		return
//	}
//	e = orm.Eloquent.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
//	if e != nil {
//		return
//	}
//	return
//}
