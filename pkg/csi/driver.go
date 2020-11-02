package csi_driver

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang/glog"

	"github.com/batazor/shortlink/internal/logger"
)

// NewDriver returns a CSI plugin that contains the necessary gRPC
// interfaces to interact with Kubernetes over unix domain sockets for
// managing ShortLink Storage
func NewDriver(log logger.Logger, driverName, nodeID, endpoint string, ephemeral bool, maxVolumesPerNode int64, version string) (*hostPath, error) {
	if driverName == "" {
		return nil, errors.New("no driver name provided")
	}

	if nodeID == "" {
		return nil, errors.New("no node id provided")
	}

	if endpoint == "" {
		return nil, errors.New("no driver endpoint provided")
	}
	if version != "" {
		vendorVersion = version
	}

	if err := os.MkdirAll(dataRoot, 0750); err != nil {
		return nil, fmt.Errorf("failed to create dataRoot: %v", err)
	}

	glog.Infof("Driver: %v ", driverName)
	glog.Infof("Version: %s", vendorVersion)

	return &hostPath{
		name:              driverName,
		version:           vendorVersion,
		nodeID:            nodeID,
		endpoint:          endpoint,
		ephemeral:         ephemeral,
		maxVolumesPerNode: maxVolumesPerNode,
		log:               log,
	}, nil
}
