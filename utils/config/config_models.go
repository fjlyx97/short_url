package config

type Config struct {
	BaseModel      `yaml:"base"`
	LogModel       `yaml:"log"`
	SnowFlakeModel `yaml:"snowflake"`
}

type BaseModel struct {
	Port int32 `yaml:"port"`
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
