package tests

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_Proxy(t *testing.T) {

	bt, err := os.ReadFile("../README.md")
	require.NoError(t, err)

	cli := resty.New()
	for i := 0; i < 251; i++ {
		res, err := cli.R().SetBody(bt).Post("http://localhost:8080/hello")
		require.NoError(t, err)
		fmt.Println(res.String())
	}
}
