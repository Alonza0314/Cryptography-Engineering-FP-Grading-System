package factory_test

import (
	"ce/backend/factory"
	"reflect"
	"testing"
)

var testCasesLoadCeCfg = []struct {
	name     string
	cfgPath  string
	expected *factory.CeCfg
}{
	{
		name:    "test load ceCfg",
		cfgPath: "../../config/ceCfg.yaml",
		expected: &factory.CeCfg{
			Logger: factory.LoggerCfg{
				FileDir:   "./logs",
				DebugMode: true,
			},
			Gin: factory.GinCfg{
				Port: "31413",
			},
		},
	},
}

func TestLoadCeCfg(t *testing.T) {
	for _, testCase := range testCasesLoadCeCfg {
		t.Run(testCase.name, func(t *testing.T) {
			ceCfg, err := factory.LoadCeCfg(testCase.cfgPath)
			if err != nil {
				t.Fatalf("Failed to load ceCfg: %v", err)
			}
			if !reflect.DeepEqual(ceCfg.Logger, testCase.expected.Logger) {
				t.Errorf("Logger config mismatch, got: %v, want: %v", ceCfg.Logger, testCase.expected.Logger)
			}
			if !reflect.DeepEqual(ceCfg.Gin, testCase.expected.Gin) {
				t.Errorf("Gin config mismatch, got: %v, want: %v", ceCfg.Gin, testCase.expected.Gin)
			}
		})
	}
}
