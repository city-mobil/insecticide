package rules

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type RedisVersion string

const (
	RedisVersion6 RedisVersion = "6"
)

var data map[RedisVersion]Parameters

//go:embed rules.json
var rulesJson []byte

func init() {
	data = make(map[RedisVersion]Parameters)
	err := json.Unmarshal(rulesJson, &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Condition string

const (
	Equal       Condition = "eq"
	NotEqual    Condition = "ne"
	LessThan    Condition = "lt"
	GreaterThan Condition = "gt"
)

type Level string

const (
	Ok         Level = "ok"
	Warn       Level = "warning"
	Critical   Level = "critical"
	Deprecated Level = "deprecated"
)

type Info struct {
	Advice string `json:"advice"`
	Reason string `json:"reason"`
	Level  Level  `json:"level"`
}

func NewDefaultInfo() *Info {
	return &Info{
		Level: Ok,
	}
}

type Rule struct {
	Index     uint8     `json:"index"`
	Value     string    `json:"value"`
	Required  bool      `json:"required"`
	Info      *Info     `json:"info"`
	Condition Condition `json:"condition"`
}

type Rules []*Rule

func (r *Rule) Check(args []string) *Info {
	value := args[r.Index]
	rvType := "string"

	rValue, err := strconv.Atoi(r.Value)
	if err == nil {
		rvType = "int"
	}

	var info *Info
	switch rvType {
	case "int":
		ivalue, err := strconv.Atoi(value)
		if err != nil {
			info = &Info{
				Advice: "Use type integer",
				Reason: fmt.Sprintf("Can not parse integer value: %s", err.Error()),
				Level:  Critical,
			}
			break
		}
		info = r.checkNumbers(ivalue, rValue)
	case "string":
		info = r.checkString(value, r.Value)
	}
	return info
}

func (r *Rule) checkString(lv, rv string) *Info {
	info := NewDefaultInfo()
	switch r.Condition {
	case Equal:
		if lv == rv {
			info = r.Info
		}
	case NotEqual:
		if lv != rv {
			info = r.Info
		}
	default:
		return &Info{
			Advice: "Use correct condition: eq or ne",
			Reason: fmt.Sprintf("unsupported condition: %s", r.Condition),
			Level:  Critical,
		}
	}
	return info
}

func (r *Rule) checkNumbers(lv, rv int) *Info {
	info := NewDefaultInfo()
	switch r.Condition {
	case Equal:
		if lv == rv {
			info = r.Info
		}
	case NotEqual:
		if lv != rv {
			info = r.Info
		}
	case LessThan:
		if lv < rv {
			info = r.Info
		}
	case GreaterThan:
		if lv > rv {
			info = r.Info
		}
	default:
		return &Info{
			Advice: "Use correct condition: eq, ne, lt or gt",
			Reason: fmt.Sprintf("unsupported condition: %s", r.Condition),
			Level:  Critical,
		}
	}
	return info
}

type Parameters map[string]Rules

func GetParameters(version RedisVersion) Parameters {
	return data[version]
}
