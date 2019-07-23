package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"os"
)

func create_RC_temp(filename string, app_name string, env string, app_group string, container_port int, build_number int, replca_count int, volumes string, volume_mounts string, args string, memory_requests string, cpu_requests string, memory_limits string, cpu_limits string) {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bn := strconv.Itoa(build_number)
	cp := strconv.Itoa(container_port)
	rc := strconv.Itoa(replca_count)

	fmt.Println(bn, cp)

	output  := strings.Replace(string(input), string("<<APP_NAME>>"), string(app_name), -1)
	output  = strings.Replace(output, string("<<ENV_NAME>>"), string(env), -1)
	output  = strings.Replace(output, string("<<APP_GROUP>>"), string(app_group), -1)
	output  = strings.Replace(output, string("<<CONTAINER_PORT>>"), cp, -1)
	output  = strings.Replace(output, string("<<BUILD_NUMBER>>"), bn, -1)
	output  = strings.Replace(output, string("<<REPLICA_COUNT>>"), rc, -1)
	output  = strings.Replace(output, string("<<VOLUMES>>"), string(volumes), -1)
	output  = strings.Replace(output, string("<<VOLUME_MOUNTS>>"), string(volume_mounts), -1)
	output  = strings.Replace(output, string("<<ARGS>>"), string(args), -1)
	output  = strings.Replace(output, string("<<MEMORY_REQUESTS>>"), string(memory_requests), -1)
	output  = strings.Replace(output, string("<<CPU_REQUESTS>>"), string(cpu_requests), -1)
	output  = strings.Replace(output, string("<<MEMORY_LIMITS>>"), string(memory_limits), -1)
	output  = strings.Replace(output, string("<<CPU_LIMITS>>"), string(cpu_limits), -1)

	rc_temp := fmt.Sprintf("./RC-%s-%d.yml", string(app_name), build_number)

	if err = ioutil.WriteFile(rc_temp, []byte(output), 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {


	var app_name = flag.String("APP_NAME", "dummy_app", "The Application name for the service.")
	var env = flag.String("ENV", "test", "The Environment name.")
	var volumes = flag.String("VOLUMES", "type1", "Volumes for the RC.")
	var app_group = flag.String("APP_GROUP", "group1", "The Application Group name.")
	var container_port = flag.Int("CONTAINER_PORT", 443, "The Container port of the RC.")
	var build_number = flag.Int("BUILD_NUMBER", 1, "The Build number(generated by jenkins).")
	var replica_count = flag.Int("REPLICA_COUNT", 1, "Replica Count for RC.")
	var volume_mounts = flag.String("VOLUME_MOUNTS", "type1", "Volume mounts for the RC.")
	var args = flag.String("ARGS", "hello", "ARGS for the RC.")
	var mem_req = flag.String("MEMORY_REQUESTS", "520mb", "Memory request.")
	var cpu_req = flag.String("CPU_REQUESTS", "20Gb", "CPU request.")
	var mem_lim = flag.String("MEMORY_LIMITS", "512mb", "Memory Limit.")
	var cpu_lim = flag.String("CPU_LIMITS", "200g", "CPU Limit.")

	flag.Parse()

	create_RC_temp("RC.yml", *app_name, *env, *app_group, *container_port, *build_number, *replica_count, *volumes, *volume_mounts, *args, *mem_req, *cpu_req, *mem_lim, *cpu_lim)

}