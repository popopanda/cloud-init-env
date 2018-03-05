package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Environments struct {
	Environment map[string]interface{} `json:"environment"`
}

func main() {

	const input string = "/var/lib/cloud/instance/user-data.txt"
	const output string = "/etc/environment"

	result := parseYML(input)

	file, err := os.OpenFile(output, os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for k, v := range result.Environment {
		_, err := file.WriteString(fmt.Sprintf("%v=\"%v\"\n", k, v))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseYML(filename string) Environments {
	var env Environments
	reader, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(buf, &env)

	return env
}
