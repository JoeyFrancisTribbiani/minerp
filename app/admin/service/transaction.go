package service

import (
	"errors"

	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/app/admin/service/dto"
	"go-admin/common/service"
)

type ErpFinanceTransaction struct {
	service.Service
}

// GetErpFinanceTransactionPage 获取ErpFinanceTransaction列表
func (e *ErpFinanceTransaction) GetErpFinanceTransactionPage(c *dto.ErpFinanceTransactionSearch, p *actions.DataPermission, list *[]models.ErpFinanceTransaction, count *int64) error {
	var err error
	var data models.ErpFinanceTransaction

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetErpFinanceTransactionPage error:%s", err)
		return err
	}
	return nil
}

// GetErpFinanceTransaction 获取ErpFinanceTransaction对象
func (e *ErpFinanceTransaction) GetErpFinanceTransaction(d *dto.ErpFinanceTransactionById, p *actions.DataPermission, model *models.ErpFinanceTransaction) error {
	var data models.ErpFinanceTransaction

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetErpFinanceTransaction error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertErpFinanceTransaction 创建ErpFinanceTransaction对象
func (e *ErpFinanceTransaction) InsertErpFinanceTransaction(c *models.ErpFinanceTransaction) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertErpFinanceTransaction error:%s", err)
		return err
	}
	return nil
}

// UpdateErpFinanceTransaction 修改ErpFinanceTransaction对象
func (e *ErpFinanceTransaction) UpdateErpFinanceTransaction(c *models.ErpFinanceTransaction, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateErpFinanceTransaction error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveErpFinanceTransaction 删除ErpFinanceTransaction
func (e *ErpFinanceTransaction) RemoveErpFinanceTransaction(d *dto.ErpFinanceTransactionById, p *actions.DataPermission) error {
	var data models.ErpFinanceTransaction

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveErpFinanceTransaction error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}