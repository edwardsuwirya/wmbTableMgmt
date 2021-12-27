package manager

import "github.com/edwardsuwirya/wmbTableMgmt/usecase"

type UseCaseManager interface {
	CustomerTableUseCase() usecase.ICustomerTableUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) CustomerTableUseCase() usecase.ICustomerTableUseCase {
	return usecase.NewCustomerTableUseCase(uc.repo.CustomerTableTransactionRepo())
}
func NewUseCaseManger(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
