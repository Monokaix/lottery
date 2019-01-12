package dao

import (
	"github.com/go-xorm/xorm"
	"myproject/lottery/models"
	"log"
)

type  BlackipDao struct {
	engine *xorm.Engine
}


func NewBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine:engine,
	}
}
func (d *BlackipDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err != nil{
		return data
	}else{
		data.Id = 0
		return data
	}
}

func (d *BlackipDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift,0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil{
		log.Println("gift_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *BlackipDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil{
		return 0
	}else {
		return num
	}
}
//软删除
func (d *BlackipDao) Delete(id int) error {
	data := &models.LtGift{Id:id,Sysstatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err

}

func (d *BlackipDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackipDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackipDao) GetByIp(ip string) *models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.Where("ip?=",ip).Decr("id").
		Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1{
		return nil
	}else {
		return &datalist[0]
	}
}