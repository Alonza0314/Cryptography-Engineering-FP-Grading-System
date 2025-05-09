package factory

import (
	"ce/backend/util"
)

func LoadCeCfg(cfgPath string) (*CeCfg, error) {
	ceCfg := &CeCfg{}
	if err := util.LoadYaml(cfgPath, ceCfg); err != nil {
		return nil, err
	}
	return ceCfg, nil
}
