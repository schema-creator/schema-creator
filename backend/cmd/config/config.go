package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// LoadEnv はconfigに記述された環境変数に読み取りConfig(グローバル変数)に代入する関数
// 引数にファイルを指定することによって、特定のenvファイルを読み取る。
func LoadEnv(envfile ...string) error {
	if len(envfile) > 0 {
		if err := godotenv.Load(envfile...); err != nil {
			return err
		}
	}

	config := config{}

	if err := env.Parse(&config.App); err != nil {
		return err
	}

	if err := env.Parse(&config.Database); err != nil {
		return err
	}

	if err := env.Parse(&config.Google); err != nil {
		return err
	}

	if err := env.Parse(&config.Github); err != nil {
		return err
	}

	Config = &config

	return nil
}
