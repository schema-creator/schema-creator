package main

import (
	"context"
	"database/sql"
	"flag"
	"strings"

	"github.com/schema-creator/schema-creator/schema-creator/cmd/config"
	"github.com/schema-creator/schema-creator/schema-creator/internal/container"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/server"
	"github.com/schema-creator/schema-creator/schema-creator/internal/route"
	"github.com/schema-creator/schema-creator/schema-creator/pkg/log"
)

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func init() {
	// Usage: eg. go run main.go -e .env -e hoge.env -e fuga.env ...
	var envFile envFlag
	flag.Var(&envFile, "e", "path to .env file \n eg. -e .env -e another.env . ")
	flag.Parse()

	if err := config.LoadEnv(envFile...); err != nil {
		log.Fatal(context.Background(), "lod Env Error", "error", err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), "failed to run", "error", err)
	}
}

func run() error {
	if err := container.NewContainer(); err != nil {
		return err
	}

	db := container.Invoke[*sql.DB]()

	srv := route.NewRouter()

	defer db.Close()

	if err := server.New(config.Config.App.Addr, srv).RunWithGraceful(); err != nil {
		log.Error(context.Background(), "failed to listen server", "error", err)
		return err
	}
	return nil
}
