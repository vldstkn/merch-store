package mappers

import (
	"merch_store/internal/api/dto"
	"merch_store/pkg/pb"
)

func InventoryFromGrpcToDto(inventory *pb.Inventory) dto.Inventory {
	return dto.Inventory{
		Type:     inventory.Type,
		Quantity: inventory.Quantity,
	}
}

func InventoryRepFromGrpcToDto(inventory []*pb.Inventory) []dto.Inventory {
	if inventory == nil {
		return []dto.Inventory{}
	}
	res := make([]dto.Inventory, len(inventory))
	for i, el := range inventory {
		res[i] = InventoryFromGrpcToDto(el)
	}
	return res
}

func FromReceivedRepGrpcToDto(rec []*pb.Received) []dto.Received {
	if rec == nil {
		return []dto.Received{}
	}
	res := make([]dto.Received, len(rec))
	for i, el := range rec {
		res[i] = dto.Received{
			FromUser: el.FromUser,
			Amount:   el.Amount,
		}
	}
	return res
}

func FromSentRepGrpcToDto(sent []*pb.Sent) []dto.Sent {
	if sent == nil {
		return []dto.Sent{}
	}
	res := make([]dto.Sent, len(sent))
	for i, el := range sent {
		res[i] = dto.Sent{
			ToUser: el.ToUser,
			Amount: el.Amount,
		}
	}
	return res
}

func CoinsHistoryFromGrpcToDto(history *pb.CoinsHistory) dto.CoinsHistory {
	if history == nil {
		return dto.CoinsHistory{
			Received: []dto.Received{},
			Sent:     []dto.Sent{},
		}
	}
	return dto.CoinsHistory{
		Received: FromReceivedRepGrpcToDto(history.Received),
		Sent:     FromSentRepGrpcToDto(history.Sent),
	}
}
