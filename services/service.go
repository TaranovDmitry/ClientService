package services

import (
	"fmt"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

type DomainService interface {
	PortsFromDomainMS() (entity.Ports, error)
	UpdatePortsInDomainMS(ports entity.Ports) error
}

type Service struct {
	domainService DomainService
}

func NewService(ds DomainService) *Service {
	return &Service{
		domainService: ds,
	}
}

func (s Service) Ports() (entity.Ports, error) {
	ports, err := s.domainService.PortsFromDomainMS()
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}
	return ports, err
}

func (s Service) UpdatePorts(ports entity.Ports) error {
	err := s.domainService.UpdatePortsInDomainMS(ports)
	if err != nil {
		return fmt.Errorf("faild to upload ports: %w", err)
	}
	return nil
}
