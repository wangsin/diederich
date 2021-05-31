package mconf

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(GlobalViper.GetString("main.port"))
}
