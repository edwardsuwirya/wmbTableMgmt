package manager

import "github.com/edwardsuwirya/wmbTableMgmt/repository"

type RepoManager interface {
	CustomerTableTransactionRepo() repository.ICustomerTableTrxRepository
}

type repoManager struct {
	infra Infra
}

func (rm *repoManager) CustomerTableTransactionRepo() repository.ICustomerTableTrxRepository {
	return repository.NewCustomerTableRepository(rm.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra}
}
