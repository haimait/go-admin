package actions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/common/dto"
	"go-admin/common/log"
	"go-admin/common/models"
	"go-admin/tools"
	"go-admin/tools/app"
)

// IndexAction 通用查询动作
func IndexAction(m models.ActiveRecord, d dto.Index, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := tools.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := tools.GenerateMsgIDFromContext(c)
		list := f()
		object := m.Generate()
		req := d.Generate()
		var count int64

		//查询列表
		err = req.Bind(c)
		if err != nil {
			app.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.WithContext(c).Model(object).
			Scopes(
				dto.MakeCondition(req.GetNeedSearch()),
				dto.Paginate(req.GetPageSize(), req.GetPageIndex()),
				Permission(object.TableName(), p),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(&count).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("MsgID[%s] Index error: %s", msgID, err)
			app.Error(c, http.StatusInternalServerError, err, "查询失败")
			return
		}
		app.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
		c.Next()
	}
}
// IndexAction 通用查询动作
/*
m model对象
d 筛洗对象参数
f 返回数据方法
*/
func IndexAction2(m models.ActiveRecord2, d dto.Index, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := tools.GetOrm(c) //获取数据库连接对象
		if err != nil {
			log.Error(err)
			return
		}

		msgID := tools.GenerateMsgIDFromContext(c) //生成请求id
		list := f() //返回的列表,一般为 切片里放的model的对象
		object := m.Generate()  //ActiveRecord2对象
		req := d.Generate() //index对象
		var count int64

		//查询列表 获取筛洗对象参数
		err = req.Bind(c)
		if err != nil {
			app.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.WithContext(c).Model(object). //查哪张表
			Scopes(
				dto.MakeCondition(req.GetNeedSearch()), //筛选条件 req.GetNeedSearch()筛洗对象参数
				dto.Paginate(req.GetPageSize(), req.GetPageIndex()), //分页
				Permission(object.TableName(), p), //数据权限
			).
			Find(list).Limit(-1).Offset(-1).
			Count(&count).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("MsgID[%s] Index error: %s", msgID, err)
			app.Error(c, http.StatusInternalServerError, err, "查询失败")
			return
		}
		app.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
		c.Next()
	}
}
