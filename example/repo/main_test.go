package repo_test

import (
	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"testing"
)

func TestMain(m *testing.M) {
	viper.SetConfigFile("../../configs/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_, err = ormx.New(viper.GetViper())
	if err != nil {
		panic(err)
	}
	ormx.FlushDB()
	m.Run()
}
