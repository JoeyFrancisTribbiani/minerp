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

type ErpFinanceSummary struct {
	service.Service
}

// GetErpFinanceSummaryPage 获取ErpFinanceSummary列表
func (e *ErpFinanceSummary) GetErpFinanceSummaryPage(c *dto.ErpFinanceSummarySearch, p *actions.DataPermission, list *[]models.ErpFinanceSummary, count *int64) error {
	var err error
	var data models.ErpFinanceSummary

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetErpFinanceSummaryPage error:%s", err)
		return err
	}
	return nil
}

// GetErpFinanceSummary 获取ErpFinanceSummary对象
func (e *ErpFinanceSummary) GetErpFinanceSummary(d *dto.ErpFinanceSummaryById, p *actions.DataPermission, model *models.ErpFinanceSummary) error {
	var data models.ErpFinanceSummary

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetErpFinanceSummary error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertErpFinanceSummary 创建ErpFinanceSummary对象
func (e *ErpFinanceSummary) InsertErpFinanceSummary(c *models.ErpFinanceSummary) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertErpFinanceSummary error:%s", err)
		return err
	}
	return nil
}

// UpdateErpFinanceSummary 修改ErpFinanceSummary对象
func (e *ErpFinanceSummary) UpdateErpFinanceSummary(c *models.ErpFinanceSummary, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateErpFinanceSummary error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveErpFinanceSummary 删除ErpFinanceSummary
func (e *ErpFinanceSummary) RemoveErpFinanceSummary(d *dto.ErpFinanceSummaryById, p *actions.DataPermission) error {
	var data models.ErpFinanceSummary

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveErpFinanceSummary error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}