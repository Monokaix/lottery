package services

import (
	"myproject/lottery/models"
	"myproject/lottery/dao"
)

type ResultService interface {
	GetAll() []models.LtResult
	CountAll() int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao:dao.NewResultDao(nil),
	}
}
func (r *resultService) GetAll() []models.LtResult {
	return r.dao.GetAll()
}
func (r *resultService) CountAll() int64 {
	return r.dao.CountAll()
}
func (r *resultService) Get(id int) *models.LtResult {
	return r.dao.Get(id)
}
func (r *resultService) Delete(id int) error {
	return r.dao.Delete(id)
}
func (r *resultService) Update(data *models.LtResult, columns []string) error {
	return r.dao.Update(data, columns)
}
func (r *resultService) Create(data *models.LtResult) error {
	return r.dao.Create(data)
}