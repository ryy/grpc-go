package service

import (
	"context"
	"user/pb"
)

type RockPaperScissorsService struct {
	Name string
}

func NewRockPaperScissorsService() *RockPaperScissorsService {
	return &RockPaperScissorsService{
		Name: "Tom",
	}
}

func (s *RockPaperScissorsService) GetUser(ctx context.Context, req *pb.Empty) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "Tom",
	}, nil
}
