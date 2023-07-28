package test

import (
	"testing"

	"github.com/zhojiew/appmesh-load/utils"
)

func TestCpuLoad(t *testing.T) {
	precision := 250_000_000_0
	utils.CpuLoad(precision)
}
