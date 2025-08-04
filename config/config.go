package config

import (
	"flag"
        log "github.com/sirupsen/logrus"
        "github.com/spf13/viper"
)


type Config struct {
        BindAddress       string  `mapstructure:"bind_address"`
        Port              string  `mapstructure:"listen_port"`
        BaseURL           string  `mapstructure:"url_base"`
        ProxyProtocolPort string  `mapstructure:"proxyprotocol_port"`
        ServerLat         float64 `mapstructure:"server_lat"`
        ServerLng         float64 `mapstructure:"server_lng"`
        IPInfoAPIKey      string  `mapstructure:"ipinfo_api_key"`

        StatsPassword string `mapstructure:"statistics_password"`
        RedactIP      bool   `mapstructure:"redact_ip_addresses"`

        AssetsPath string `mapstructure:"assets_path"`

        DatabaseType     string `mapstructure:"database_type"`
        DatabaseHostname string `mapstructure:"database_hostname"`
        DatabaseName     string `mapstructure:"database_name"`
        DatabaseUsername string `mapstructure:"database_username"`
        DatabasePassword string `mapstructure:"database_password"`

        DatabaseFile string `mapstructure:"database_file"`

        EnableHTTP2 bool   `mapstructure:"enable_http2"`
        EnableTLS   bool   `mapstructure:"enable_tls"`
        TLSCertFile string `mapstructure:"tls_cert_file"`
        TLSKeyFile  string `mapstructure:"tls_key_file"`
}

var (
        configFile   string
	optConfig = flag.String("c", "", "config file to be used, defaults to settings.toml in the same directory")
        loadedConfig *Config = nil
)

func GetConfigFile() string {
	return configFile
}

func Load(configPath string) Config {
        var conf Config

        viper.Reset()
        viper.SetDefault("listen_port", "8989")
        viper.SetDefault("url_base", "")
        viper.SetDefault("proxyprotocol_port", "0")
        viper.SetDefault("download_chunks", 4)
        viper.SetDefault("distance_unit", "K")
        viper.SetDefault("enable_cors", false)
        viper.SetDefault("statistics_password", "PASSWORD")
        viper.SetDefault("redact_ip_addresses", false)
        viper.SetDefault("database_type", "postgresql")
        viper.SetDefault("database_hostname", "localhost")
        viper.SetDefault("database_name", "speedtest")
        viper.SetDefault("database_username", "postgres")
        viper.SetDefault("enable_tls", false)
        viper.SetDefault("enable_http2", false)

        if configPath != "" {
	    configFile = configPath
            viper.SetConfigFile(configFile)
            if err := viper.ReadInConfig(); err != nil {
               log.Fatalf("Error reading config file: %v", err)
            }
            if err := viper.Unmarshal(&conf); err != nil {
               log.Fatalf("Error parsing config: %v", err)
            }
            loadedConfig = &conf
            return conf
            } else {
            log.Warnf("Custom config not found!, falling back to default")
            }

        viper.SetConfigName("settings")
        viper.AddConfigPath("/usr/local/etc/")
        viper.AddConfigPath(".")
        viper.SetConfigType("toml")
        viper.SetEnvPrefix("speedtest")
        viper.AutomaticEnv()

        // Добавлена строка с обработкой ошибки чтения конфига
        if err := viper.ReadInConfig(); err != nil {
            log.Fatalf("Error reading config file: %v", err)
        }
    
        if err := viper.Unmarshal(&conf); err != nil {
            log.Fatalf("Error parsing config: %v", err)
        }

        loadedConfig = &conf
        return conf
}

func init() {
    log.Println("Config initialization")
    flag.Parse()
    if *optConfig != "" {
        log.Printf("Using custom config file: %s", *optConfig)
	} else {
        *optConfig = "settings.toml"
        log.Print("Using default config file: %s",*optConfig)
	}
    configFile=*optConfig
}

func LoadedConfig() *Config {
	if loadedConfig == nil {
	    Load(configFile)
	}
    return loadedConfig
}
