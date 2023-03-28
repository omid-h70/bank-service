package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/adapter/action"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"time"
)

type (
	//	// TransferService input port
	//	TransferService interface {
	//		Execute(context.Context, CreateTransferInput) (CreateTransferOutput, error)
	//	}
	//
	//	// CreateTransferInput input data
	//	CreateTransferInput struct {
	//		AccountIdFrom string `json:"account_from_id" validate:"required,uuid4"`
	//		AccountIdTo   string `json:"account_to_id" validate:"required,uuid4"`
	//		Amount        int64  `json:"amount" validate:"gt=0,required"`
	//	}
	//
	//	// CreateTransferPresenter output port
	//	//CreateTransferPresenter interface {
	//	//	Output(domain.Transfer) CreateTransferOutput
	//	//}
	//
	//	// CreateTransferOutput output data
	//	CreateTransferOutput struct {
	//		ID            string  `json:"id"`
	//		AccountFromId string  `json:"account_from_id"`
	//		AccountIdTo   string  `json:"account_to_id"`
	//		Amount        float64 `json:"amount"`
	//		TransferTime  string  `json:"date_time"`
	//	}
	//
	TransferTransaction struct {
		//	//transferRepo domain.TransferRepository
		//	//accountRepo  domain.AccountRepository
		notifyHandler domain.PushNotificationService
		actionHandler action.CreateTransferAction
		ctxTimeout    time.Duration
	}
)

// NewTransferService creates new NewTransferService with its dependencies
func NewTransferService(
	// // transferRepo domain.TransferRepository,
	// // accountRepo domain.AccountRepository,
	a action.CreateTransferAction,
	t time.Duration,
) TransferTransaction {
	return TransferTransaction{
		//		//transferRepo: transferRepo,
		//		//accountRepo:  accountRepo,
		actionHandler: a,
		ctxTimeout:    t,
	}
}

func (t TransferTransaction) ExecuteCardTransfer(ctx context.Context, input domain.CardTransferInput) (domain.CardTransferOutput, error) {
	return t.actionHandler.DoCardTransfer(ctx, input)
}

//
//func (t createTransferInteractor) process(ctx context.Context, input CreateTransferInput) error {
//	origin, err := t.accountRepo.FindByID(ctx, domain.AccountID(input.AccountOriginID))
//	if err != nil {
//		switch err {
//		case domain.ErrAccountNotFound:
//			return domain.ErrAccountOriginNotFound
//		default:
//			return err
//		}
//	}
//
//	if err := origin.Withdraw(domain.Money(input.Amount)); err != nil {
//		return err
//	}
//
//	destination, err := t.accountRepo.FindByID(ctx, domain.AccountID(input.AccountDestinationID))
//	if err != nil {
//		switch err {
//		case domain.ErrAccountNotFound:
//			return domain.ErrAccountDestinationNotFound
//		default:
//			return err
//		}
//	}
//
//	destination.Deposit(domain.Money(input.Amount))
//
//	if err = t.accountRepo.UpdateBalance(ctx, origin.ID(), origin.Balance()); err != nil {
//		return err
//	}
//
//	if err = t.accountRepo.UpdateBalance(ctx, destination.ID(), destination.Balance()); err != nil {
//		return err
//	}
//
//	return nil
//}
