package galxe

import "github.com/ethereum/go-ethereum/common"

const galxeURL = "https://graphigo.prd.galaxy.eco/query"

const (
	operationTypeAppend = "APPEND"
)

type (
	// SubmitAddressRequest is the input to the Galxe API when submitting a participant.
	SubmitAddressRequest struct {
		CredentialId int64
		Operation    string
		Address      common.Address
	}

	// SubmitAddressResponse is the response from the Galxe API when submitting a participant.
	SubmitAddressResponse struct {
		Data   SubmitAddressResponseData `json:"data"`
		Errors []GraphQLError            `json:"errors"`
	}

	SubmitAddressResponseData struct {
		Eligibility Eligibility `json:"credentialItems"`
	}

	Eligibility struct {
		Eligible int64 `json:"eligible"`
	}

	// Submitter is the interface for submitting participants to the Galxe API.
	Submitter interface {
		SubmitParticipant(address common.Address) (eligible bool, err error)
	}
)

type (
	// GraphQLError is the error response from the Galxe API.
	GraphQLError struct {
		Message    string     `json:"message"`
		Locations  []Location `json:"locations"`
		Path       []string   `json:"path"`
		Extensions Extensions `json:"extensions"`
	}

	Location struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	}

	Extensions struct {
		Code string `json:"code"`
	}
)
