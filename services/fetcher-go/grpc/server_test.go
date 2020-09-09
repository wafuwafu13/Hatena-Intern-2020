package grpc

import (
	"context"
	"testing"

	pb "github.com/wafuwafu13/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetch(t *testing.T) {
	s := NewServer()
	src := `foo https://google.com/ bar`
	reply, err := s.Fetch(context.Background(), &pb.FetcherRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, "foo <a href=\"https://google.com/\">https://google.com/</a> bar", reply.Title)
}
