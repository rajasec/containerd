package server

import (
	"context"
	"os"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/sys"
)

const (
	// DefaultRootDir is the default location used by containerd to store
	// persistent data
	DefaultRootDir = "/var/lib/containerd"
	// DefaultStateDir is the default location used by containerd to store
	// transient data
	DefaultStateDir = "/run/containerd"
	// DefaultAddress is the default unix socket address
	DefaultAddress = "/run/containerd/containerd.sock"
	// DefaultDebuggAddress is the default unix socket address for pprof data
	DefaultDebugAddress = "/run/containerd/debug.sock"
)

// apply sets config settings on the server process
func apply(ctx context.Context, config *Config) error {
	if config.Subreaper {
		log.G(ctx).Info("setting subreaper...")
		if err := sys.SetSubreaper(1); err != nil {
			return err
		}
	}
	if config.OOMScore != 0 {
		log.G(ctx).Infof("changing OOM score to %d", config.OOMScore)
		if err := sys.SetOOMScore(os.Getpid(), config.OOMScore); err != nil {
			return err
		}
	}
	return nil
}
