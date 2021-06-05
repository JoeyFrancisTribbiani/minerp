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

type ErpFinanceAsinSummary struct {
	service.Service
}

// GetErpFinanceAsinSummaryPage 获取ErpFinanceAsinSummary列表
func (e *ErpFinanceAsinSummary) GetErpFinanceAsinSummaryPage(c *dto.ErpFinanceAsinSummarySearch, p *actions.DataPermission, list *[]models.ErpFinanceAsinSummary, count *int64) error {
	var err error
	var data models.ErpFinanceAsinSummary

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetErpFinanceAsinSummaryPage error:%s", err)
		return err
	}
	return nil
}

// GetErpFinanceAsinSummary 获取ErpFinanceAsinSummary对象
func (e *ErpFinanceAsinSummary) GetErpFinanceAsinSummary(d *dto.ErpFinanceAsinSummaryById, p *actions.DataPermission, model *models.ErpFinanceAsinSummary) error {
	var data models.ErpFinanceAsinSummary

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetErpFinanceAsinSummary error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertErpFinanceAsinSummary 创建ErpFinanceAsinSummary对象
func (e *ErpFinanceAsinSummary) InsertErpFinanceAsinSummary(c *models.ErpFinanceAsinSummary) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertErpFinanceAsinSummary error:%s", err)
		return err
	}
	return nil
}

// UpdateErpFinanceAsinSummary 修改ErpFinanceAsinSummary对象
func (e *ErpFinanceAsinSummary) UpdateErpFinanceAsinSummary(c *models.ErpFinanceAsinSummary, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateErpFinanceAsinSummary error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveErpFinanceAsinSummary 删除ErpFinanceAsinSummary
func (e *ErpFinanceAsinSummary) RemoveErpFinanceAsinSummary(d *dto.ErpFinanceAsinSummaryById, p *actions.DataPermission) error {
	var data models.ErpFinanceAsinSummary

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveErpFinanceAsinSummary error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}