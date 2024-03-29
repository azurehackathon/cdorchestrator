package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	//"os/exec"
	"strconv"
	"strings"
)

func create_svc_temp(filename string, service_name string, app_name string, env string, service_type string, app_group string, container_port string, build_number string) {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output  := strings.Replace(string(input), string("<<APP_NAME>>"), string(app_name), -1)
	output  = strings.Replace(output, string("<<ENV_NAME>>"), string(env), -1)
	output  = strings.Replace(output, string("<<TYPE>>"), string(service_type), -1)
	output  = strings.Replace(output, string("<<APP_GROUP>>"), string(app_group), -1)
	output  = strings.Replace(output, string("<<CONTAINER_PORT>>"), container_port, -1)
	output  = strings.Replace(output, string("<<BUILD_NUMBER>>"), build_number, -1)

	sv_temp := fmt.Sprintf("./SVC-%s-%s-%s.yml", string(app_name), build_number, service_name)

	if err = ioutil.WriteFile(sv_temp, []byte(output), 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func create_svc_temp_v2(filename string, args []string) {

	input, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output  := strings.Replace(string(input), string("<<APP_NAME>>"), string(args[1]), -1)
	output  = strings.Replace(output, string("<<ENV_NAME>>"), string(args[2]), -1)
	output  = strings.Replace(output, string("<<TYPE>>"), string(args[3]), -1)
	output  = strings.Replace(output, string("<<APP_GROUP>>"), string(args[4]), -1)
	output  = strings.Replace(output, string("<<CONTAINER_PORT>>"), args[5], -1)
	output  = strings.Replace(output, string("<<BUILD_NUMBER>>"), args[6], -1)
	output  = strings.Replace(output, string("<<PUBLIC_IP>>"), args[7], -1)
	output  = strings.Replace(output, string("<<RESOURCE_GROUP>>"), args[8], -1)

	sv_temp := fmt.Sprintf("./SVC-%s-%s-%s.yml", string(args[1]), args[6], args[0])

	if err = ioutil.WriteFile(sv_temp, []byte(output), 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	var service_name = flag.String("SERVICE_NAME", "beta", "Type of the service - LIVE or BETA.")
	var app_name = flag.String("APP_NAME", "dummy_app", "The Application name for the service.")
	var env = flag.String("ENV", "test", "The Environment name.")
	var service_type = flag.String("TYPE", "type1", "Type of the Service.")
	var app_group = flag.String("APP_GROUP", "group1", "The Application Group name.")
	var container_port = flag.Int("CONTAINER_PORT", 443, "The Container port to be exposed by the service.")
	var build_number = flag.Int("BUILD_NUMBER", 1, "The Build number(generated by jenkins).")
	var public_ip = flag.String("PUBLIC_IP", "127.0.0.1", "Public IP for the service.")
	var ip_resource_group = flag.String("RESOURCE_GROUP", "rs", "Resource group for PIP.")

	flag.Parse()

	args := []string{*service_name, *app_name, *env, *service_type, *app_group, strconv.Itoa(*container_port), strconv.Itoa(*build_number), *public_ip, *ip_resource_group}

	//create_svc_temp("SVC.yml", *service_name, *app_name, *env, *service_type, *app_group, strconv.Itoa(*container_port), strconv.Itoa(*build_number))

	create_svc_temp_v2("./manifest-generator/service_template_gen/SVC.yml", args)

	//cmd := exec.Command("git", "add", "-A")
	//cmd.Dir = "./service_test"
	//out, err := cmd.Output()
	//
	//cmd = exec.Command("git", "commit", "-m", "message1")
	//cmd.Dir = "./service_test"
	//out, err = cmd.Output()
	//
	//cmd = exec.Command("git", "push", "origin", "master")
	//cmd.Dir = "./service_test"
	//out, err = cmd.Output()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(out)
}
