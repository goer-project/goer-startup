// Package apiserver does all the work necessary to create a iam APIServer.
package apiserver

import (
	"github.com/marmotedu/iam/pkg/log"

	"goer-startup/internal/goer-apiserver/config"
	"goer-startup/internal/goer-apiserver/options"
	"goer-startup/pkg/app"
)

const commandDesc = `The API server validates and configures data
for the api objects. The API Server services REST operations to do the api objects management.`

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
