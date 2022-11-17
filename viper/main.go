package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type configType struct {
	InitParam InitParam `json:"initParam"`
	// TagLib    TagLib    `json:"tagLib"`
}

type InitParam struct {
	InstallAt  string `json:"installAt"`
	Version    string `json:"version"`
	StaticPath string `json:"staticPath"`
}

type TagLib struct {
	TagLibASD string `json:"tagLibASD"`
	TagLibLoc string `json:"taglibLoc"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	var config configType
	var s string

	viper.SetConfigType("json")
	viper.SetConfigFile("./testJSONConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if viper.IsSet("init-param.installAt") {
		fmt.Println("init-param.installAt:", viper.Get("init-param.installAt"))
	} else {
		fmt.Println("init-param.installAt not set.")
	}

	if viper.IsSet("init-param.staticPath") {
		fmt.Println("init-param.staticPath:", viper.Get("init-param.staticPath"))
	} else {
		fmt.Println("init-param.staticPath is not set.")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Config Loaded:", config)

	if err := viper.UnmarshalKey("init-param.installAt", &s); err != nil {
		log.Fatal(err)
	}

	fmt.Println("key Loaded:", s)
}
