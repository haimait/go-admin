package models

import (
	"go-admin/common/models"
	"time"
)

type Casusers struct {
	ModelBModel
	CasUid              int64     `json:"casUid" gorm:"type:int(11);comment:cas_uid"`                                                  //
	Isadmin             string    `json:"isadmin" gorm:"type:tinyint(1);comment:是否为后台管理员 0:否 1:是"`                                     //
	Username            string    `json:"username" gorm:"type:varchar(255);comment:用户名"`                                               //
	Nickname            string    `json:"nickname" gorm:"type:varchar(200);comment:昵称"`                                                //
	Email               string    `json:"email" gorm:"type:varchar(255);comment:Email"`                                                //
	AvatarMedium        string    `json:"avatarMedium" gorm:"type:varchar(255);comment:用户头像中尺寸"`                                       //
	AvatarThumb         string    `json:"avatarThumb" gorm:"type:varchar(255);comment:用户头像缩略图"`                                        //
	Level               string    `json:"level" gorm:"type:int(11);comment:暂时没有用"`                                                     //
	Key                 string    `json:"key" gorm:"type:varchar(255);comment:Key"`                                                    //
	Isvip               string    `json:"isvip" gorm:"type:tinyint(1);comment:是否为会员 0:否 1:是"`                                          //
	FofaCoin            int64     `json:"fofaCoin" gorm:"type:int(11);comment:F币 已经包含OnlyPayFofaCoin"`                                 //
	Point               int64     `json:"point" gorm:"type:int(11);comment:积分"`                                                        //
	IsAuthenticated     string    `json:"isAuthenticated" gorm:"type:tinyint(1);comment:是否认证,好像没有用"`                                   //
	CoinState           string    `json:"coinState" gorm:"type:int(11);comment:F币状态"`                                                  //
	VipLevel            string    `json:"vipLevel" gorm:"type:int(11);default:0;comment:会员等级 1:普通会员 2:高级会员 3:企业会员"`                    //
	OnlyPayFofaCoin     int64     `json:"onlyPayFofaCoin" gorm:"type:int(11);comment:不能提现F币"`                                          //
	OnlyPayCoin         int64     `json:"onlyPayCoin" gorm:"type:int(11);comment:好像没有用"`                                               //
	VipDuration         time.Time `json:"vipDuration" gorm:"type:datetime;comment:暂时没有用"`                                              //
	RoleId              int64     `json:"roleId" gorm:"type:int(11);comment:RoleId"`                                                   //
	EncryptionCoin      string    `json:"encryptionCoin" gorm:"type:varchar(255);comment:F币加密码串"`                                      //
	AccountState        string    `json:"accountState" gorm:"type:varchar(255);comment:对账状态 success:正常 failure:异常"`                    //
	Category            string    `json:"category" gorm:"type:varchar(255);comment:用户类型   user:用户 common_admin:管理员 super_admin:超级管理员"` //
	EncryptedPassword   string    `json:"encryptedPassword" gorm:"type:varchar(255);comment:EncryptedPassword"`                        //
	ResetPasswordToken  string    `json:"resetPasswordToken" gorm:"type:varchar(255);comment:ResetPasswordToken"`                      //
	ResetPasswordSentAt time.Time `json:"resetPasswordSentAt" gorm:"type:datetime;comment:ResetPasswordSentAt"`                        //
	RememberCreatedAt   time.Time `json:"rememberCreatedAt" gorm:"type:datetime;comment:RememberCreatedAt"`                            //
	SignInCount         int64     `json:"signInCount" gorm:"type:int(11);comment:SignInCount"`                                         //
	CurrentSignInAt     time.Time `json:"currentSignInAt" gorm:"type:datetime;comment:CurrentSignInAt"`                                //
	LastSignInAt        time.Time `json:"lastSignInAt" gorm:"type:datetime;comment:LastSignInAt"`                                      //
	CurrentSignInIp     string    `json:"currentSignInIp" gorm:"type:varchar(255);comment:CurrentSignInIp"`                            //
	LastSignInIp        string    `json:"lastSignInIp" gorm:"type:varchar(255);comment:LastSignInIp"`                                  //
	AvatarFileName      string    `json:"avatarFileName" gorm:"type:varchar(255);comment:AvatarFileName"`                              //
	AvatarContentType   string    `json:"avatarContentType" gorm:"type:varchar(255);comment:AvatarContentType"`                        //
	AvatarFileSize      int       `json:"avatarFileSize" gorm:"type:int(11);comment:AvatarFileSize"`                                   //
	AvatarUpdatedAt     time.Time `json:"avatarUpdatedAt" gorm:"type:datetime;comment:AvatarUpdatedAt"`                                //
	SashId              int       `json:"sashId" gorm:"type:int(11);comment:SashId"`                                                   //
	Duration            time.Time `json:"duration" gorm:"type:datetime;comment:会员过期时间"`                                                //
	DueTime             int64     `json:"dueTime" gorm:"type:int(11);comment:DueTime"`                                                 //
	ConfirmedAt         time.Time `json:"confirmedAt" gorm:"type:datetime;comment:ConfirmedAt"`                                        //
	ConfirmationToken   string    `json:"confirmationToken" gorm:"type:varchar(255);comment:ConfirmationToken"`                        //
	ConfirmationSentAt  time.Time `json:"confirmationSentAt" gorm:"type:datetime;comment:ConfirmationSentAt"`                          //
	UnconfirmedEmail    string    `json:"unconfirmedEmail" gorm:"type:varchar(255);comment:UnconfirmedEmail"`                          //
	FailedAttempts      string    `json:"failedAttempts" gorm:"type:int(11);comment:FailedAttempts"`                                   //
	UnlockToken         string    `json:"unlockToken" gorm:"type:varchar(255);comment:UnlockToken"`                                    //
	LockedAt            time.Time `json:"lockedAt" gorm:"type:datetime;comment:用户锁定时间 锁定用户不能调用api"`                                    //
	Ltime               int64     `json:"ltime" gorm:"type:int(11);comment:暂时没有用"`                                                     //
	Loginlog            string    `json:"loginlog" gorm:"type:text;comment:Loginlog"`                                                  //
	Adminlevel          int       `json:"adminlevel" gorm:"type:int(11);comment:Adminlevel"`                                           //
	EntIpLimit          int       `json:"entIpLimit" gorm:"type:bigint(20);comment:限制登陆次数 -1:不限制"`                                     //
}

func (Casusers) TableName() string {
	return "casusers"
}

func (e *Casusers) Generate() models.ActiveRecord2 {
	o := *e
	return &o
}

func (e *Casusers) GetId() interface{} {
	return e.ID
}
