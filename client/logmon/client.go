package logmon

import (
	"context"

	"github.com/hashicorp/nomad/client/logmon/proto"
)

type logmonClient struct {
	client proto.LogMonClient
}

func (c *logmonClient) Start(cfg *LogConfig) error {
	req := &proto.StartRequest{
		LogDir:         cfg.LogDir,
		StdoutFileName: cfg.StdoutLogFile,
		StderrFileName: cfg.StderrLogFile,
		MaxFiles:       uint32(cfg.MaxFiles),
		MaxFileSizeMb:  uint32(cfg.MaxFileSizeMB),
		Uid:            uint32(cfg.UID),
		Gid:            uint32(cfg.GID),
		FifoDir:        cfg.FifoDir,
	}
	_, err := c.client.Start(context.Background(), req)
	return err
}

func (c *logmonClient) Stop() error {
	req := &proto.StopRequest{}
	_, err := c.client.Stop(context.Background(), req)
	return err
}
