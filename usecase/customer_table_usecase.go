package usecase

import (
	"errors"
	"github.com/edwardsuwirya/wmbTableMgmt/apperror"
	"github.com/edwardsuwirya/wmbTableMgmt/constant"
	"github.com/edwardsuwirya/wmbTableMgmt/dto"
	"github.com/edwardsuwirya/wmbTableMgmt/entity"
	"github.com/edwardsuwirya/wmbTableMgmt/repository"
	"log"
)

type ICustomerTableUseCase interface {
	GetTodayListCustomerTable() ([]dto.TableAvailablity, error)
	TableCheckIn(trx dto.CheckInRequest) (*entity.CustomerTableTransaction, error)
	TableCheckOut(billNo string) error
}

type CustomerTableUseCase struct {
	customerTableTrxRepo repository.ICustomerTableTrxRepository
}

func NewCustomerTableUseCase(customerTableTrxRepo repository.ICustomerTableTrxRepository) ICustomerTableUseCase {
	return &CustomerTableUseCase{customerTableTrxRepo}
}

func (t *CustomerTableUseCase) GetTodayListCustomerTable() ([]dto.TableAvailablity, error) {
	return t.customerTableTrxRepo.GetAllByBusinessDate()
}

func (t *CustomerTableUseCase) TableCheckIn(trx dto.CheckInRequest) (*entity.CustomerTableTransaction, error) {
	tbl, err := t.customerTableTrxRepo.CountByTableId(trx.TableId, constant.TableOccupied)
	if err != nil {
		return nil, errors.New("Failed to check in")
	}
	log.Println(tbl)
	if tbl == 0 {
		return t.customerTableTrxRepo.CreateOne(entity.CustomerTableTransaction{BillNo: trx.BillNo, CustomerTableID: trx.TableId})
	} else {
		return nil, apperror.TableOccupiedError
	}

}

func (t *CustomerTableUseCase) TableCheckOut(billNo string) error {
	return t.customerTableTrxRepo.Delete(billNo)
}
