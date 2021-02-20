package dto

import (
	"github.com/gin-gonic/gin"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)




/**
 * 	search 生成where条件参数
 * 	exact / iexact 等于
 * 	contains / icontains 包含
 *	gt / gte 大于 / 大于等于
 *	lt / lte 小于 / 小于等于
 *	startswith / istartswith 以…起始
 *	endswith / iendswith 以…结束
 *	in
 *	isnull
 *  order 排序		e.g. order[key]=desc     order[key]=asc
*/
type CasusersSearch struct {
	dto.Pagination `search:"-"`
	//Isadmin        string `form:"isadmin" search:"type:exact;column:isadmin;table:casusers" comment:""`
	//Nickname string `form:"nickname" search:"type:contains;column:nickname;table:casusers" comment:""`
	Username string `form:"username" search:"type:contains;column:username;table:casusers" comment:""`
	Email string `form:"email" search:"type:contains;column:email;table:casusers" comment:""`
	AccountState string `form:"accountState" search:"type:exact;column:account_state;table:casusers" comment:""`
}

func (m *CasusersSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CasusersSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *CasusersSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CasusersControl struct {
	ID                  uint      `uri:"ID" comment:""` //
	CasUid              int64     `json:"casUid" comment:""`
	Isadmin             string    `json:"isadmin" comment:""`
	Username            string    `json:"username" comment:""`
	Nickname            string    `json:"nickname" comment:""`
	Email               string    `json:"email" comment:""`
	AvatarMedium        string    `json:"avatarMedium" comment:""`
	AvatarThumb         string    `json:"avatarThumb" comment:""`
	Level               string    `json:"level" comment:""`
	Key                 string    `json:"key" comment:""`
	Isvip               string    `json:"isvip" comment:""`
	FofaCoin            int64     `json:"fofaCoin" comment:""`
	Point               int64     `json:"point" comment:""`
	IsAuthenticated     string    `json:"isAuthenticated" comment:""`
	CoinState           string    `json:"coinState" comment:""`
	VipLevel            string    `json:"vipLevel" comment:""`
	OnlyPayFofaCoin     int64     `json:"onlyPayFofaCoin" comment:""`
	OnlyPayCoin         int64     `json:"onlyPayCoin" comment:""`
	VipDuration         time.Time `json:"vipDuration" comment:""`
	RoleId              int64     `json:"roleId" comment:""`
	EncryptionCoin      string    `json:"encryptionCoin" comment:""`
	AccountState        string    `json:"accountState" comment:""`
	Category            string    `json:"category" comment:""`
	EncryptedPassword   string    `json:"encryptedPassword" comment:""`
	ResetPasswordToken  string    `json:"resetPasswordToken" comment:""`
	ResetPasswordSentAt time.Time `json:"resetPasswordSentAt" comment:""`
	RememberCreatedAt   time.Time `json:"rememberCreatedAt" comment:""`
	SignInCount         int64     `json:"signInCount" comment:""`
	CurrentSignInAt     time.Time `json:"currentSignInAt" comment:""`
	LastSignInAt        time.Time `json:"lastSignInAt" comment:""`
	CurrentSignInIp     string    `json:"currentSignInIp" comment:""`
	LastSignInIp        string    `json:"lastSignInIp" comment:""`
	AvatarFileName      string    `json:"avatarFileName" comment:""`
	AvatarContentType   string    `json:"avatarContentType" comment:""`
	AvatarFileSize      int       `json:"avatarFileSize" comment:""`
	AvatarUpdatedAt     time.Time `json:"avatarUpdatedAt" comment:""`
	SashId              int       `json:"sashId" comment:""`
	Duration            time.Time `json:"duration" comment:""`
	DueTime             int64     `json:"dueTime" comment:""`
	ConfirmedAt         time.Time `json:"confirmedAt" comment:""`
	ConfirmationToken   string    `json:"confirmationToken" comment:""`
	ConfirmationSentAt  time.Time `json:"confirmationSentAt" comment:""`
	UnconfirmedEmail    string    `json:"unconfirmedEmail" comment:""`
	FailedAttempts      string    `json:"failedAttempts" comment:""`
	UnlockToken         string    `json:"unlockToken" comment:""`
	LockedAt            time.Time `json:"lockedAt" comment:""`
	Ltime               int64     `json:"ltime" comment:""`
	Loginlog            string    `json:"loginlog" comment:""`
	Adminlevel          int       `json:"adminlevel" comment:""`
	EntIpLimit          int       `json:"entIpLimit" comment:""`
}

func (s *CasusersControl) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
	}
	return err
}

func (s *CasusersControl) Generate() dto.Control2 {
	cp := *s
	return &cp
}

func (s *CasusersControl) GenerateM() (common.ActiveRecord2, error) {
	return &models.Casusers{
		//Model:               gorm.Model{ID: s.ID},
		ModelBModel:         models.ModelBModel{ID: s.ID},
		CasUid:              s.CasUid,
		Isadmin:             s.Isadmin,
		Username:            s.Username,
		Nickname:            s.Nickname,
		Email:               s.Email,
		AvatarMedium:        s.AvatarMedium,
		AvatarThumb:         s.AvatarThumb,
		Level:               s.Level,
		Key:                 s.Key,
		Isvip:               s.Isvip,
		FofaCoin:            s.FofaCoin,
		Point:               s.Point,
		IsAuthenticated:     s.IsAuthenticated,
		CoinState:           s.CoinState,
		VipLevel:            s.VipLevel,
		OnlyPayFofaCoin:     s.OnlyPayFofaCoin,
		OnlyPayCoin:         s.OnlyPayCoin,
		VipDuration:         s.VipDuration,
		RoleId:              s.RoleId,
		EncryptionCoin:      s.EncryptionCoin,
		AccountState:        s.AccountState,
		Category:            s.Category,
		EncryptedPassword:   s.EncryptedPassword,
		ResetPasswordToken:  s.ResetPasswordToken,
		ResetPasswordSentAt: s.ResetPasswordSentAt,
		RememberCreatedAt:   s.RememberCreatedAt,
		SignInCount:         s.SignInCount,
		CurrentSignInAt:     s.CurrentSignInAt,
		LastSignInAt:        s.LastSignInAt,
		CurrentSignInIp:     s.CurrentSignInIp,
		LastSignInIp:        s.LastSignInIp,
		AvatarFileName:      s.AvatarFileName,
		AvatarContentType:   s.AvatarContentType,
		AvatarFileSize:      s.AvatarFileSize,
		AvatarUpdatedAt:     s.AvatarUpdatedAt,
		SashId:              s.SashId,
		Duration:            s.Duration,
		DueTime:             s.DueTime,
		ConfirmedAt:         s.ConfirmedAt,
		ConfirmationToken:   s.ConfirmationToken,
		ConfirmationSentAt:  s.ConfirmationSentAt,
		UnconfirmedEmail:    s.UnconfirmedEmail,
		FailedAttempts:      s.FailedAttempts,
		UnlockToken:         s.UnlockToken,
		LockedAt:            s.LockedAt,
		Ltime:               s.Ltime,
		Loginlog:            s.Loginlog,
		Adminlevel:          s.Adminlevel,
		EntIpLimit:          s.EntIpLimit,
	}, nil
}

func (s *CasusersControl) GetId() interface{} {
	return s.ID
}

type CasusersById struct {
	dto.ObjectById
}

func (s *CasusersById) Generate() dto.Control2 {
	cp := *s
	return &cp
}

func (s *CasusersById) GenerateM() (common.ActiveRecord2, error) {
	return &models.Casusers{}, nil
}
