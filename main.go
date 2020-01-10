package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 4 {
		fmt.Println("USAGE : systemd-gen 'service-file-name(.service)' 'logged username' 'working dir' 'exec start'")
		return
	}

	if os.Geteuid() != 0 {
		fmt.Println("please run as sudo")
		return
	}

	name := argsWithoutProg[0]
	user := argsWithoutProg[1]
	directory := argsWithoutProg[2]
	execCmd := argsWithoutProg[3]

	template := `[Unit]
Description=%v Service
After=network.target

[Service]
User=%v
WorkingDirectory=%v
ExecStart=%v
Restart=always

[Install]
WantedBy=multi-user.target`

	serviceFile := fmt.Sprintf(template, name, user, directory, execCmd)

	servicePath := fmt.Sprintf("/etc/systemd/system/%v.service", name)
	ioutil.WriteFile(servicePath, []byte(serviceFile), 0644)

	exec.Command("systemctl", "enable", name).Run()
	exec.Command("service", name, "restart").Run()

	fmt.Println("Service File Created.")
}