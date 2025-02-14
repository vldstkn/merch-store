package saga

import (
	"context"
	"merch_store/internal/interfaces"
	"merch_store/pkg/pb"
)

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

func (saga *BuySaga) Start(userName, productType string) error {
	price, err := saga.ProductService.GetPriceProduct(productType)
	if err != nil {
		return err
	}
	_, err = saga.AccountClient.DeductBalance(context.Background(), &pb.DeductBalanceReq{
		Amount:   price,
		UserName: userName,
	})
	if err != nil {
		return err
	}
	err = saga.ProductService.AddProductToInventory(userName, productType)
	if err != nil {
		saga.AccountClient.Refund(context.Background(), &pb.RefundReq{
			Amount:   price,
			UserName: userName,
		})
		return err
	}
	return nil
}
