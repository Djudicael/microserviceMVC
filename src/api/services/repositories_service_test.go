package services

import (
	"os"
	"testing"

	"github.com/Djudicael/microserviceMVC/src/api/clients/restclient"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}
