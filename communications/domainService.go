package communications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

const (
	portsAPI = "/domain/v1/ports"
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
	resp, err := d.client.Get(d.host + portsAPI)
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
	b, err := json.Marshal(ports)
	if err != nil {
		return fmt.Errorf("failed to marhsal struct: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, d.host+portsAPI, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("received bad status code: %d", resp.StatusCode)
	}

	return nil
}
