package config_test

import (
	"strings"
	"testing"

	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/spf13/viper"
)

func TestEnvFile(t *testing.T) {
	viper.AddConfigPath("./..")
	_, key, secret := config.Init()
	keySplit := strings.Split(key, "")
	secretSplit := strings.Split(secret, "")
	if keySplit[0] != "P" {
		t.Error("Failure to get key value from .env file")
	} else if secretSplit[0] != "d" {
		t.Error("Failure to get secret value from .env file")
	}
}
