package application

import (
	"reflect"
	"sync"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/domain"
)

// NewEventBus create new EventBus
func NewEventBus() EventBus {
	return &EventBusImplement{}
}

// EventBus EventBus interface
type EventBus interface {
	Excute(event interface{})
}

// EventBusImplement EventBus struct
type EventBusImplement struct{}

// Excute excute given event
func (b *EventBusImplement) Excute(event interface{}) {
	var handler EventHandler
	GetEventHandler(event, handler)
	handler.handle(event)
}

// GetEventHandler set event handler to receiver
func GetEventHandler(event interface{}, receiver interface{}) {
	reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(eventHandlers[reflect.TypeOf(event)]))
}

// RegisterEventHandler register event handler by event
func RegisterEventHandler(event interface{}, handler interface{}) {
	eventHandlerRWMutex.Lock()
	defer eventHandlerRWMutex.Unlock()
	eventHandlers[reflect.ValueOf(event).Type()] = handler
}

var eventHandlers = map[reflect.Type]interface{}{}
var eventHandlerRWMutex = sync.RWMutex{}

// EventHandler event handler interface
type EventHandler interface {
	handle(event interface{})
}

// NewAccountOpenedEventHandler create new AccountOpenedEventHandler
func NewAccountOpenedEventHandler(eventPublisher EventPublisher) AccountOpenedEventHandler {
	return &AccountOpenedEventHandlerImplement{EventPublisher: eventPublisher}
}

// AccountOpenedEventHandler AccountOpenedEventHandler interface
type AccountOpenedEventHandler interface {
	handle(event domain.AccountOpenedDomainEvent)
}

// AccountOpenedEventHandlerImplement AccountOpenedEventHandler struct
type AccountOpenedEventHandlerImplement struct {
	EventPublisher
}

func (h *AccountOpenedEventHandlerImplement) handle(event domain.AccountOpenedDomainEvent) {
	h.EventPublisher.Publish(event.AccountID())
}

// NewAccountClosedEventHandler create new AccountClosedEventHandler
func NewAccountClosedEventHandler(eventPublisher EventPublisher) AccountClosedEventHandler {
	return &AccountClosedEventHandlerImplement{EventPublisher: eventPublisher}
}

// AccountClosedEventHandler AccountEventHandler interface
type AccountClosedEventHandler interface {
	handle(event domain.AccountClosedDomainEvent)
}

// AccountClosedEventHandlerImplement AccountEventHandler struct
type AccountClosedEventHandlerImplement struct {
	EventPublisher
}

func (h *AccountClosedEventHandlerImplement) handle(event domain.AccountClosedDomainEvent) {
	h.EventPublisher.Publish(event.AccountID())
}

// NewAccountPasswordUpdatedEventHandler create new AccountPasswordUpdatedEventHandler
func NewAccountPasswordUpdatedEventHandler(eventPublisher EventPublisher) AccountPasswordUpdatedEventHandler {
	return &AccountPasswordUpdatedEventHandlerImplement{EventPublisher: eventPublisher}
}

// AccountPasswordUpdatedEventHandler AccountPasswordUpdatedEventHandler interface
type AccountPasswordUpdatedEventHandler interface {
	handle(event domain.AccountPasswordUpdatedDomainEvent)
}

// AccountPasswordUpdatedEventHandlerImplement AccountPasswordUpdatedEventHandler struct
type AccountPasswordUpdatedEventHandlerImplement struct {
	EventPublisher
}

func (h *AccountPasswordUpdatedEventHandlerImplement) handle(event domain.AccountPasswordUpdatedDomainEvent) {
	h.EventPublisher.Publish(event.AccountID())
}

// NewWithdrawnEventHandler create new WithdrawnEventHandler
func NewWithdrawnEventHandler(eventPublisher EventPublisher) WithdrawnEventHandler {
	return &WithdrawnEventHandlerImplement{EventPublisher: eventPublisher}
}

// WithdrawnEventHandler WithdrawnEventHandler interface
type WithdrawnEventHandler interface {
	handle(event domain.WithdrawnDomainEvent)
}

// WithdrawnEventHandlerImplement WithdrawnEventHandler struct
type WithdrawnEventHandlerImplement struct {
	EventPublisher
}

func (h *WithdrawnEventHandlerImplement) handle(event domain.WithdrawnDomainEvent) {
	h.EventPublisher.Publish(event.AccountID())
}

// NewDepositedEventHandler create new DepositedEventHandler
func NewDepositedEventHandler(eventPublisher EventPublisher) DepositedEventHandler {
	return &DepositedEventHandlerImplement{EventPublisher: eventPublisher}
}

// DepositedEventHandler DepositedEventHandler interface
type DepositedEventHandler interface {
	handle(event domain.DepositedDomainEvent)
}

// DepositedEventHandlerImplement DepositedEventHandler struct
type DepositedEventHandlerImplement struct {
	EventPublisher
}

func (h *DepositedEventHandlerImplement) handle(event domain.DepositedDomainEvent) {
	h.EventPublisher.Publish(event)
}

// EventPublisher integration event publisher
type EventPublisher interface {
	Publish(event interface{})
}
