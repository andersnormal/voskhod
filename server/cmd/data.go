package cmd

import (
	"path"

	"github.com/andersnormal/voskhod/server/config"
	"github.com/andersnormal/voskhod/utils"
)

func mkdirDataFolder(cfg *config.Config) error {
	dirs := []string{
		cfg.DataDir,
		path.Join(cfg.DataDir, cfg.EtcdDataDir),
		path.Join(cfg.DataDir, cfg.NatsDataDir),
	}

	// create data
	for _, d := range dirs {
		if err := utils.MkdirFolder(d); err != nil {
			return err
		}
	}

	return nil
}
