package factory

type CeCfg struct {
	Logger LoggerCfg `yaml:"logger" binding:"required"`
}

type LoggerCfg struct {
	FileDir   string `yaml:"fileDir" binding:"required"`
	DebugMode bool   `yaml:"debugMode" binding:"required"`
}
