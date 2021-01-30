package application

import (
	"reflect"
	"sync"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/domain"
)

// NewCommandBus create command bus
func NewCommandBus(accountRepository domain.AccountRepository, eventBus EventBus) CommandBus {
	RegisterCommandHandler(OpenAccountCommand{}, NewOpenAccountCommandHandler(accountRepository, eventBus))
	RegisterCommandHandler(UpdatePasswordCommand{}, NewUpdatePasswordCommandHandler(accountRepository, eventBus))
	RegisterCommandHandler(CloseAccountCommand{}, NewCloseAccountCommandHandler(accountRepository, eventBus))
	RegisterCommandHandler(WithdrawCommand{}, NewWithdrawCommandHandler(accountRepository, eventBus))
	RegisterCommandHandler(DepositCommand{}, NewDepositCommandHandler(accountRepository, eventBus))
	RegisterCommandHandler(RemittanceCommand{}, NewRemittanceCommandHandler(accountRepository, eventBus))
	return &CommandBusImplement{}
}

// RegisterCommandHandler register command handler by command
func RegisterCommandHandler(command interface{}, handler interface{}) {
	commandHandlerRWMutex.Lock()
	defer commandHandlerRWMutex.Unlock()
	commandHandlers[reflect.ValueOf(command).Type()] = handler
}

// CommandBus command bus interface
type CommandBus interface {
	Excute(command interface{})
}

// CommandBusImplement command bus struct
type CommandBusImplement struct{}

// Excute excute given command
func (c *CommandBusImplement) Excute(command interface{}) {
	var handler CommandHandler
	GetCommandHandler(command, handler)
	handler.handle(command)
}

// CommandHandler command handler interface
type CommandHandler interface {
	handle(command interface{})
}

// GetCommandHandler set command handler to receiver
func GetCommandHandler(command interface{}, receiver interface{}) {
	reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(commandHandlers[reflect.TypeOf(command)]))
}

var commandHandlers = map[reflect.Type]interface{}{}
var commandHandlerRWMutex = sync.RWMutex{}

// NewOpenAccountCommandHandler create new OpenAccountCommandHandler
func NewOpenAccountCommandHandler(repository domain.AccountRepository, eventBus EventBus) OpenAccountCommandHandler {
	return &OpenAccountCommandHandlerImplement{AccountRepository: repository, EventBus: eventBus}
}

// OpenAccountCommandHandler OpenAccountCommandHandler interface
type OpenAccountCommandHandler interface {
	handle(command OpenAccountCommand)
}

// OpenAccountCommandHandlerImplement OpenAccountCommandHandler struct
type OpenAccountCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
}

func (h *OpenAccountCommandHandlerImplement) handle(command OpenAccountCommand) {
	id, err := h.AccountRepository.FindNewID()
	if err != nil {
		panic(err)
	}

	account := domain.NewAccount(domain.AccountOptions{
		ID:       id,
		Name:     command.Name,
		Password: command.Password,
	})

	if err := h.AccountRepository.Save(account); err != nil {
		panic(err)
	}

	for _, event := range account.Events() {
		h.EventBus.Excute(event)
	}
}

// OpenAccountCommand command for open account
type OpenAccountCommand struct {
	Name     string
	Password string
}

// NewUpdatePasswordCommandHandler create new UpdatePasswordCommandHandler
func NewUpdatePasswordCommandHandler(repository domain.AccountRepository, eventBus EventBus) UpdatePasswordCommandHandler {
	return &UpdatePasswordCommandHandlerImplement{AccountRepository: repository, EventBus: eventBus}
}

// UpdatePasswordCommandHandler UpdatePasswordCommandHandler interface
type UpdatePasswordCommandHandler interface {
	handle(command UpdatePasswordCommand)
}

// UpdatePasswordCommandHandlerImplement UpdatePasswordCommandHandler struct
type UpdatePasswordCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
}

func (h *UpdatePasswordCommandHandlerImplement) handle(command UpdatePasswordCommand) {
	account, err := h.AccountRepository.FindByID(command.ID)
	if err != nil {
		panic(err)
	}

	account.UpdatePassword(command.New, command.Password)

	if err := h.AccountRepository.Save(account); err != nil {
		panic(err)
	}

	for _, event := range account.Events() {
		h.EventBus.Excute(event)
	}
}

// UpdatePasswordCommand command for update password
type UpdatePasswordCommand struct {
	ID       string
	Password string
	New      string
}

// NewCloseAccountCommandHandler create new CloseAccountCommandHandler
func NewCloseAccountCommandHandler(repository domain.AccountRepository, eventbus EventBus) CloseAccountCommandHandler {
	return &CloseAccountCommandHandlerImplement{AccountRepository: repository, EventBus: eventbus}
}

