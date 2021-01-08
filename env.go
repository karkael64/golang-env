package env

import (
	"bufio"
	"os"
	"strings"
	"syscall"
)

var envs map[string]string

func init() {
	envs, _ = load()
}

func splitLine(line string) (string, string) {
	var pos = strings.Index(line, "=")
	if -1 == pos {
		return line, ""
	} else {
		var key = strings.ToUpper(strings.TrimSpace(line[0:pos]))
		var value = line[pos+1:]
		return key, value
	}
}

func load() (map[string]string, error) {
	var result = make(map[string]string)

	var envs []string = syscall.Environ()
	for _, line := range envs {
		var key, value = splitLine(line)
		result[key] = value
	}

	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var key, value = splitLine(line)
		result[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func LoadEnv() map[string]string {
	return envs
}

func GetEnv(name string) string {
	var nameUpper = strings.ToUpper(name)
	return envs[nameUpper]
}

func HasEnv(name string) bool {
	var nameUpper = strings.ToUpper(name)
	for key := range envs {
		if key == nameUpper {
			return true
		}
	}
	return false
}
