package config

import(
    "os"
    "encoding/json"
    "io/ioutil"

)

var config Config

type DbConfig struct {
	Host	string `json:"host"`
	Port	int	`json:"port"`
	User	string `json:"user"`
	Password	string `json:"password"`
	DbName	string	`json:"db_name"`
}

type Config struct {
	Port	int `json:"port"`
	Db	DbConfig	`json:"db"`
}

func init() {
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/config/config.json")
    if err != nil {
        panic("error opening config file: " + err.Error())
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &config)
}

func GetDbHost() string {
	return config.Db.Host
}

func GetDbPort() int {
	return config.Db.Port
}

func GetDbUser() string {
	return config.Db.User
}

func GetDbPassword() string {
	return config.Db.Password
}

func GetDbName() string {
	return config.Db.DbName
}

func GetPort() int {
	return config.Port
}