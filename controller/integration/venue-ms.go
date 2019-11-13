package integration

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rockspoon/rs.com.ordering-ms/errors"
	"github.com/rockspoon/rs.cor.common-model/address"
	request "github.com/rockspoon/rs.cor.common-request"
	m "github.com/rockspoon/rs.cor.middleware/v2"
	mmodel "github.com/rockspoon/rs.cor.middleware/v2/model"
	vm "github.com/rockspoon/rs.cor.venue-model/v4/model"
)

const (
	// GetVenueByID is the get venue by id endpoint.
	GetVenueByID = "%svenue/general-information"
)

// VenueMS venue microservice integration point
type VenueMS struct{}

// NewVenueMS creates a new Venue Microservice integration point
func NewVenueMS() VenueMS {
	return VenueMS{}
}

// GeneralInformationResponse fdasfsa
type GeneralInformationResponse struct {
	Name            string                `json:"name"`
	Description     *string               `json:"description"`
	Media           []vm.Media            `json:"media"`
	Logo            *vm.ImageDTO          `json:"logo"`
	Address         *address.Address      `json:"address"`
	OperatingEntity vm.OperatingEntityDTO `json:"operatingEntity"`
}

// GetGeneralInformation get informations about the venue
func (v VenueMS) GetGeneralInformation(ctx context.Context) (*GeneralInformationResponse, error) {
	cData, err := m.DataFromContext(ctx)
	if err != nil {
		return nil, err
	}

	venuePath := cData.Paths[mmodel.MerchantMS]

	req, err := request.NewJSONRequestWithContext(ctx, "GET", fmt.Sprintf(GetVenueByID, venuePath), nil)
	if err != nil {
		return nil, err
	}

	var si GeneralInformationResponse
	if statusCode, err := request.DoJSON(req, &si); err != nil {
		return nil, err
	} else if statusCode != http.StatusOK {
		return nil, errors.InvalidMicroserviceResponse("Venue GeneralInformation", strconv.Itoa(statusCode))
	}
	return &si, nil
}
