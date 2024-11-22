package config

type Config struct {
	Env       string
	Kitex     Kitex    `yaml:"kitex" yaml:"kitex"`
	MySQLWk   MySQL    `mapstructure:"mysql-wk" yaml:"mysql-wk"`
	MySQLYoda MySQL    `mapstructure:"mysql-yoda" yaml:"mysql-yoda"`
	Redis     Redis    `mapstructure:"redis" yaml:"redis"`
	Registry  Registry `mapstructure:"registry" yaml:"registry"`
	Etcd      Etcd     `mapstructure:"etcd" yaml:"etcd"`
}

type MySQL struct {
	DSN string `mapstructure:"dsn" yaml:"dsn"`
}

type Redis struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

type Kitex struct {
	Service         string `mapstructure:"service" yaml:"service"`
	Address         string `mapstructure:"address" yaml:"address"`
	EnablePprof     bool   `mapstructure:"enable_pprof" yaml:"enable_pprof"`
	EnableGzip      bool   `mapstructure:"enable_gzip" yaml:"enable_gzip"`
	EnableAccessLog bool   `mapstructure:"enable_access_log" yaml:"enable_access_log"`
	LogLevel        string `mapstructure:"log_level" yaml:"log_level"`
	LogFileName     string `mapstructure:"log_file_name" yaml:"log_file_name"`
	LogMaxSize      int    `mapstructure:"log_max_size" yaml:"log_max_size"`
	LogMaxBackups   int    `mapstructure:"log_max_backups" yaml:"log_max_backups"`
	LogMaxAge       int    `mapstructure:"log_max_age" yaml:"log_max_age"`
}

type Registry struct {
	RegistryAddress []string `mapstructure:"registry_address" yaml:"registry_address"`
	Username        string   `mapstructure:"username" yaml:"username"`
	Password        string   `mapstructure:"password" yaml:"password"`
}

type Etcd struct {
	Address string `mapstructure:"address" yaml:"address"`
}
