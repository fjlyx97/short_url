package config

type Config struct {
	LogModel  `yaml:"log"`
	BaseModel `yaml:"base"`
}

type BaseModel struct {
	Port int32 `yaml:"port"`
}

type LogModel struct {
	LogPath      string `yaml:"log_path"`
	RotationTime int32  `yaml:"rotation_time"`
}
