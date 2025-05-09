package factory

type CeCfg struct {
	Logger LoggerCfg `yaml:"logger"`
}

type LoggerCfg struct {
	FileDir string `yaml:"fileDir"`
	DebugMode bool `yaml:"debugMode"`
}
