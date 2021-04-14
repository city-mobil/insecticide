package rconf

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type RedisConf map[string][]string

func LoadParameters(redisConfigPath string) (RedisConf, error) {
	if redisConfigPath == "" {
		return nil, errors.New("redis config path is empty")
	}
	file, err := os.Open(redisConfigPath)
	if err != nil {
		return nil, err
	}

	conf := make(RedisConf)
	scanner := bufio.NewScanner(file)
	var param []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		param = strings.Split(line, " ")
		conf[param[0]] = param[1:]
	}
	return conf, nil
}
