package config

type Config struct {
	LogModel `yaml:"log"`
}

type LogModel struct {
	LogPath      string `yaml:"log_path"`
	RotationTime int32  `yaml:"rotation_time"`
}
