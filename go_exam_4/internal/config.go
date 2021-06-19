package internal

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

type Configs struct {
	vn         *viper.Viper
	ConfigPath string
	State      string

	Validator *validator.Validate
	TimeZone  string `mapstructure:"time_zone"`

	BangkokLocation *time.Location
}

func (c *Configs) InitAllConfigurations(s string) error {
	if s == "" {
		s = "local"
	}

	name := fmt.Sprintf("config.%s", s)
	log.Infof("config file using : %s", name)

	if c.ConfigPath == "" {
		c.ConfigPath = "./go_exam_4/configs"
	}

	vn := viper.New()
	vn.AddConfigPath(c.ConfigPath)
	vn.SetConfigName(name)
	c.vn = vn
	c.State = s

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}

	loc, err := time.LoadLocation(c.TimeZone)
	if err != nil {
		return errors.Wrapf(err, "load location %s error", c.TimeZone)
	}
	c.BangkokLocation = loc

	log.Infof("all config loaded : %#v", c)
	return nil
}
