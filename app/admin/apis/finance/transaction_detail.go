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

type TransactionDetail struct {
	apis.Api
}

// GetTransactionDetailList 获取结算明细列表
// @Summary 获取结算明细列表
// @Description 获取结算明细列表
// @Tags 结算明细
// @Param shopId query string false "店铺id"
// @Param shopName query string false "店铺名字"
// @Param type query string false "交易类型"
// @Param sku query string false "MSKU"
// @Param marketplace query string false "销售市场"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TransactionDetail}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance [get]
// @Security Bearer
func (e TransactionDetail) GetTransactionDetailList(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	d := new(dto.TransactionDetailSearch)
	//查询列表
	err = d.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.TransactionDetail, 0)
	var count int64
	serviceTransactionDetail := service.TransactionDetail{}
	serviceTransactionDetail.Log = log
	serviceTransactionDetail.Orm = db
	err = serviceTransactionDetail.GetTransactionDetailPage(d, p, &list, &count)
	if err != nil {
		log.Errorf("Get TransactionDetail Page error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK( list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetTransactionDetail 获取结算明细
// @Summary 获取结算明细
// @Description 获取结算明细
// @Tags 结算明细
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.TransactionDetail} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance/{id} [get]
// @Security Bearer
func (e TransactionDetail) GetTransactionDetail(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TransactionDetailById)

	//查看详情
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.TransactionDetail

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceTransactionDetail := service.TransactionDetail{}
	serviceTransactionDetail.Log = log
	serviceTransactionDetail.Orm = db
	err = serviceTransactionDetail.GetTransactionDetail(control, p, &object)
	if err != nil {
		log.Errorf("Get TransactionDetail error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK( object, "查看成功")
}

// InsertTransactionDetail 创建结算明细
// @Summary 创建结算明细
// @Description 创建结算明细
// @Tags 结算明细
// @Accept application/json
// @Product application/json
// @Param data body dto.TransactionDetailControl true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance [post]
// @Security Bearer
func (e TransactionDetail) InsertTransactionDetail(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TransactionDetailControl)

	//新增操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate TransactionDetail model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(user.GetUserId(c))

	serviceTransactionDetail := service.TransactionDetail{}
	serviceTransactionDetail.Orm = db
	serviceTransactionDetail.Log = log
	err = serviceTransactionDetail.InsertTransactionDetail(object.(*models.TransactionDetail))
	if err != nil {
		log.Errorf("Insert TransactionDetail error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK( object.GetId(), "创建成功")
}

// UpdateTransactionDetail 修改结算明细
// @Summary 修改结算明细
// @Description 修改结算明细
// @Tags 结算明细
// @Accept application/json
// @Product application/json
// @Param data body dto.TransactionDetailControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance/{id} [put]
// @Security Bearer
func (e TransactionDetail) UpdateTransactionDetail(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TransactionDetailControl)

	//更新操作
	err = control.Bind(c)
	if err != nil {
	    log.Warnf("request body bind error, %s", err.Error())
		e.Error( http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.GenerateM()
	if err != nil {
		log.Errorf("generate TransactionDetail model error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceTransactionDetail := service.TransactionDetail{}
	serviceTransactionDetail.Orm = db
	serviceTransactionDetail.Log = log
	err = serviceTransactionDetail.UpdateTransactionDetail(object.(*models.TransactionDetail), p)
	if err != nil {
		log.Errorf("Update TransactionDetail error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK( object.GetId(), "更新成功")
}

// DeleteTransactionDetail 删除结算明细
// @Summary 删除结算明细
// @Description 删除结算明细
// @Tags 结算明细
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance [delete]
// @Security Bearer
func (e TransactionDetail) DeleteTransactionDetail(c *gin.Context) {
    e.SetContext(c)
    log := e.GetLogger()
    db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}

	control := new(dto.TransactionDetailById)

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

	serviceTransactionDetail := service.TransactionDetail{}
	serviceTransactionDetail.Orm = db
	serviceTransactionDetail.Log = log
	err = serviceTransactionDetail.RemoveTransactionDetail(control, p)
	if err != nil {
		log.Errorf("Remove TransactionDetail error, %s", err.Error())
		e.Error( http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK( control.GetId(), "删除成功")
}