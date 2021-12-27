package repository

import (
	"errors"
	"github.com/edwardsuwirya/wmbTableMgmt/constant"
	"github.com/edwardsuwirya/wmbTableMgmt/dto"
	"github.com/edwardsuwirya/wmbTableMgmt/entity"
	"github.com/edwardsuwirya/wmbTableMgmt/util"
	"gorm.io/gorm"
	"log"
)

type ICustomerTableTrxRepository interface {
	CreateOne(table entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error)
	GetAllByBusinessDate() ([]dto.TableAvailablity, error)
	CountByTableId(tableId string, status constant.TableStatus) (int64, error)
	Delete(billNo string) error
}

type CustomerTableTrxRepository struct {
	db *gorm.DB
}

func NewCustomerTableRepository(resource *gorm.DB) ICustomerTableTrxRepository {
	customerTableTrx := &CustomerTableTrxRepository{db: resource}
	return customerTableTrx
}

func (s *CustomerTableTrxRepository) GetAllByBusinessDate() ([]dto.TableAvailablity, error) {
	var tableListResult []dto.TableAvailablity
	sd, ed := util.GetTodayWithTime()
	err := s.db.Model(&entity.CustomerTableTransaction{}).
		Select("customer_table.id as table_id,count(customer_table_transaction.created_at) as is_occupied").
		Group("customer_table.id").
		Joins(`
			right join customer_table on customer_table_transaction.customer_table_id = customer_table.id and
			customer_table_transaction.created_at between ? and ? and customer_table_transaction.deleted_at is null
		`, sd, ed).
		Scan(&tableListResult).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return tableListResult, nil
}

func (s *CustomerTableTrxRepository) Delete(billNo string) error {
	err := s.db.Where("bill_no = ?", billNo).Delete(&entity.CustomerTableTransaction{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *CustomerTableTrxRepository) CreateOne(table entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error) {
	err := s.db.Create(&table).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &table, nil
}
func (s *CustomerTableTrxRepository) CountByTableId(tableId string, status constant.TableStatus) (int64, error) {
	var count int64
	switch status {
	case constant.TableAllStatus:
		err := s.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ?", tableId).Count(&count).Error
		if err != nil {
			log.Println(err)
			return -1, err
		}
	case constant.TableVacant:
		err := s.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ? and deleted_at is not null", tableId).Count(&count).Error
		if err != nil {
			log.Println(err)
			return -1, err
		}
	case constant.TableOccupied:
		err := s.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ? and deleted_at is  null", tableId).Count(&count).Error
		if err != nil {
			log.Println(err)
			return -1, err
		}
	default:
		return -1, errors.New("Unknown status")
	}
	return count, nil
}
