package config

import "fmt"

type Config struct {
  Home string
  Path string
  DbFile string
}

func (c *Config) UpdatePath(newPath string) {
  c.Path = newPath
}

func (c *Config) String() string {
  return fmt.Sprintf("%s/%s/%s", c.Home, c.Path, c.DbFile)
}

func NewConfig(home, path, db string) *Config {
  return &Config{
    Home: home,
    Path: path,
    DbFile: db,
  }
}
