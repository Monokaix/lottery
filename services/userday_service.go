package services

import (
	"myproject/lottery/models"
	"myproject/lottery/dao"
)

type UserdayService interface {
	GetAll() []models.LtUserday
	CountAll() int64
	Get(id int) *models.LtUserday
	//Delete(id int) error
	Update(data *models.LtUserday, columns []string) error
	Create(data *models.LtUserday) error
}

type userdayService struct {
	dao *dao.UserdayDao
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao:dao.NewUserdayDao(nil),
	}
}
func (s *userdayService) GetAll() []models.LtUserday {
	return s.dao.GetAll()
}
func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}
//func (s *userdayService) Delete(id int) error {
//	return s.dao.Delete(id)
//}
func (s *userdayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}
