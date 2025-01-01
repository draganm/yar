package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	cfg := struct {
		output string
	}{}
	app := &cli.App{
		Name:                 "yar",
		Usage:                "yaml archiver",
		Args:                 false,
		ArgsUsage:            "<directory>",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output file",
				EnvVars:     []string{"YAR_OUTPUT"},
				Required:    true,
				Destination: &cfg.output,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("missing directory argument")
			}

			content := map[string]any{}

			dfs := os.DirFS(c.Args().First())
			err := fs.WalkDir(
				dfs,
				".",
				func(path string, d fs.DirEntry, err error) error {
					if d.IsDir() {
						return nil
					}

					f, err := dfs.Open(path)
					if err != nil {
						return fmt.Errorf("error opening file %s: %w", path, err)
					}
					defer f.Close()

					b, err := io.ReadAll(f)
					if err != nil {
						return fmt.Errorf("error reading file %s: %w", path, err)
					}

					content[path] = string(b)

					return nil
				},
			)

			if err != nil {
				return fmt.Errorf("error walking directory: %w", err)
			}

			d, err := yaml.Marshal(content)
			if err != nil {
				return fmt.Errorf("error marshalling content: %w", err)
			}

			err = os.WriteFile(cfg.output, d, 0644)
			if err != nil {
				return fmt.Errorf("error writing output file: %w", err)
			}

			return nil
		},
	}
	app.RunAndExitOnError()
}
