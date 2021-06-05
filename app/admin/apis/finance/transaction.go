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

type ErpFinanceTransaction struct {
	apis.Api
}

// GetErpFinanceTransactionList 获取结算明细列表
// @Summary 获取结算明细列表
// @Description 获取结算明细列表
// @Tags 结算明细
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ErpFinanceTransaction}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance [get]
// @Security Bearer
func (e ErpFinanceTransaction) GetErpFinanceTransactionList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.ErpFinanceTransactionSearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.ErpFinanceTransaction, 0)
	var count int64
	serviceErpFinanceTransaction := service.ErpFinanceTransaction{}
	serviceErpFinanceTransaction.Log = log
	serviceErpFinanceTransaction.Orm = db
	err = serviceErpFinanceTransaction.GetErpFinanceTransactionPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get ErpFinanceTransaction Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetErpFinanceTransaction 获取结算明细
// @Summary 获取结算明细
// @Description 获取结算明细
// @Tags 结算明细
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.ErpFinanceTransaction} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance/{id} [get]
// @Security Bearer
func (e ErpFinanceTransaction) GetErpFinanceTransaction(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceTransactionById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.ErpFinanceTransaction

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceTransaction := service.ErpFinanceTransaction{}
	serviceErpFinanceTransaction.Log = log
	serviceErpFinanceTransaction.Orm = db
	err = serviceErpFinanceTransaction.GetErpFinanceTransaction(control, p, &object)
	if err != nil {
		log.Errorf("Get ErpFinanceTransaction error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertErpFinanceTransaction 创建结算明细
// @Summary 创建结算明细
// @Description 创建结算明细
// @Tags 结算明细
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceTransactionControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance [post]
// @Security Bearer
func (e ErpFinanceTransaction) InsertErpFinanceTransaction(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceTransactionControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceTransaction model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceErpFinanceTransaction := service.ErpFinanceTransaction{}
	serviceErpFinanceTransaction.Orm = db
	serviceErpFinanceTransaction.Log = log
	err = serviceErpFinanceTransaction.InsertErpFinanceTransaction(object.(*models.ErpFinanceTransaction))
	if err != nil {
		log.Errorf("Insert ErpFinanceTransaction error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateErpFinanceTransaction 修改结算明细
// @Summary 修改结算明细
// @Description 修改结算明细
// @Tags 结算明细
// @Accept application/json
// @Product application/json
// @Param data body dto.ErpFinanceTransactionControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance/{id} [put]
// @Security Bearer
func (e ErpFinanceTransaction) UpdateErpFinanceTransaction(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceTransactionControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate ErpFinanceTransaction model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceErpFinanceTransaction := service.ErpFinanceTransaction{}
	serviceErpFinanceTransaction.Orm = db
	serviceErpFinanceTransaction.Log = log
	err = serviceErpFinanceTransaction.UpdateErpFinanceTransaction(object.(*models.ErpFinanceTransaction), p)
	if err != nil {
		log.Errorf("Update ErpFinanceTransaction error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteErpFinanceTransaction 删除结算明细
// @Summary 删除结算明细
// @Description 删除结算明细
// @Tags 结算明细
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance [delete]
// @Security Bearer
func (e ErpFinanceTransaction) DeleteErpFinanceTransaction(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.ErpFinanceTransactionById)

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

	serviceErpFinanceTransaction := service.ErpFinanceTransaction{}
	serviceErpFinanceTransaction.Orm = db
	serviceErpFinanceTransaction.Log = log
	err = serviceErpFinanceTransaction.RemoveErpFinanceTransaction(control, p)
	if err != nil {
		log.Errorf("Remove ErpFinanceTransaction error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}