package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/rodaine/hclencoder"

	"github.com/juliogreff/datadog-to-terraform/pkg/types"
)

const (
	dashboardResource = "dashboard"
	monitorResource   = "monitor"
)

func main() {
	var err error

	args := os.Args[1:]

	if len(args) != 3 {
		fail("usage: dd2hcl [dashboard|monitor] [ressource_name] [json_path]")
	}

	resourceType := args[0]
	resourceName := args[1]
	jsonFile, err := os.Open(args[2])
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	resource := types.Resource{Name: resourceName}

	switch resourceType {
	case dashboardResource:
		var dashboard *types.Board
		err = json.Unmarshal(byteValue, &dashboard)
		if err != nil {
			fail("%s %s: unable to parse JSON: %s", resourceType, err)
		}

		resource.Type = "datadog_dashboard"
		resource.Board = dashboard
	case monitorResource:
		var monitor *types.Monitor
		err = json.Unmarshal(byteValue, &monitor)
		if err != nil {
			fail("%s %s: unable to parse JSON: %s", resourceType, err)
		}

		resource.Type = "datadog_monitor"
		resource.Monitor = monitor
	}

	hcl, err := hclencoder.Encode(types.ResourceWrapper{
		Resource: resource,
	})
	if err != nil {
		fail("%s %s: unable to encode hcl: %s", resourceType, err)
	}

	fmt.Println(string(hcl))
}

func fail(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}
