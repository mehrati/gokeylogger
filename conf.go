package main

import (
	"github.com/spf13/viper"
)

var conf *viper.Viper

func init() {
	conf = viper.New()
	conf.SetConfigType("toml")
	conf.SetConfigName("conf")
	conf.AddConfigPath(".")
	// conf.Set("cron_duration", "@every 60s") // send mail every 60 second
	// conf.Set("host_mail", "hostmail")
	// conf.Set("port_mail", 587)
	// conf.Set("user_mail", "user mail")
	// conf.Set("pass_mail", "password")
	// conf.Set("from_mail", "addr from mail")
	// conf.Set("to_mail", "addr to mail")
	// conf.Set("subject_mail", "keylogger")
	// conf.WriteConfig()
	conf.ReadInConfig()
}
