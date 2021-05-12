package analyzer

import (
	"fmt"

	"github.com/city-mobil/insecticide/internal/rules"
)

type Param struct {
	Key   string
	Args  []string
	Rules rules.Rules
}

type Data []Param

type Analyzer struct {
	Data Data
}

type Infos []*rules.Info
type Report map[string]Infos

func (p *Analyzer) Analyze() Report {
	report := make(Report)
	for _, param := range p.Data {
		for _, rule := range param.Rules {
			var info *rules.Info
			argsLen := uint8(len(param.Args))
			if argsLen == 0 || (argsLen-1) < rule.Index {
				if !rule.Required {
					continue
				}
				info = &rules.Info{
					Advice: fmt.Sprintf("Parameter %s is not set. Index: %d", param.Key, rule.Index),
					Reason: fmt.Sprintf("Set variable for Param %s and Index %d", param.Key, rule.Index),
					Level:  rules.Critical,
				}
			} else {
				info = rule.Check(param.Args)
				if info.Level == rules.Ok {
					continue
				}
			}
			report[param.Key] = append(report[param.Key], info)
		}
	}
	return report
}

func NewAnalyzer(data Data) *Analyzer {
	return &Analyzer{
		Data: data,
	}
}
