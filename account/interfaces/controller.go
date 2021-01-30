package interfaces

import "github.com/kyhsa93/gin-rest-cqrs-example/account/application"

// Controller controller interface
type Controller interface {
	OpenAccount(dto OpenAccountDTO)
	UpdateAccountPassword(dto UpdateAccountPasswordDTO)
	CloseAccount(dto CloseAccountDTO)
	Withdraw(dto WithdrawDTO)
	Deposit(dto DepositDTO)
	Remittance(dto RemittanceDTO)
}

// ControllerImplement controller struct
type ControllerImplement struct {
	application.CommandBus
}

// OpenAccount handle open account request
func (c *ControllerImplement) OpenAccount(dto OpenAccountDTO) {
	c.CommandBus.Excute(application.OpenAccountCommand{
		Name:     dto.Name,
		Password: dto.Password,
	})
}

// OpenAccountDTO dto for open account
type OpenAccountDTO struct {
	Name     string
	Password string
}

// UpdateAccountPassword handle update account password request
func (c *ControllerImplement) UpdateAccountPassword(dto UpdateAccountPasswordDTO) {
	c.CommandBus.Excute(application.UpdatePasswordCommand{
		ID:       dto.ID,
		Password: dto.Password,
		New:      dto.New,
	})
}

// UpdateAccountPasswordDTO dto for update account password
type UpdateAccountPasswordDTO struct {
	ID       string
	Password string
	New      string
}

// CloseAccount handle close account request
func (c *ControllerImplement) CloseAccount(dto CloseAccountDTO) {
	c.CommandBus.Excute(application.CloseAccountCommand{
		ID:       dto.ID,
		Password: dto.Password,
	})
}

// CloseAccountDTO dto for close account
type CloseAccountDTO struct {
	ID       string
	Password string
}

// Withdraw handle withdraw from account request
func (c *ControllerImplement) Withdraw(dto WithdrawDTO) {
	c.CommandBus.Excute(application.WithdrawCommand{
		ID:       dto.ID,
		Password: dto.Password,
		Amount:   dto.Amount,
	})
}

// WithdrawDTO dto for withdraw
type WithdrawDTO struct {
	ID       string
	Password string
	Amount   int
}

// Deposit handle deposit to account request
func (c *ControllerImplement) Deposit(dto DepositDTO) {
	c.CommandBus.Excute(application.DepositCommand{
		ID:     dto.ID,
		Amount: dto.Amount,
	})
}

// DepositDTO dto for deposit
type DepositDTO struct {
	ID       string
	Password string
	Amount   int
}

// Remittance handle remittance request
func (c *ControllerImplement) Remittance(dto RemittanceDTO) {
	c.CommandBus.Excute(application.RemittanceCommand{
		SenderID:   dto.SenderID,
		Password:   dto.Password,
		ReceiverID: dto.ReceiverID,
		Amount:     dto.Amount,
	})
}

// RemittanceDTO dto for remittance
type RemittanceDTO struct {
	SenderID   string
	Password   string
	ReceiverID string
	Amount     int
}
