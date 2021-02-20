package system

import (
	"fmt"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-admin/app/admin/models"
	"go-admin/tools"
)

type SysSetting struct {
	apis.Api
}

// @Summary 查询系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/setting [get]
func (e *SysSetting) GetSetting(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	sysSettingService := service.SysSetting{}
	sysSettingService.MsgID=msgID
	sysSettingService.Orm=db
	var model = models.SysSetting{}
	err = sysSettingService.GetSysSetting(&model)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	if model.Logo != "" {
		if !strings.HasPrefix(model.Logo, "http") {
			model.Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, model.Logo)
		}
	}

	e.OK(c, model, "查询成功")
}

// @Summary 更新或提交系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Param data body models.SysUser true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/system/setting [post]
func (e *SysSetting) CreateOrUpdateSetting(c *gin.Context) {
	control := new(dto.SysSettingControl)
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	//更新操作
	err = control.Bind(c)
	control.Logo = utils.FilterToLocalUrl(control.Logo) //过滤为本地的相对路径
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.Generate()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	sysSettingService := service.SysSetting{}
	sysSettingService.MsgID=msgID
	sysSettingService.Orm=db
	err = sysSettingService.UpdateSysSetting(object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "更新失败")
		return
	}


	if object.Logo != "" {
		if !strings.HasPrefix(object.Logo, "http") {
			object.Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, object.Logo)
		}
	}
	e.OK(c, object, "提交成功")
}

