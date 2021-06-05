package product

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

type ErpProductSpecifics struct {
	apis.Api
}

// GetErpProductSpecificsList 获取在线产品列表
// @Summary 获取在线产品列表
// @Description 获取在线产品列表
// @Tags 在线产品
// @Param shopId query string false "店铺id"
// @Param marketplaceId query string false "市场id"
// @Param dxmState query string false "产品状态"
// @Param sku query string false "SKU"
// @Param title query string false "产品标题"
// @Param saleStartDate query string false "开始售出日期"
// @Param saleEndDate query string false "销售结束日期"
// @Param fnsku query string false "FNSKU"
// @Param commodityName query string false "商品名称"
// @Param siteName query string false "站点名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ErpProductSpecifics}} "{"code": 200, "data": [...]}"
// @Router /api/v1/product [get]
// @Security Bearer
func (e ErpProductSpecifics) GetErpProductSpecificsList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.ErpProductSpecificsSearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.ErpProductSpecifics, 0)
	var count int64
	serviceErpProductSpecifics := service.ErpProductSpecifics{}
	serviceErpProductSpecifics.Log = log
	serviceErpProductSpecifics.Orm = db
	err = serviceErpProductSpecifics.GetErpProductSpecificsPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get ErpProductSpecifics Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetErpProductSpecifics 获取在线产品
// @Summary 获取在线产品
// @Description 获取在线产品
// @Tags 在线产品
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.ErpProductSpecifics} "{"code": 200, "data": [...]}"
// @Router /api/v1/product/{id} [get]
// @Security Bearer
func (e ErpProductSpecifics) GetErpProductSpecifics(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpProductSpecificsById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.ErpProductSpecifics

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpProductSpecifics := service.ErpProductSpecifics{}
	serviceErpProductSpecifics.Log = log
	serviceErpProductSpecifics.Orm = db
	err = serviceErpProductSpecifics.GetErpProductSpecifics(control, p, &object)
	if err != nil {
		log.Errorf("Get ErpProductSpecifics error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertErpProductSpecifics 创建在线产品
// @Summary 创建在线产品
// @Description 创建在线产品
// @Tags 在线产品
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpProductSpecificsControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/product [post]
// @Security Bearer
func (e ErpProductSpecifics) InsertErpProductSpecifics(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpProductSpecificsControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpProductSpecifics model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceErpProductSpecifics := service.ErpProductSpecifics{}
	serviceErpProductSpecifics.Orm = db
	serviceErpProductSpecifics.Log = log
	err = serviceErpProductSpecifics.InsertErpProductSpecifics(object.(*models.ErpProductSpecifics))
	if err != nil {
		log.Errorf("Insert ErpProductSpecifics error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateErpProductSpecifics 修改在线产品
// @Summary 修改在线产品
// @Description 修改在线产品
// @Tags 在线产品
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpProductSpecificsControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/product/{id} [put]
// @Security Bearer
func (e ErpProductSpecifics) UpdateErpProductSpecifics(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpProductSpecificsControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpProductSpecifics model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpProductSpecifics := service.ErpProductSpecifics{}
	serviceErpProductSpecifics.Orm = db
	serviceErpProductSpecifics.Log = log
	err = serviceErpProductSpecifics.UpdateErpProductSpecifics(object.(*models.ErpProductSpecifics), p)
	if err != nil {
		log.Errorf("Update ErpProductSpecifics error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteErpProductSpecifics 删除在线产品
// @Summary 删除在线产品
// @Description 删除在线产品
// @Tags 在线产品
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/product [delete]
// @Security Bearer
func (e ErpProductSpecifics) DeleteErpProductSpecifics(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpProductSpecificsById)

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

	serviceErpProductSpecifics := service.ErpProductSpecifics{}
	serviceErpProductSpecifics.Orm = db
	serviceErpProductSpecifics.Log = log
	err = serviceErpProductSpecifics.RemoveErpProductSpecifics(control, p)
	if err != nil {
		log.Errorf("Remove ErpProductSpecifics error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}