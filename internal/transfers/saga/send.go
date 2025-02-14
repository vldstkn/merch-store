package saga

import (
	"context"
	"merch_store/internal/interfaces"
	"merch_store/pkg/pb"
)

type SendSaga struct {
	AccountClient   pb.AccountClient
	TransferService interfaces.TransfersService
}

type SendSagaDeps struct {
	AccountClient   pb.AccountClient
	TransferService interfaces.TransfersService
}

func NewSendSaga(deps *SendSagaDeps) *SendSaga {
	return &SendSaga{
		AccountClient:   deps.AccountClient,
		TransferService: deps.TransferService,
	}
}

func (saga *SendSaga) Start(userFromName, userToName string, amount int64) error {
	_, err := saga.AccountClient.TransferCoins(context.Background(), &pb.TransferCoinsReq{
		UserFromName: userFromName,
		UserToName:   userToName,
		Amount:       amount,
	})
	if err != nil {
		return err
	}
	err = saga.TransferService.Create(userFromName, userToName, amount)
	if err != nil {
		saga.AccountClient.TransferCoins(context.Background(), &pb.TransferCoinsReq{
			UserFromName: userToName,
			UserToName:   userFromName,
			Amount:       amount,
		})
		return err
	}
	return nil
}
