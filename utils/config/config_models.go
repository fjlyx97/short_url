package config

type Config struct {
	BaseModel      `yaml:"base"`
	LogModel       `yaml:"log"`
	SnowFlakeModel `yaml:"snowflake"`
	MysqlModel     `yaml:"mysql"`
}

type BaseModel struct {
	Port            int32  `yaml:"port"`
	BaseUrl         string `yaml:"base_url"`
	HealthCheckUrl  string `yaml:"health_check_url"`
	HealthCheckPort int32  `yaml:"health_check_port"`
}

type LogModel struct {
	LogPath      string `yaml:"log_path"`
	RotationTime int32  `yaml:"rotation_time"`
}

type SnowFlakeModel struct {
	StartTimeStamp   int64 `yaml:"start_timestamp"`
	DatacenterId     int64 `yaml:"datacenterId"`
	WorkerId         int64 `yaml:"worker_id"`
	SequenceBits     int64 `yaml:"sequence_bits"`
	DataCenterIdBits int64 `yaml:"datacenterid_bits"`
	WorkerIdBits     int64 `yaml:"workerid_bits"`
}

type MysqlModel struct {
	Ip             string `yaml:"ip"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	DatabaseName   string `yaml:"database_name"`
	ConnectTimeout string `yaml:"connect_timeout"`
	ReadTimeout    string `yaml:"read_timeout"`
	WriteTimeout   string `yaml:"write_timeout"`
}
