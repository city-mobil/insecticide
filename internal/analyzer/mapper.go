package analyzer

import (
	"github.com/city-mobil/insecticide/internal/rconf"
	"github.com/city-mobil/insecticide/internal/rules"
)

func MapConfigAndRules(conf rconf.RedisConf, params rules.Parameters) Data {
	var data Data
	for parameter, rule := range params {
		data = append(data, Param{
			Key:   parameter,
			Args:  conf[parameter],
			Rules: rule,
		})
	}
	return data
}
