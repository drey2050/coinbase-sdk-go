package coinbase

import (
	"context"
	"fmt"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

type ValidatorStatus string

const (
	ValidatorStatusUnknown             ValidatorStatus = "unknown"
	ValidatorStatusProvisioning        ValidatorStatus = "provisioning"
	ValidatorStatusProvisioned         ValidatorStatus = "provisioned"
	ValidatorStatusDeposited           ValidatorStatus = "deposited"
	ValidatorStatusPendingActivation   ValidatorStatus = "pending_activation"
	ValidatorStatusActive              ValidatorStatus = "active"
	ValidatorStatusExiting             ValidatorStatus = "exiting"
	ValidatorStatusExited              ValidatorStatus = "exited"
	ValidatorStatusWithdrawalAvailable ValidatorStatus = "withdrawal_available"
	ValidatorStatusWithdrawalComplete  ValidatorStatus = "withdrawal_complete"
	ValidatorStatusActiveSlashed       ValidatorStatus = "active_slashed"
	ValidatorStatusExitedSlashed       ValidatorStatus = "exited_slashed"
	ValidatorStatusReaped              ValidatorStatus = "reaped"
)

// getAPIValidatorStatus maps the user facing ValidatorStatus to api facing ValidatorStatus
func getAPIValidatorStatus(status *ValidatorStatus) client.ValidatorStatus {
	if status == nil {
		return client.VALIDATORSTATUS_UNKNOWN
	}

	switch *status {
	case ValidatorStatusUnknown:
		return client.VALIDATORSTATUS_UNKNOWN
	case ValidatorStatusProvisioning:
		return client.VALIDATORSTATUS_PROVISIONING
	case ValidatorStatusProvisioned:
		return client.VALIDATORSTATUS_PROVISIONED
	case ValidatorStatusDeposited:
		return client.VALIDATORSTATUS_DEPOSITED
	case ValidatorStatusPendingActivation:
		return client.VALIDATORSTATUS_PENDING_ACTIVATION
	case ValidatorStatusActive:
		return client.VALIDATORSTATUS_ACTIVE
	case ValidatorStatusExiting:
		return client.VALIDATORSTATUS_EXITING
	case ValidatorStatusExited:
		return client.VALIDATORSTATUS_EXITED
	case ValidatorStatusWithdrawalAvailable:
		return client.VALIDATORSTATUS_WITHDRAWAL_AVAILABLE
	case ValidatorStatusWithdrawalComplete:
		return client.VALIDATORSTATUS_WITHDRAWAL_COMPLETE
	case ValidatorStatusActiveSlashed:
		return client.VALIDATORSTATUS_ACTIVE_SLASHED
	case ValidatorStatusExitedSlashed:
		return client.VALIDATORSTATUS_EXITED_SLASHED
	case ValidatorStatusReaped:
		return client.VALIDATORSTATUS_REAPED
	default:
		return client.VALIDATORSTATUS_UNKNOWN
	}
}

type Validator struct {
	model client.Validator
}

func NewValidator(validator client.Validator) Validator {
	return Validator{
		model: validator,
	}
}

func (v Validator) ID() string {
	return v.model.ValidatorId
}

func (v Validator) Status() client.ValidatorStatus {
	return v.model.Status
}

func (v Validator) ToString() string {
	return fmt.Sprintf(
		"Validator { Id: '%s' Status: '%s' }",
		v.ID(),
		v.Status(),
	)
}

func (c *Client) ListValidators(ctx context.Context, networkId string, assetId string, status *ValidatorStatus) ([]Validator, error) {
	listValidatorReq := c.client.ValidatorsAPI.ListValidators(ctx, normalizeNetwork(networkId), assetId)
	listValidatorReq = listValidatorReq.Status(getAPIValidatorStatus(status))

	validatorList, _, err := listValidatorReq.Execute()
	if err != nil {
		return nil, err
	}

	validators := make([]Validator, len(validatorList.GetData()))
	for i, validator := range validatorList.GetData() {
		validators[i] = NewValidator(validator)
	}

	return validators, nil
}

func (c *Client) GetValidator(ctx context.Context, networkId string, assetId string, validatorId string) (Validator, error) {
	validator, _, err := c.client.ValidatorsAPI.GetValidator(
		ctx,
		normalizeNetwork(networkId),
		assetId,
		validatorId,
	).Execute()
	if err != nil {
		return Validator{}, err
	}

	return NewValidator(*validator), nil
}
