package intercom

import "github.com/stefanoschrs/go-intercom/interfaces"

// DataAttributeRepository defines the interface for working with DataAttributes through the API.
type DataAttributeRepository interface {
	create(*DataAttribute) error
}

// DataAttributeAPI implements DataAttributeRepository
type DataAttributeAPI struct {
	httpClient interfaces.HTTPClient
}

func (api DataAttributeAPI) create(dataAttribute *DataAttribute) error {
	_, err := api.httpClient.Post("/data_attributes", dataAttribute)
	return err
}
