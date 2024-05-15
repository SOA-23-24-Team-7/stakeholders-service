package server

import (
	"context"
	"stakeholders-service/service"
)

type StakeholdersMicroservice struct {
	UnimplementedStakeholdersMicroserviceServer
	JwtService *service.JWTService
}

func (server *StakeholdersMicroservice) GenerateAccessToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	token, err := server.JwtService.GenerateAccsessToken(req.UserId, req.Username, req.PersonId, req.Role)
	if err != nil {
		println("Error while generating token")
		return nil, err
	}

	response := &TokenResponse{
		Id:          token.Id,
		AccessToken: token.AccessToken,
	}

	return response, nil
}
