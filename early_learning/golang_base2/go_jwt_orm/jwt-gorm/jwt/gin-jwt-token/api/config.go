/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:config.go
*Author:Xuebiao Xu
*Date:2020年03月15日
*Description:
*
================================================================*/
package api

import (
    "fmt"
    "github.com/spf13/viper"
)

type DBConfig struct {
    Host    string    `json:"host"` 
    Port    int64     `json:"port"`
    User    string    `json:"user"`
    Passwd  string    `json:"passwd"`
    Database  string  `json:"database"`
}

var (
    DbConfig DBConfig
)

func ParserConfig() {
    config := viper.New()
    config.AddConfigPath("./config")
    config.SetConfigName("config")
    config.SetConfigType("ini")

    if err := config.ReadInConfig(); err != nil {
        panic(err)
    }

    fmt.Println(config.GetString("ss.host"))
    fmt.Println(config.GetInt64("ss.port"))
    DbConfig.Host = config.GetString("ss.host")
    DbConfig.Port = config.GetInt64("ss.port")
    DbConfig.User = config.GetString("ss.user")
    DbConfig.Passwd = config.GetString("ss.passwd")
    DbConfig.Database = config.GetString("ss.database")
}
