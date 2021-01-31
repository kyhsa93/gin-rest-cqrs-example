package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/interfaces"
)

// NewRouter create new http router
func NewRouter(controller interfaces.Controller) Router {
	router := &RouterImplement{Controller: controller}
	engine := gin.Default()
	engine.POST("/accounts", router.handleCreateAccountRequest)
	engine.POST("/accounts/:accountID/withdraw", router.handleWithdrawRequest)
	engine.POST("/accounts/:accountID/deposit", router.handleDepositRequest)
	engine.POST("/accounts/:accountID/remittance", router.handleRemittanceRequest)
	engine.PUT("/accounts/:accountID/password", router.handleUpdateAccountPasswordRequest)
	engine.DELETE("/accounts/:accountID", router.handleCloseAccountRequest)
	return router
}

func (r *RouterImplement) handleCreateAccountRequest(context *gin.Context) {
	var body interfaces.OpenAccountDTO
	context.ShouldBindJSON(body)
	r.Controller.OpenAccount(body)
}

func (r *RouterImplement) handleUpdateAccountPasswordRequest(context *gin.Context) {
	var body updateAccountPasswordRequestBody
	context.ShouldBindJSON(body)
	r.Controller.UpdateAccountPassword(interfaces.UpdateAccountPasswordDTO{
		ID:       context.Param("accountID"),
		Password: body.Password,
		New:      body.New,
	})
}

type updateAccountPasswordRequestBody struct {
	Password string
	New      string
}

func (r *RouterImplement) handleCloseAccountRequest(context *gin.Context) {
	r.Controller.CloseAccount(interfaces.CloseAccountDTO{
		ID:       context.Param("accountID"),
		Password: context.Query("password"),
	})
}

func (r *RouterImplement) handleWithdrawRequest(context *gin.Context) {
	var body withdrawRequestBody
	context.ShouldBindJSON(body)
	r.Controller.Withdraw(interfaces.WithdrawDTO{
		ID:       context.Param("accountID"),
		Password: body.Password,
		Amount:   body.Amount,
	})
}

type withdrawRequestBody struct {
	Password string
	Amount   int
}

func (r *RouterImplement) handleDepositRequest(context *gin.Context) {
	var body depositRequestBody
	context.ShouldBindJSON(body)
	r.Controller.Deposit(interfaces.DepositDTO{
		ID:       context.Param("accountID"),
		Password: body.Password,
		Amount:   body.Amount,
	})
}

type depositRequestBody struct {
	Password string
	Amount   int
}

func (r *RouterImplement) handleRemittanceRequest(context *gin.Context) {
	var body remittanceRequestBody
	context.ShouldBindJSON(body)
	r.Controller.Remittance(interfaces.RemittanceDTO{
		SenderID:   context.Param("accountID"),
		Password:   body.Password,
		ReceiverID: body.receiverID,
		Amount:     body.Amount,
	})
}

type remittanceRequestBody struct {
	Password   string
	receiverID string
	Amount     int
}

// Router http router
type Router interface{}

// RouterImplement http route implement
type RouterImplement struct {
	interfaces.Controller
}
