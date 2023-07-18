package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestAreaCollectDateHanler(t *testing.T) {
  t.Run("リクエストパラメタが不正のときエラーを返す", func(t *testing.T) {
    invalidRequest := events.APIGatewayProxyRequest{}

    _, err := handler(invalidRequest)
    fmt.Println(err)
    assert.Error(t, err)
  })
}
