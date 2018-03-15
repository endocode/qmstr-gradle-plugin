package master

import (
	"errors"
	"log"

	"github.com/QMSTR/qmstr/pkg/service"
)

type serverPhaseBuild struct {
	genericServerPhase
}

func (phase *serverPhaseBuild) Activate() error {
	return nil
}

func (phase *serverPhaseBuild) Shutdown() error {
	phase.db.AwaitBuildComplete()
	return nil
}

func (phase *serverPhaseBuild) GetPhaseId() int32 {
	return phase.phaseId
}

func (phase *serverPhaseBuild) Build(in *service.BuildMessage) (*service.BuildResponse, error) {
	for _, node := range in.FileNodes {
		log.Printf("Adding file node %s", node.Path)
		phase.db.AddNode(node)
	}
	return &service.BuildResponse{Success: true}, nil
}

func (phase *serverPhaseBuild) GetConfig(in *service.ConfigRequest) (*service.ConfigResponse, error) {
	return nil, errors.New("Wrong phase")
}

func (phase *serverPhaseBuild) GetNodes(in *service.NodeRequest) (*service.NodeResponse, error) {
	return nil, errors.New("Wrong phase")
}

func (phase *serverPhaseBuild) SendNodes(in *service.AnalysisMessage) (*service.AnalysisResponse, error) {
	return nil, errors.New("Wrong phase")
}

func (phase *serverPhaseBuild) Report(in *service.ReportRequest, streamServer service.ReportService_ReportServer) error {
	return errors.New("Wrong phase")
}