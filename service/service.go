package service

import (
	"fmt"

	"github.com/NajmiddinAbdulhakim/iman/api-gateway/config"
	pb "github.com/NajmiddinAbdulhakim/iman/api-gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/internal/resolver"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	GetService() pb.GetServiceClient
	CRUDService() pb.CRUDServiceClient
}

type serviceManager struct {
	getService  pb.GetServiceClient
	crudService pb.CRUDServiceClient
}

func (s *serviceManager) GetService() pb.GetServiceClient {
	return s.getService
}

func (s *serviceManager) CRUDService() pb.CRUDServiceClient {
	return s.crudService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connGet, err := grpc.Dial(
		fmt.Sprintf(`%s:%d`, conf.GetServiceHost, conf.GetServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}	

	connCRUD, err := grpc.Dial(
		fmt.Sprintf(`%s:%d`, conf.CRUDServiceHost, conf.CRUDServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		getService: pb.NewGetServiceClient(connGet),
		crudService: pb.NewCRUDServiceClient(connCRUD),	
	}

	return serviceManager, nil

}
