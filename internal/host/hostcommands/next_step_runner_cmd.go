package hostcommands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/openshift/assisted-service/internal/common"
)

type NextStepRunnerConfig struct {
	ServiceBaseURL       string
	InfraEnvID           string
	HostID               string
	UseCustomCACert      bool
	NextStepRunnerImage  string
	SkipCertVerification bool
}

func GetNextStepRunnerCommand(config *NextStepRunnerConfig) (string, *[]string) {

	arguments := []string{"run", "--rm", "-ti", "--privileged", "--pid=host", "--net=host",
		"-v", "/dev:/dev:rw", "-v", "/opt:/opt:rw",
		"-v", "/run/systemd/journal/socket:/run/systemd/journal/socket",
		"-v", "/var/log:/var/log:rw",
		"-v", "/run/media:/run/media:rw",
		"-v", "/etc/pki:/etc/pki"}

	if config.UseCustomCACert {
		arguments = append(arguments, "-v", fmt.Sprintf("%s:%s", common.HostCACertPath, common.HostCACertPath))
	}

	arguments = append(arguments,
		"--env", "PULL_SECRET_TOKEN",
		"--env", "CONTAINERS_CONF",
		"--env", "CONTAINERS_STORAGE_CONF",
		"--env", "HTTP_PROXY", "--env", "HTTPS_PROXY", "--env", "NO_PROXY",
		"--env", "http_proxy", "--env", "https_proxy", "--env", "no_proxy",
		"--name", "next-step-runner", config.NextStepRunnerImage, "next_step_runner",
		"--url", strings.TrimSpace(config.ServiceBaseURL), "--infra-env-id", config.InfraEnvID, "--host-id", config.HostID,
		"--agent-version", config.NextStepRunnerImage, fmt.Sprintf("--insecure=%s", strconv.FormatBool(config.SkipCertVerification)))

	if config.UseCustomCACert {
		arguments = append(arguments, "--cacert", common.HostCACertPath)
	}

	return "podman", &arguments
}
