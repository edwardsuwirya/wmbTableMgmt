package delivery

import (
	"github.com/edwardsuwirya/wmbTableMgmt/manager"
	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func NewServer(engine *gin.Engine, useCaseManager manager.UseCaseManager) *Routes {
	newServer := new(Routes)

	publicRoute := engine.Group("/api")
	NewCustomerTableApi(publicRoute, useCaseManager.CustomerTableUseCase())
	return newServer
}
