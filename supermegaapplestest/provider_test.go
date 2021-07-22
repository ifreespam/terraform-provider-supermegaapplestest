package apples_old

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestAccFirst(t *testing.T) {
	skipTestIfNoTFAccFlag(t)
	testRandomEnvAvailable(t)
	fmt.Println("[DEBUG] TestAccFirst start")
	time.Sleep(time.Second * 30)
	fmt.Println("[DEBUG] TestAccFirst done")
}

func TestAccSecond(t *testing.T) {
	skipTestIfNoTFAccFlag(t)
	testRandomEnvAvailable(t)
	fmt.Println("[DEBUG] TestAccSecond start")
	time.Sleep(time.Second * 30)
	fmt.Println("[DEBUG] TestAccSecond done")
}

func testRandomEnvAvailable(t *testing.T) {
	if os.Getenv("RANDOM_ENV") != "VAL" {
		t.Fatal("RANDOM_ENV is not VAL!")
	}
}

func skipTestIfNoTFAccFlag(t *testing.T) {
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("Skipped unless env 'TF_ACC' set")
	}
}
