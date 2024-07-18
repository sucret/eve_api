package config

type Logger struct {
	Level      string `mapstructure:"level"`
	FilePath   string `mapstructure:"file_path"`
	FileName   string `mapstructure:"file_name"`
	Format     string `mapstructure:"format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}
