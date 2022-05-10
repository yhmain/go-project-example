package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhmain/go-project-example/repository"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}
func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, pageInfo)
	// fmt.Printf("%#v", pageInfo.PostList)
	// fmt.Printf("%#v", pageInfo)
	for _, e := range pageInfo.PostList {
		fmt.Printf("%#v", e)
	}
	assert.Equal(t, 5, len(pageInfo.PostList))
}
