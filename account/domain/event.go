package domain

// AccountClosedDomainEventImplement account domain event struct
type AccountClosedDomainEventImplement struct {
	accountID string
}

// AccountClosedDomainEvent account domain event interface
type AccountClosedDomainEvent interface {
	AccountID() string
}

// AccountID get accountID
func (e *AccountClosedDomainEventImplement) AccountID() string {
	return e.accountID
}

// AccountOpenedDomainEventImplement account domain event struct
type AccountOpenedDomainEventImplement struct {
	accountID string
}

// AccountOpenedDomainEvent account domain event interface
type AccountOpenedDomainEvent interface {
	AccountID() string
}

// AccountID get accountID
func (e *AccountOpenedDomainEventImplement) AccountID() string {
	return e.accountID
}

// AccountPasswordUpdatedDomainEventImplement account domain event struct
type AccountPasswordUpdatedDomainEventImplement struct {
	accountID string
}

// AccountPasswordUpdatedDomainEvent account domain event interface
type AccountPasswordUpdatedDomainEvent interface {
	AccountID() string
}

// AccountID get accountID
func (e *AccountPasswordUpdatedDomainEventImplement) AccountID() string {
	return e.accountID
}

// WithdrawnDomainEventImplement account domain event struct
type WithdrawnDomainEventImplement struct {
	accountID string
}

// WithdrawnDomainEvent account domain event interface
type WithdrawnDomainEvent interface {
	AccountID() string
}

// AccountID get accountID
func (e *WithdrawnDomainEventImplement) AccountID() string {
	return e.accountID
}

// DepositedDomainEventImplement account domain event struct
type DepositedDomainEventImplement struct {
	accountID string
}

// DepositedDomainEvent account domain event interface
type DepositedDomainEvent interface {
	AccountID() string
}

// AccountID get accountID
func (e *DepositedDomainEventImplement) AccountID() string {
	return e.accountID
}
