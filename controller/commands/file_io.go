package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	log.Printf("[%d] Downloading file '%s'.", req.AgentId, req.Path)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.DownloadFile(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Downloaded file.")
	return resp, nil
}

func (s *AgentManagerServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	log.Printf("[%d] Uploading file '%s'.", req.AgentId, req.Path)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.UploadFile(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Uploaded file.")
	return resp, nil
}
