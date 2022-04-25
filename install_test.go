package anet

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInstallEnc(t *testing.T) {
	var args InstallArgs
	data, _ := json.Marshal(args)
	fmt.Println(string(data))
}
