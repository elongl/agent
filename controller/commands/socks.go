package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) StartSocksServer(ctx context.Context, req *pb.StartSocksServerRequest) (*pb.StartSocksServerResponse, error) {
	log.Printf("[%d] Starting SOCKS server.", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.StartSocksServer(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Started SOCKS server.")
	return resp, nil
}

func (s *AgentManagerServer) StopSocksServer(ctx context.Context, req *pb.StopSocksServerRequest) (*pb.StopSocksServerResponse, error) {
	log.Printf("[%d] Stopping SOCKS server.", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.StopSocksServer(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Stopped SOCKS server.")
	return resp, nil
}
