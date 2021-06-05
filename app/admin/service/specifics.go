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

type ErpProductSpecifics struct {
	service.Service
}

// GetErpProductSpecificsPage 获取ErpProductSpecifics列表
func (e *ErpProductSpecifics) GetErpProductSpecificsPage(c *dto.ErpProductSpecificsSearch, p *actions.DataPermission, list *[]models.ErpProductSpecifics, count *int64) error {
	var err error
	var data models.ErpProductSpecifics

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetErpProductSpecificsPage error:%s", err)
		return err
	}
	return nil
}

// GetErpProductSpecifics 获取ErpProductSpecifics对象
func (e *ErpProductSpecifics) GetErpProductSpecifics(d *dto.ErpProductSpecificsById, p *actions.DataPermission, model *models.ErpProductSpecifics) error {
	var data models.ErpProductSpecifics

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetErpProductSpecifics error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertErpProductSpecifics 创建ErpProductSpecifics对象
func (e *ErpProductSpecifics) InsertErpProductSpecifics(c *models.ErpProductSpecifics) error {
	err := e.Orm.Create(c).Error
	if err != nil {
		e.Log.Errorf("Service InsertErpProductSpecifics error:%s", err)
		return err
	}
	return nil
}

// UpdateErpProductSpecifics 修改ErpProductSpecifics对象
func (e *ErpProductSpecifics) UpdateErpProductSpecifics(c *models.ErpProductSpecifics, p *actions.DataPermission) error {
	db := e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if err := db.Error; err != nil {
        e.Log.Errorf("Service UpdateErpProductSpecifics error:%s", err)
        return err
    }
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// RemoveErpProductSpecifics 删除ErpProductSpecifics
func (e *ErpProductSpecifics) RemoveErpProductSpecifics(d *dto.ErpProductSpecificsById, p *actions.DataPermission) error {
	var data models.ErpProductSpecifics

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveErpProductSpecifics error:%s", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}