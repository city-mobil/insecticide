package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lartie/insecticide/internal/analyzer"
	"github.com/lartie/insecticide/internal/rconf"
	"github.com/lartie/insecticide/internal/rules"
)

var (
	redisConfigFilePath = flag.String("redis-config", "", "Redis config file path")
	redisConfigVersion  = flag.String("redis-version", string(rules.RedisVersion6), "Redis version")
)

func main() {
	flag.Parse()

	redisConf, err := rconf.LoadParameters(*redisConfigFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	confRules := rules.GetParameters(rules.RedisVersion(*redisConfigVersion))
	data := analyzer.MapConfigAndRules(redisConf, confRules)

	report := analyzer.NewAnalyzer(data).Analyze()
	for cmd, infos := range report {
		fmt.Printf("Parameter: %s\n", cmd)
		for _, info := range infos {
			fmt.Printf("[%s]\n", strings.ToUpper(string(info.Level)))
			fmt.Printf("Advice: %s\n", info.Advice)
			fmt.Printf("Reason: %s\n", info.Reason)
		}
		fmt.Println()
	}
}
