package haberdasherserver

import (
	"context"
	"math/rand"

	pb "github.com/mehuled/twirp/haberdasher"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (hat *pb.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}

func (s *Server) ColorHat(ctx context.Context, color *pb.Color) (hat *pb.Hat, err error) {
	if color.Value == "" {
		return nil, twirp.InvalidArgumentError("color value", "I can't make a non coloured hat!!")
	}
	return &pb.Hat{
		Inches: 10,
		Color:  color.Value,
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}

func (s *Server) ColorSizeHat(ctx context.Context, colorsize *pb.ColorSize) (hat *pb.Hat, err error) {
	if colorsize.Value == "" {
		return nil, twirp.InvalidArgumentError("color value", "I can't make a non coloured hat!!")
	}
	if colorsize.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: colorsize.Inches,
		Color:  colorsize.Value,
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
