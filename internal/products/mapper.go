package products

import (
	"merch_store/internal/models"
	"merch_store/pkg/pb"
)

func InventoryFromModelToGrpc(inventory *models.Inventory) *pb.Inventory {
	return &pb.Inventory{
		Type:     string(inventory.Type),
		Quantity: inventory.Quantity,
	}
}
func InventoryFromModelsToGrpc(inventory []models.Inventory) []*pb.Inventory {
	var res []*pb.Inventory
	for _, el := range inventory {
		res = append(res, InventoryFromModelToGrpc(&el))
	}
	return res
}
