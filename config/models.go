package config


//postgres presents postgres schema
type postgres struct {
	Username           string `yaml:"USERNAME"`
	Password           string `yaml:"PASSWORD"`
	Port               string `yaml:"PORT"`
	Host               string `yaml:"HOST"`
	DBName             string `yaml:"DB_NAME"`
	MaxOpenConnections int    `yaml:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `yaml:"MAX_IDLE_CONNECTIONS"`
}

//redis presents postgres schema
type redis struct {
	Addr     string `yaml:"ADDRESS"`
	Password string `yaml:"PASSWORD"`
}

//SMTP server config
type smtpServer struct {
	Host     string `yaml:"HOST"`
	Port     int    `yaml:"PORT"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
}

//app presents app model schema
type app struct {
	Port       int    `yaml:"PORT"`
	LogLevel   string `yaml:"LOG_LEVEL"`
	LogPath    string `yaml:"LOG_PATH"`
	QueueCount int    `yaml:"DEBT_COUNT"`
}
