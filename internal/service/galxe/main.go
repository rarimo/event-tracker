package galxe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/identity/event-tracker/internal/config"
	"io"
	"net/http"
)

// galxeSubmitter is the implementation of the Submitter interface.
type galxeSubmitter struct {
	httpClient *http.Client
	settings   config.GalxeSubmitterSettings
}

// NewSubmitter returns a new instance of the Submitter interface.
func NewSubmitter(settings config.GalxeSubmitterSettings) Submitter {
	return galxeSubmitter{httpClient: http.DefaultClient, settings: settings}
}

// SubmitParticipant submits a participant to the Galxe API.
func (s galxeSubmitter) SubmitParticipant(address common.Address) (eligible bool, err error) {
	if s.settings.SkipSubmit {
		return true, nil
	}

	reqBody, err := s.createSubmitAddressMutationBody(SubmitAddressRequest{
		CredentialId: s.settings.CredentialId,
		Operation:    operationTypeAppend,
		Address:      address,
	})
	if err != nil {
		return eligible, errors.Wrap(err, "failed to create request body")
	}

	req, err := http.NewRequest(http.MethodPost, galxeURL, reqBody)
	if err != nil {
		return eligible, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", s.settings.AccessToken)

	var response SubmitAddressResponse

	if err = s.perform(req, &response); err != nil {
		return eligible, errors.Wrap(err, "failed to perform request")
	}

	if len(response.Errors) > 0 {
		return eligible, errors.New(fmt.Sprintf("unexpected error: %s", response.Errors[0].Message))
	}

	switch response.Data.Eligibility.Eligible {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return eligible, errors.New(fmt.Sprintf("unexpected eligibility: %d", response.Data.Eligibility.Eligible))

	}
}

func (s galxeSubmitter) perform(r *http.Request, dst interface{}) error {
	resp, err := s.httpClient.Do(r)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(resp.Body)

	if resp.StatusCode > 200 {
		return errors.New(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	if err = json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}

	return nil
}

// createPrepareParticipateMutationBody creates the request body for the Galxe API.
func (s galxeSubmitter) createSubmitAddressMutationBody(input SubmitAddressRequest) (io.Reader, error) {
	mutation := fmt.Sprintf(
		"mutation {\n  credentialItems(\n    input: {\n      credId: \"%v\"\n      operation: %s\n      items: [\n        \"%s\"\n      ]\n    }\n  ) {\n    eligible(address:\"%s\")\n  }\n}",
		input.CredentialId,
		input.Operation,
		input.Address.Hex(),
		input.Address.Hex(),
	)

	raw, err := json.Marshal(map[string]string{
		"query": mutation,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request body")
	}

	return bytes.NewBuffer(raw), nil
}
