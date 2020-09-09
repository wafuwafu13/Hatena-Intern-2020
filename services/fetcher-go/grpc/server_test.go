package grpc

import (
	"context"
	"testing"

	pb "github.com/wafuwafu13/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetch_Normaly(t *testing.T) {
	s := NewServer()
	src := "https://example.com/"
	reply, err := s.Fetch(context.Background(), &pb.FetcherRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, "Example Domain", reply.Title)
}

func Test_Server_Fetch_Anormaly(t *testing.T) {
	s := NewServer()
	src := "https://hogehoge/"
	reply, err := s.Fetch(context.Background(), &pb.FetcherRequest{Src: src})
	assert.Error(t, err)
	// TODO
	assert.Equal(t, "https://hogehoge/", reply.Title)
}