// CloseAccountCommandHandler CloseAccountCommandHandler interface
type CloseAccountCommandHandler interface {
	handle(command CloseAccountCommand)
}

// CloseAccountCommandHandlerImplement CloseAccountCommandHandler struct
type CloseAccountCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
}

func (h *CloseAccountCommandHandlerImplement) handle(command CloseAccountCommand) {
	account, err := h.AccountRepository.FindByID(command.ID)
	if err != nil {
		panic(err)
	}

	account.Close(command.Password)

	if err := h.AccountRepository.Save(account); err != nil {
		panic(err)
	}

	for _, event := range account.Events() {
		h.EventBus.Excute(event)
	}
}

// CloseAccountCommand command for close account
type CloseAccountCommand struct {
	ID       string
	Password string
}

// NewWithdrawCommandHandler create new WithdrawCommandHandler
func NewWithdrawCommandHandler(repository domain.AccountRepository, eventBus EventBus) WithdrawCommandHandler {
	return &WithdrawCommandHandlerImplement{AccountRepository: repository, EventBus: eventBus}
}

// WithdrawCommandHandler WithdrawCommandHandler interface
type WithdrawCommandHandler interface {
	handle(command WithdrawCommand)
}

// WithdrawCommandHandlerImplement WithdrawCommandHandler struct
type WithdrawCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
}

func (h *WithdrawCommandHandlerImplement) handle(command WithdrawCommand) {
	account, err := h.AccountRepository.FindByID(command.ID)
	if err != nil {
		panic(err)
	}

	account.Withdraw(command.Amount, command.Password)

	if err := h.AccountRepository.Save(account); err != nil {
		panic(err)
	}

	for _, event := range account.Events() {
		h.EventBus.Excute(event)
	}
}

// WithdrawCommand command for withdraw
type WithdrawCommand struct {
	ID       string
	Password string
	Amount   int
}

// NewDepositCommandHandler create new DepositCommandHandler
func NewDepositCommandHandler(repository domain.AccountRepository, eventBus EventBus) DepositCommandHandler {
	return &DepositCommandHandlerImplement{AccountRepository: repository, EventBus: eventBus}
}

// DepositCommandHandler DepositCommandHandler interface
type DepositCommandHandler interface {
	handle(command DepositCommand)
}

// DepositCommandHandlerImplement DepositCommandHandler struct
type DepositCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
}

func (h *DepositCommandHandlerImplement) handle(command DepositCommand) {
	account, err := h.AccountRepository.FindByID(command.ID)
	if err != nil {
		panic(err)
	}

	account.Deposit(command.Amount)

	if err := h.AccountRepository.Save(account); err != nil {
		panic(err)
	}

	for _, event := range account.Events() {
		h.EventBus.Excute(event)
	}
}

// DepositCommand command for deposit
type DepositCommand struct {
	ID     string
	Amount int
}

// NewRemittanceCommandHandler create new RemittanceCommandHandler
func NewRemittanceCommandHandler(repository domain.AccountRepository, eventBus EventBus) RemittanceCommandHandler {
	return &RemittanceCommandHandlerImplement{AccountRepository: repository, EventBus: eventBus}
}

// RemittanceCommandHandler RemittanceCommandHandler interface
type RemittanceCommandHandler interface {
	handle(command RemittanceCommand)
}

// RemittanceCommandHandlerImplement RemittanceCommandHandler struct
type RemittanceCommandHandlerImplement struct {
	EventBus
	domain.AccountRepository
	domain.AccountDomainService
}

func (h *RemittanceCommandHandlerImplement) handle(command RemittanceCommand) {
	sender, err := h.AccountRepository.FindByID(command.SenderID)
	if err != nil {
		panic(err)
	}

	receiver, err := h.AccountRepository.FindByID(command.ReceiverID)
	if err != nil {
		panic(err)
	}

	h.AccountDomainService.Remit(domain.RemittanceOptions{
		Password: command.Password,
		Sender:   sender,
		Receiver: receiver,
		Amount:   command.Amount,
	})

	if err := h.AccountRepository.Save(sender); err != nil {
		panic(err)
	}

	if err := h.AccountRepository.Save(receiver); err != nil {
		panic(err)
	}

	for _, event := range sender.Events() {
		h.EventBus.Excute(event)
	}

	for _, event := range receiver.Events() {
		h.EventBus.Excute(event)
	}
}

// RemittanceCommand command for remittance
type RemittanceCommand struct {
	SenderID   string
	ReceiverID string
	Password   string
	Amount     int
}
