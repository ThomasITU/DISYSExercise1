package main

import (
	"context"


	
)

const (
	port = ":8080"
)

type Server struct {
}

func main() {

}

func (s *Server) GetTime(ctx context.Context, message *GetTimeRequest) (*Response, error) {
	return &Response{Response: time.dat}
}
