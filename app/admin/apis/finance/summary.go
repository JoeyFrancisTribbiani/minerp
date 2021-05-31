package finance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
)

type ErpFinanceSummary struct {
	apis.Api
}

// GetErpFinanceSummaryList 获取店铺汇总列表
// @Summary 获取店铺汇总列表
// @Description 获取店铺汇总列表
// @Tags 店铺汇总
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ErpFinanceSummary}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance [get]
// @Security Bearer
func (e ErpFinanceSummary) GetErpFinanceSummaryList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.ErpFinanceSummarySearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.ErpFinanceSummary, 0)
	var count int64
	serviceErpFinanceSummary := service.ErpFinanceSummary{}
	serviceErpFinanceSummary.Log = log
	serviceErpFinanceSummary.Orm = db
	err = serviceErpFinanceSummary.GetErpFinanceSummaryPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get ErpFinanceSummary Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetErpFinanceSummary 获取店铺汇总
// @Summary 获取店铺汇总
// @Description 获取店铺汇总
// @Tags 店铺汇总
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.ErpFinanceSummary} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance/{id} [get]
// @Security Bearer
func (e ErpFinanceSummary) GetErpFinanceSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceSummaryById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.ErpFinanceSummary

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceSummary := service.ErpFinanceSummary{}
	serviceErpFinanceSummary.Log = log
	serviceErpFinanceSummary.Orm = db
	err = serviceErpFinanceSummary.GetErpFinanceSummary(control, p, &object)
	if err != nil {
		log.Errorf("Get ErpFinanceSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertErpFinanceSummary 创建店铺汇总
// @Summary 创建店铺汇总
// @Description 创建店铺汇总
// @Tags 店铺汇总
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceSummaryControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance [post]
// @Security Bearer
func (e ErpFinanceSummary) InsertErpFinanceSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceSummaryControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceSummary model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceErpFinanceSummary := service.ErpFinanceSummary{}
	serviceErpFinanceSummary.Orm = db
	serviceErpFinanceSummary.Log = log
	err = serviceErpFinanceSummary.InsertErpFinanceSummary(object.(*models.ErpFinanceSummary))
	if err != nil {
		log.Errorf("Insert ErpFinanceSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateErpFinanceSummary 修改店铺汇总
// @Summary 修改店铺汇总
// @Description 修改店铺汇总
// @Tags 店铺汇总
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceSummaryControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance/{id} [put]
// @Security Bearer
func (e ErpFinanceSummary) UpdateErpFinanceSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceSummaryControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceSummary model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceSummary := service.ErpFinanceSummary{}
	serviceErpFinanceSummary.Orm = db
	serviceErpFinanceSummary.Log = log
	err = serviceErpFinanceSummary.UpdateErpFinanceSummary(object.(*models.ErpFinanceSummary), p)
	if err != nil {
		log.Errorf("Update ErpFinanceSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteErpFinanceSummary 删除店铺汇总
// @Summary 删除店铺汇总
// @Description 删除店铺汇总
// @Tags 店铺汇总
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance [delete]
// @Security Bearer
func (e ErpFinanceSummary) DeleteErpFinanceSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceSummaryById)

	//删除操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	// 设置编辑人
	control.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceSummary := service.ErpFinanceSummary{}
	serviceErpFinanceSummary.Orm = db
	serviceErpFinanceSummary.Log = log
	err = serviceErpFinanceSummary.RemoveErpFinanceSummary(control, p)
	if err != nil {
		log.Errorf("Remove ErpFinanceSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}