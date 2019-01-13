package services

import (
	"myproject/lottery/models"
	"myproject/lottery/dao"
)

type BlackipService interface {
	GetAll() []models.LtBlackip
	CountAll() int64
	Get(id int) *models.LtBlackip
	Delete(id int) error
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao:dao.NewBlackipDao(nil),
	}
}
func (b *blackipService) GetAll() []models.LtBlackip {
	return b.dao.GetAll()
}
func (b *blackipService) CountAll() int64 {
	return b.dao.CountAll()
}
func (b *blackipService) Get(id int) *models.LtBlackip {
	return b.dao.Get(id)
}
func (b *blackipService) Delete(id int) error {
	return b.dao.Delete(id)
}
func (b *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return b.dao.Update(data, columns)
}
func (b *blackipService) Create(data *models.LtBlackip) error {
	return b.dao.Create(data)
}
