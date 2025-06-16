package shared

import "github.com/docker/docker/api/types/container"

const (
	ObservabilityNetworkName = "observability-net"
	LocalnetNetworkName      = "kt-localnet"
)

var (
	NetworkMode = container.NetworkMode(ObservabilityNetworkName)
	Labels      = map[string]string{"stack": "localnet-observability"}
)
