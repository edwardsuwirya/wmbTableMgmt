package delivery

import (
	"github.com/edwardsuwirya/wmbTableMgmt/apperror"
	"github.com/edwardsuwirya/wmbTableMgmt/delivery/appresponse"
	"github.com/edwardsuwirya/wmbTableMgmt/dto"
	"github.com/edwardsuwirya/wmbTableMgmt/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerTableApi struct {
	usecase     usecase.ICustomerTableUseCase
	publicRoute *gin.RouterGroup
}

func NewCustomerTableApi(publicRoute *gin.RouterGroup, usecase usecase.ICustomerTableUseCase) *CustomerTableApi {
	customerTableApi := CustomerTableApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	customerTableApi.initRouter()
	return &customerTableApi
}
func (api *CustomerTableApi) initRouter() {
	userRoute := api.publicRoute.Group("/table")
	userRoute.GET("", api.getTableList)
	userRoute.POST("/checkin", api.tableCheckIn)
	userRoute.PUT("/checkout", api.tableCheckOut)
}
func (api *CustomerTableApi) tableCheckOut(c *gin.Context) {
	billNo := c.Query("billNo")
	err := api.usecase.TableCheckOut(billNo)
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Check Out"))
		return
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Success Check Out", nil))
}
func (api *CustomerTableApi) tableCheckIn(c *gin.Context) {
	var checkInRequest dto.CheckInRequest
	if err := c.ShouldBindJSON(&checkInRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := api.usecase.TableCheckIn(checkInRequest)
	if err != nil {
		if err == apperror.TableOccupiedError {
			appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("FAILED", err.Error(), nil))
			return
		}
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Check In"))
		return
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Success Check In", nil))
}

func (api *CustomerTableApi) getTableList(c *gin.Context) {
	tableList, err := api.usecase.GetTodayListCustomerTable()
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewBadRequestError(err, "Failed Get Table List"))
		return
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "List Table", tableList))
}
