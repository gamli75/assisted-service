package operators

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/swag"
	logutil "github.com/openshift/assisted-service/pkg/log"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

type ValidationReply struct {
	ValidationId string
	IsValid      bool
	Reasons      []string
}

type ManifestReply struct {
	Manifests map[string]string
}

type API interface {
	GetType() models.OperatorType
	ValidateCluster(ctx context.Context, cluster *common.Cluster) (*ValidationReply, error)
	ValidateHost(context.Context, *common.Cluster, *models.Host) (*ValidationReply, error)
	GetCPURequirementForWorker(context.Context, *common.Cluster) (int, error)
	GetCPURequirementForMaster(context.Context, *common.Cluster) (int, error)
	GetMemoryRequirementForWorker(ctx context.Context, cluster *common.Cluster) (int, error)
	GetMemoryRequirementForMaster(ctx context.Context, cluster *common.Cluster) (int, error)
	GenerateManifest(context.Context, *common.Cluster) (*ManifestReply, error)

	//monitor operator status
}

type typeToOperatorMap map[models.OperatorType]API

type OperatorsManager struct {
	log                logrus.FieldLogger
	db                 *gorm.DB
	typeToOperatorsMap typeToOperatorMap
}

func NewOperatorsManager(log logrus.FieldLogger, db *gorm.DB) *OperatorsManager {
	lsoOperator := NewLsoOperator(log, db)

	return &OperatorsManager{
		log: log,
		db:  db,
		typeToOperatorsMap: typeToOperatorMap{
			models.OperatorTypeLso: lsoOperator,
		},
	}
}

func (i *OperatorsManager) IsEnable(ctx context.Context, cluster *models.Cluster, operatorType models.OperatorType) bool {

	log := logutil.FromContext(ctx, i.log)
	result := false
	if cluster.Operators != "" {
		var operators models.Operators
		if err := json.Unmarshal([]byte(cluster.Operators), &operators); err != nil {
			log.WithError(err).Warn("Failed to unmarshal operators")
			return false
		}
		for _, operator := range operators {
			if operator.OperatorType == operatorType && swag.BoolValue(operator.Enabled) {
				result = true
				break
			}
		}
	}
	return result
}
