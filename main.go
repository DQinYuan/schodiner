package main

import (
	"context"
	"github.com/pingcap/schrodinger/pkg/test_client"
	"log"
	"os"
	"time"
)

const (
	managerAddrEnv = "MANAGER_ADDR"
	testIDEnv      = "TEST_ID"
	tidbServiceEnv = "TIDB_SERVICE"
	pdServiceEnv   = "PD_SERVICE"
	tidbsEnv       = "TIDBS"
	pdsEnv         = "PDS"
	tikvsEnv       = "TIKVS"
	// KVName is tikv name
	KVName = "tikv"
	// DBName is tidb name
	DBName = "tidb"
	// PDName is pd name
	PDName = "pd"
)

func main() {
	// for local test
	tidbServiceStr := "172.16.30.11:32333"
	tidbStr := "10.5.4.2:4000,10.5.4.3:4000"
	pdStr := "10.5.4.4:2379,10.5.4.5:2379,10.5.4.6:2379"
	tikvStr := "10.5.4.7:22334,10.5.4.8:22334,10.5.4.9:22334"

	if err := os.Setenv(tidbServiceEnv, tidbServiceStr); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv(tidbsEnv, tidbStr); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv(pdsEnv, pdStr); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv(tikvsEnv, tikvStr); err != nil {
		log.Fatal(err)
	}

	cli := client.New(context.Background(), client.ModeOnline)
	cli.KillWithRecover(context.Background(), "tikv", 25 * time.Second)

}
