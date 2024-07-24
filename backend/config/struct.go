package config

type Config struct {
	HTTPPort  string `yaml:"http_port"`
	OriginURL string `yaml:"origin_url"`
	LogPath   string `yaml:"log_path"`
	Ffmpeg    struct {
		AudioWatermark string `yaml:"audio_watermark"`
		VideoWatermark string `yaml:"video_watermark"`
		ImgWatermark   string `yaml:"img_watermark"`
	} `yaml:"ffmpeg"`
	DbConnection struct {
		Server   string `yaml:"server"`
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		DataBase string `yaml:"database"`
		Driver   string `yaml:"driver"`
	} `yaml:"connection"`
	Notifier struct {
		Host   string `yaml:"host"`
		APIKey string `yaml:"api_key"`
	} `yaml:"notifier"`
	Session struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"session"`
	Gateway struct {
		ApiKey string `yaml:"api_key"`
	} `yaml:"gateway"`
}
