package hostcommands

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/hardware"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

type diskCleanupCmd struct {
	baseCmd
	hwValidator      hardware.Validator
	diskCleanupImage string
}

func NewDiskCleanupCmd(log logrus.FieldLogger, diskCleanupImage string, hwValidator hardware.Validator) *diskCleanupCmd {
	return &diskCleanupCmd{
		baseCmd:          baseCmd{log: log},
		diskCleanupImage: diskCleanupImage,
		hwValidator:      hwValidator,
	}
}

func (c *diskCleanupCmd) GetSteps(_ context.Context, host *models.Host) ([]*models.Step, error) {
	bootDevice, err := hardware.GetBootDevice(c.hwValidator, host)
	if err != nil {
		return nil, err
	}

	// Skip cleanup if SaveDiskPartitions is set
	if hostutil.SaveDiskPartitionsIsSet(host.InstallerArgs) {
		return nil, nil
	}

	args, err := c.GetArgs(bootDevice)
	if err != nil {
		return nil, err
	}

	step := &models.Step{
		StepType: models.StepTypeInstallationDiskCleanup,
		Args:     args,
	}
	return []*models.Step{step}, nil
}

func (c *diskCleanupCmd) GetArgs(bootDevice string) ([]string, error) {

	request := models.DiskCleanupRequest{
		Path: swag.String(bootDevice),
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		c.log.WithError(err).Errorf("failed to marshal DiskCleanupRequest")
		return nil, err
	}

	arguments := []string{
		string(requestBytes),
	}

	return arguments, nil
}
