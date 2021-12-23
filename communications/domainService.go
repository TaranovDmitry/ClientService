package communications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

const (
	portsAPI = "domain/v1/ports"
)

type DomainService struct {
	host   string
	client *http.Client
}

func NewDomain(h string, c *http.Client) *DomainService {
	return &DomainService{
		host:   h,
		client: c,
	}
}

func (d DomainService) PortsFromDomainMS() (entity.Ports, error) {
	resp, err := d.client.Get(fmt.Sprintf("%s/%s", d.host, portsAPI))
	if err != nil {
		return nil, fmt.Errorf("failed to call domain ports GET API %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var ports entity.Ports
	err = json.Unmarshal(b, &ports)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarhsal response: %w", err)
	}

	return ports, nil
}

func (d DomainService) UpdatePortsInDomainMS(ports entity.Ports) error {
	return nil
}
