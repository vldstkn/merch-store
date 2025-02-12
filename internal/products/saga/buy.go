package saga

import (
	"context"
	"merch-store/internal/interfaces"
	"merch-store/pkg/pb"
)

type Status string

type BuySaga struct {
	AccountClient  pb.AccountClient
	ProductService interfaces.ProductsService
}

type BuySagaDeps struct {
	AccountClient  pb.AccountClient
	ProductService interfaces.ProductsService
}

func NewBuySaga(deps *BuySagaDeps) *BuySaga {
	return &BuySaga{
		AccountClient:  deps.AccountClient,
		ProductService: deps.ProductService,
	}
}

func (saga *BuySaga) Start(userId, productId int64) error {
	price, err := saga.ProductService.GetPriceProduct(productId)
	if err != nil {
		return err
	}
	_, err = saga.AccountClient.DeductBalance(context.Background(), &pb.DeductBalanceReq{
		Amount: price,
		UserId: userId,
	})
	if err != nil {
		return err
	}
	err = saga.ProductService.AddProductToInventory(userId, productId)
	if err != nil {
		saga.AccountClient.Refund(context.Background(), &pb.RefundReq{
			Amount: price,
			UserId: userId,
		})
		return err
	}
	return nil
}
