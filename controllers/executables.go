package controllers

import (
	"os"
	"os/exec"
)

func killProcess(hOS string) {
	if hOS == "windows" {
		cmd := exec.Command("taskkill", "/IM", "app-godemon-app-godemon-tmp-generated.exe", "/F")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	} else {
		cmd := exec.Command("killall", "-9", "app-godemon-app-godemon-tmp-generated")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	}
}

func execMOD(hOS string) {
	if hOS == "windows" {
		cmd := exec.Command("app-godemon-app-godemon-tmp-generated.exe")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	} else {
		cmd := exec.Command("./app-godemon-app-godemon-tmp-generated")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		ErrorHandle(err)
	}
}

func execFile(filepath string) {
	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ErrorHandle(err)
}