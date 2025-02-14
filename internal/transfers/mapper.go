package transfers

import (
	"merch_store/internal/models"
	"merch_store/pkg/pb"
)

func ReceivedFromModelToGrpc(rec []models.Received) []*pb.Received {
	res := make([]*pb.Received, len(rec))
	for i, el := range rec {
		res[i] = &pb.Received{
			FromUser: el.FromUser,
			Amount:   el.Amount,
		}
	}
	return res
}

func SentFromModelToGrpc(sent []models.Sent) []*pb.Sent {
	res := make([]*pb.Sent, len(sent))
	for i, el := range sent {
		res[i] = &pb.Sent{
			ToUser: el.ToUser,
			Amount: el.Amount,
		}
	}
	return res
}

func HistoryFromModelToGrpc(history models.History) *pb.CoinsHistory {
	return &pb.CoinsHistory{
		Received: ReceivedFromModelToGrpc(history.Received),
		Sent:     SentFromModelToGrpc(history.Sent),
	}
}
