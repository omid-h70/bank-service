package action

import (
	"context"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/helper"
	"github.com/pkg/errors"
)

var (
	errInvalidFromCardNum = errors.New("invalid from Card Number")
	errInvalidToCardNum   = errors.New("invalid to Card Number")
	errInvalidFromCardLen = errors.New("invalid from card len")
	errInvalidToCardLen   = errors.New("invalid To card len")
)

type CreateTransferAction struct {
	repo domain.AccountRepositoryDB
}

func NewCreateTransferAction(repo domain.AccountRepositoryDB) CreateTransferAction {
	return CreateTransferAction{
		repo,
		//		uc:        uc,
		//		log:       log,
		//		validator: v,
		//		logKey:    "create_transfer",
		//		logMsg:    "creating a new transfer",
	}
}

func (t CreateTransferAction) DoCardTransfer(ctx context.Context, input domain.CardTransferInput) (domain.CardTransferOutput, error) {
	out, err := t.repo.MakeTransferFromCardToCard(ctx, input)
	if err != nil {
		fmt.Println("Do Notify")
	}
	return out, err
}

func (t CreateTransferAction) isTransferValid(input domain.CardTransferInput) error {
	if len(input.CardFrom.CardNum) != 16 {
		return errInvalidFromCardLen
	}

	if len(input.CardTo.CardNum) != 16 {
		return errInvalidToCardLen
	}

	if !helper.CheckCardNumber(input.CardFrom.CardNum) {
		return errInvalidFromCardNum
	}

	if !helper.CheckCardNumber(input.CardTo.CardNum) {
		return errInvalidToCardNum
	}

	return nil
}

//
//func (t CreateTransferAction) handleErr(w http.ResponseWriter, err error) {
//	switch err {
//	case domain.ErrInsufficientBalance:
//		logging.NewError(
//			t.log,
//			err,
//			t.logKey,
//			http.StatusUnprocessableEntity,
//		).Log(t.logMsg)
//
//		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
//		return
//	case domain.ErrAccountOriginNotFound:
//		logging.NewError(
//			t.log,
//			err,
//			t.logKey,
//			http.StatusUnprocessableEntity,
//		).Log(t.logMsg)
//
//		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
//		return
//	case domain.ErrAccountDestinationNotFound:
//		logging.NewError(
//			t.log,
//			err,
//			t.logKey,
//			http.StatusUnprocessableEntity,
//		).Log(t.logMsg)
//
//		response.NewError(err, http.StatusUnprocessableEntity).Send(w)
//		return
//	default:
//		logging.NewError(
//			t.log,
//			err,
//			t.logKey,
//			http.StatusInternalServerError,
//		).Log(t.logMsg)
//
//		response.NewError(err, http.StatusInternalServerError).Send(w)
//		return
//	}
//}
//
//func (t CreateTransferAction) validateInput(input usecase.CreateTransferInput) []string {
//	var (
//		msgs              []string
//		errAccountsEquals = errors.New("account origin equals destination account")
//		accountIsEquals   = input.AccountOriginID == input.AccountDestinationID
//		accountsIsEmpty   = input.AccountOriginID == "" && input.AccountDestinationID == ""
//	)
//
//	if !accountsIsEmpty && accountIsEquals {
//		msgs = append(msgs, errAccountsEquals.Error())
//	}
//
//	err := t.validator.Validate(input)
//	if err != nil {
//		for _, msg := range t.validator.Messages() {
//			msgs = append(msgs, msg)
//		}
//	}
//
//	return msgs
//}
