package dao

import (
	"github.com/go-xorm/xorm"
	"myproject/lottery/models"
	"log"
)

type  CodeDao struct {
	engine *xorm.Engine
}


func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine:engine,
	}
}
func (d *CodeDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err != nil{
		return data
	}else{
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift,0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil{
		log.Println("code_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil{
		return 0
	}else {
		return num
	}
}
//软删除
func (d *CodeDao) Delete(id int) error {
	data := &models.LtGift{Id:id,Sysstatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err

}

func (d *CodeDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}