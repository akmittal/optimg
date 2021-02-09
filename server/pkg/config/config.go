package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	dbPath string

	dbName  string
	appHost string
	appPort int
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbPath, "dbpath", os.Getenv("SQLITE_PATH"), "DB PATH")

	flag.StringVar(&conf.dbName, "dbname", os.Getenv("SQLITE_DB"), "DB name")
	flag.StringVar(&conf.appHost, "apphost", os.Getenv("APP_HOST"), "APP HOST")
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	flag.IntVar(&conf.appPort, "appport", port, "APP PORT")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbPath, c.dbName)
}

func (c *Config) GetAppHost() string {
	return fmt.Sprintf("%v:%v", c.appHost, c.appPort)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"%s/%s",
		c.dbPath,
		dbname,
	)
}
