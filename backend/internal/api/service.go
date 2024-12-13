package api

import (
	ctx "context"
	"log"

	data "github.com/metao1/creativefabrica/backend/internal/api/data"
	aps "github.com/metao1/creativefabrica/backend/internal/api/proto"
)

type ActiveCreatorsConfig struct {
	aps.UnimplementedCreatorServiceServer
	activeEmails []string
	FilePath     string
}

func (s *ActiveCreatorsConfig) Init() {
	p, err := data.ReadData(s.FilePath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	s.activeEmails = data.CalcActiveCreators(p, 3)
}

func (a *ActiveCreatorsConfig) GetTopActiveCreators(ctx.Context, *aps.TopActiveCreatorsRequest) (*aps.TopActiveCreatorsResponse, error) {
	resp := &aps.TopActiveCreatorsResponse{
		Emails: a.activeEmails,
	}
	return resp, nil
}
