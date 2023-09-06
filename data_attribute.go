package intercom

import "fmt"

// DataAttributeService handles interactions with the API through an DataAttributeRepository.
type DataAttributeService struct {
	Repository DataAttributeRepository
}

// A DataAttribute represents a new dataAttribute that happens to a User.
type DataAttribute struct {
	Name     string `json:"name"`
	Model    string `json:"model"`     // The model that the data attribute belongs to. Enum: "contact" "company" "conversation"
	DataType string `json:"data_type"` // The type of data stored for this attribute. Enum: "string" "integer" "float" "boolean" "datetime" "date"

	Description string   `json:"description,omitempty"` // The readable description you see in the UI for the attribute.
	Options     []string `json:"options,omitempty"`     // To create list attributes. Provide a set of hashes with value as the key of the options you want to make. data_type must be string.
}

// Create a new DataAttribute
func (e *DataAttributeService) Create(dataAttribute *DataAttribute) error {
	return e.Repository.create(dataAttribute)
}

func (e DataAttribute) String() string {
	return fmt.Sprintf("[intercom] dataAttribute { name: %s }", e.Name)
}
