package services

import (
	"myproject/lottery/models"
	"myproject/lottery/dao"
)

type CodeService interface {
	GetAll() []models.LtCode
	CountAll() int64
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(data *models.LtCode, columns []string) error
	Create(data *models.LtCode) error
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao:dao.NewCodeDao(nil),
	}
}
func (c *codeService) GetAll() []models.LtCode {
	return c.dao.GetAll()
}
func (c *codeService) CountAll() int64 {
	return c.dao.CountAll()
}
func (c *codeService) Get(id int) *models.LtCode {
	return c.dao.Get(id)
}
func (c *codeService) Delete(id int) error {
	return c.dao.Delete(id)
}
func (c *codeService) Update(data *models.LtCode, columns []string) error {
	return c.dao.Update(data, columns)
}
func (c *codeService) Create(data *models.LtCode) error {
	return c.dao.Create(data)
}
