package main

import (
  "fmt"
  "os"
  "log"
  "os/exec"
  _ "errors"
)


func parseArgsAsRoleCommand(args []string)(containerName string){
  containerName = args[2]
  return containerName
}

func CheckContainer(stage string, targetContainerName string) bool{
  return true
}

func getGlobalIP(targetContainerName string,stage string) string {
  return ""
}

func wantedlyIndicate() bool {
	return os.Args[1] == ""
}

func checkArgsLen() bool {
  return len(os.Args) > 2
  }


func main(){
  fmt.Println("s is ssh for wantedly")
  if !checkArgsLen(){
    log.Println("Usage: s <stage> <containerName>")
    return
  }
  if wantedlyIndicate(){
    stage := ""
    sshOption := "-i"
    sshUser := "core"
    privateSSHKeyPath := ""
    targetContainerName := parseArgsAsRoleCommand(os.Args)
    if CheckContainer(stage, targetContainerName) {
      globalIP := getGlobalIP(targetContainerName, stage)
      err := exec.Command("ssh", sshOption, privateSSHKeyPath, sshUser + "@" + globalIP).Run()
	    if err != nil {
	      log.Fatal(err)
	    }
    } else {
      log.Printf("Not found Containeer: " + targetContainerName)
    }
  }
}
