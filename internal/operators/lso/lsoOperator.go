package lso

import (
	"context"

	"github.com/openshift/assisted-service/internal/operators"

	"github.com/jinzhu/gorm"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

type lsoOperator struct {
	log logrus.FieldLogger
	db  *gorm.DB
}

func NewLsoOperator(log logrus.FieldLogger, db *gorm.DB) *lsoOperator {
	return &lsoOperator{
		log: log,
		db:  db,
	}
}

func (l *lsoOperator) GetType() models.OperatorType {
	return models.OperatorTypeLso
}

func (l *lsoOperator) ValidateCluster(ctx context.Context, cluster *common.Cluster) (*operators.ValidationReply, error) {
	return nil, nil
}

func (l *lsoOperator) ValidateHost(context.Context, *common.Cluster, *models.Host) (*operators.ValidationReply, error) {
	return nil, nil
}

func (l *lsoOperator) GetCPURequirementForWorker(context.Context, *common.Cluster) (int, error) {
	return 0, nil
}

func (l *lsoOperator) GetCPURequirementForMaster(context.Context, *common.Cluster) (int, error) {
	return 0, nil
}

func (l *lsoOperator) GetMemoryRequirementForWorker(ctx context.Context, cluster *common.Cluster) (int, error) {
	return 0, nil
}

func (l *lsoOperator) GetMemoryRequirementForMaster(ctx context.Context, cluster *common.Cluster) (int, error) {
	return 0, nil
}

func (l *lsoOperator) GenerateManifest(ctx context.Context, c *common.Cluster) (*operators.ManifestReply, error) {
	reply, err := Manifests(c.Cluster.OpenshiftVersion)
	return &operators.ManifestReply{Manifests: reply}, err

}
