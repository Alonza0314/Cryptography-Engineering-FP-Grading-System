package factory

type CeCfg struct {
	Logger LoggerCfg `yaml:"logger" binding:"required"`
	Gin    GinCfg    `yaml:"gin" binding:"required"`
	Admin  AdminCfg  `yaml:"admin" binding:"required"`
}

type LoggerCfg struct {
	FileDir   string `yaml:"fileDir" binding:"required"`
	DebugMode bool   `yaml:"debugMode" binding:"required"`
}

type GinCfg struct {
	Port string `yaml:"port" binding:"required"`
}

type AdminCfg struct {
	Ta []TA `yaml:"ta" binding:"required"`
}

type TA struct {
	TaId     string `yaml:"ta_id" binding:"required"`
	TaPassword string `yaml:"ta_password" binding:"required"`
}
