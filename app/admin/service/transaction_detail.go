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

type TransactionDetail struct {
	service.Service
}

// GetTransactionDetailPage 获取TransactionDetail列表
func (e *TransactionDetail) GetTransactionDetailPage(c *dto.TransactionDetailSearch, p *actions.DataPermission, list *[]models.TransactionDetail, count *int64) error {
	var err error
	var data models.TransactionDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetTransactionDetailPage error:%s", err)
		return err
	}
	return nil
}

// GetTransactionDetail 获取TransactionDetail对象
func (e *TransactionDetail) GetTransactionDetail(d *dto.TransactionDetailById, p *actions.DataPermission, model *models.TransactionDetail) error {
	var data models.TransactionDetail

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTransactionDetail error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertTransactionDetail 创建TransactionDetail对象
func (e *TransactionDetail) InsertTransactionDetail(c *models.TransactionDetail) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertTransactionDetail error:%s", err)
		return err
	}
	return nil
}

// UpdateTransactionDetail 修改TransactionDetail对象
func (e *TransactionDetail) UpdateTransactionDetail(c *models.TransactionDetail, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateTransactionDetail error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveTransactionDetail 删除TransactionDetail
func (e *TransactionDetail) RemoveTransactionDetail(d *dto.TransactionDetailById, p *actions.DataPermission) error {
	var data models.TransactionDetail

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTransactionDetail error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}