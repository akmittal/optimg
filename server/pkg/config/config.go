package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbPath string

	dbName  string
	appHost string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbPath, "dbpath", os.Getenv("SQLITE_PATH"), "DB PATH")

	flag.StringVar(&conf.dbName, "dbname", os.Getenv("SQLITE_DB"), "DB name")
	flag.StringVar(&conf.appHost, "apphost", os.Getenv("APP_HOST"), "APP HOST")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbPath, c.dbName)
}

func (c *Config) GetAppHost() string {
	return c.appHost
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"%s/%s",
		c.dbPath,
		dbname,
	)
}
