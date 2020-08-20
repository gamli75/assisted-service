package cluster

import (
	"net/http"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/pkg/errors"

	"github.com/openshift/assisted-service/models"
)

type validationID models.ClusterValidationID

const (
	IsMachineCidrDefined                = validationID(models.ClusterValidationIDMachineCidrDefined)
	isMachineCidrEqualsToCalculatedCidr = validationID(models.ClusterValidationIDMachineCidrEqualsToCalculatedCidr)
	isApiVipDefined                     = validationID(models.ClusterValidationIDAPIVipDefined)
	isApiVipValid                       = validationID(models.ClusterValidationIDAPIVipValid)
	isIngressVipDefined                 = validationID(models.ClusterValidationIDIngressVipDefined)
	isIngressVipValid                   = validationID(models.ClusterValidationIDIngressVipValid)
	AllHostsAreReadyToInstall           = validationID(models.ClusterValidationIDAllHostsAreReadyToInstall)
	SufficientMastersCount              = validationID(models.ClusterValidationIDSufficientMastersCount)
	IsDNSDomainDefined                  = validationID(models.ClusterValidationIDDNSDomainDefined)
	IsPullSecretSet                     = validationID(models.ClusterValidationIDPullSecretSet)
)

func (v validationID) category() (string, error) {
	switch v {
	case IsMachineCidrDefined, isMachineCidrEqualsToCalculatedCidr, isApiVipDefined, isApiVipValid, isIngressVipDefined,
		isIngressVipValid, IsDNSDomainDefined:
		return "network", nil
	case AllHostsAreReadyToInstall, SufficientMastersCount:
		return "hosts-data", nil
	case IsPullSecretSet:
		return "configuration", nil
	}
	return "", common.NewApiError(http.StatusInternalServerError, errors.Errorf("Unexpected cluster validation id %s", string(v)))
}

func (v validationID) String() string {
	return string(v)
}
