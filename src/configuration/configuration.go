package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

func InitConfig[Cfg any](configPath string, config *Cfg) (err error) {
	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	viper.SetConfigFile(configPath)

	// Set environment variable prefix for nested configs
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	BindEnvs(config, "")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(config); err != nil {
		return
	}
	return
}

func BindEnvs(iface interface{}, parentKey string) {
	t := reflect.TypeOf(iface)
	v := reflect.ValueOf(iface)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			continue
		}

		fullKey := tag
		if parentKey != "" {
			fullKey = parentKey + "." + tag
		}

		// Handle nested structs
		if fieldVal.Kind() == reflect.Struct {
			BindEnvs(fieldVal.Addr().Interface(), fullKey)
			continue
		}

		// Bind environment variable
		err := viper.BindEnv(fullKey)
		if err != nil {
			panic("error binding env: " + err.Error())
		}
	}
}
