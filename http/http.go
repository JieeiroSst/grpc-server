package http

import (
	"context"
	"log"

	"github.com/JIeeiroSst/go-app/inventory"
	"github.com/JIeeiroSst/go-app/repositories"
)

type InventoryService struct {
	repo repositories.IventoryRepository
	inventory.UnimplementedCheckInventoryServiceServer
}

func NewInventoryService(repo repositories.IventoryRepository) InventoryService {
	return InventoryService{
		repo : repo,
	}
}

func (i *InventoryService) CheckAndOrder(ctx context.Context,in *inventory.CheckInventory) (*inventory.CheckInventoryResponse,error){
	items:=in.GetItems()
	ok,err:=i.repo.CheckAndOrder(items)
	if err!=nil{
		log.Println(err)
	}
	respone:=&inventory.CheckInventoryResponse{Ok: ok}
	return respone,nil
}