package apiserver

import "goer-startup/internal/goer-apiserver/config"

// Run runs the specified APIServer. This should never exit.
func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
