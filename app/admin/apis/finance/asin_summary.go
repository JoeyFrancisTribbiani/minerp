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

type ErpFinanceAsinSummary struct {
	apis.Api
}

// GetErpFinanceAsinSummaryList 获取ASIN月费列表
// @Summary 获取ASIN月费列表
// @Description 获取ASIN月费列表
// @Tags ASIN月费
// @Param commoditySku query string false "商品SKU"
// @Param commodityName query string false "商品名称"
// @Param shopId query string false "店铺id"
// @Param marketplaceId query string false "市场id"
// @Param currency query string false "币种"
// @Param asin query string false "ASIN"
// @Param sku query string false "SKU"
// @Param month query string false "月份"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ErpFinanceAsinSummary}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance [get]
// @Security Bearer
func (e ErpFinanceAsinSummary) GetErpFinanceAsinSummaryList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.ErpFinanceAsinSummarySearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.ErpFinanceAsinSummary, 0)
	var count int64
	serviceErpFinanceAsinSummary := service.ErpFinanceAsinSummary{}
	serviceErpFinanceAsinSummary.Log = log
	serviceErpFinanceAsinSummary.Orm = db
	err = serviceErpFinanceAsinSummary.GetErpFinanceAsinSummaryPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get ErpFinanceAsinSummary Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetErpFinanceAsinSummary 获取ASIN月费
// @Summary 获取ASIN月费
// @Description 获取ASIN月费
// @Tags ASIN月费
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.ErpFinanceAsinSummary} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance/{id} [get]
// @Security Bearer
func (e ErpFinanceAsinSummary) GetErpFinanceAsinSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceAsinSummaryById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.ErpFinanceAsinSummary

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceAsinSummary := service.ErpFinanceAsinSummary{}
	serviceErpFinanceAsinSummary.Log = log
	serviceErpFinanceAsinSummary.Orm = db
	err = serviceErpFinanceAsinSummary.GetErpFinanceAsinSummary(control, p, &object)
	if err != nil {
		log.Errorf("Get ErpFinanceAsinSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertErpFinanceAsinSummary 创建ASIN月费
// @Summary 创建ASIN月费
// @Description 创建ASIN月费
// @Tags ASIN月费
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceAsinSummaryControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance [post]
// @Security Bearer
func (e ErpFinanceAsinSummary) InsertErpFinanceAsinSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceAsinSummaryControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceAsinSummary model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceErpFinanceAsinSummary := service.ErpFinanceAsinSummary{}
	serviceErpFinanceAsinSummary.Orm = db
	serviceErpFinanceAsinSummary.Log = log
	err = serviceErpFinanceAsinSummary.InsertErpFinanceAsinSummary(object.(*models.ErpFinanceAsinSummary))
	if err != nil {
		log.Errorf("Insert ErpFinanceAsinSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateErpFinanceAsinSummary 修改ASIN月费
// @Summary 修改ASIN月费
// @Description 修改ASIN月费
// @Tags ASIN月费
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceAsinSummaryControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance/{id} [put]
// @Security Bearer
func (e ErpFinanceAsinSummary) UpdateErpFinanceAsinSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceAsinSummaryControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceAsinSummary model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceAsinSummary := service.ErpFinanceAsinSummary{}
	serviceErpFinanceAsinSummary.Orm = db
	serviceErpFinanceAsinSummary.Log = log
	err = serviceErpFinanceAsinSummary.UpdateErpFinanceAsinSummary(object.(*models.ErpFinanceAsinSummary), p)
	if err != nil {
		log.Errorf("Update ErpFinanceAsinSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteErpFinanceAsinSummary 删除ASIN月费
// @Summary 删除ASIN月费
// @Description 删除ASIN月费
// @Tags ASIN月费
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance [delete]
// @Security Bearer
func (e ErpFinanceAsinSummary) DeleteErpFinanceAsinSummary(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceAsinSummaryById)

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

	serviceErpFinanceAsinSummary := service.ErpFinanceAsinSummary{}
	serviceErpFinanceAsinSummary.Orm = db
	serviceErpFinanceAsinSummary.Log = log
	err = serviceErpFinanceAsinSummary.RemoveErpFinanceAsinSummary(control, p)
	if err != nil {
		log.Errorf("Remove ErpFinanceAsinSummary error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}