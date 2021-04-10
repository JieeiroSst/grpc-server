package main

import (
	"log"
	"net"

	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/http"
	"github.com/JIeeiroSst/go-app/inventory"
	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"google.golang.org/grpc"
)

func main(){
	dao:=mysql.NewMysqlConnRepo(&config.Config.MysqlConfig)
	inventoryServiceServer:=http.NewInventoryService(dao)

	lis,err:=net.Listen("tcp",config.Config.URL)
	if err!=nil{
		log.Println("programming run error",err)
	}
	log.Println("server running....",config.Config.URL)
	grpcServer := grpc.NewServer()
	if grpcServer != nil {
			log.Println("ket noi client khong thanh cong")
	}else {
		log.Println("ket noi client thanh cong")
	}

	inventory.RegisterCheckInventoryServiceServer(grpcServer,&inventoryServiceServer)
	grpcServer.Serve(lis)

}