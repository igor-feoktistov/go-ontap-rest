package main

import (
	"fmt"
	"time"
	"encoding/json"

	"go-ontap-rest/ontap"
)

func main() {
	c := ontap.NewClient(
		"https://mytestsvm.example.com",
		&ontap.ClientOptions {
		    BasicAuthUser: "vsadmin",
		    BasicAuthPassword: "secret",
		    SSLVerify: false,
		    Debug: true,
    		    Timeout: 60 * time.Second,
		},
	)
	var parameters []string
	parameters = []string{"name=my_test_vol01"}
	volumes, _, err := c.VolumeGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(volumes) > 0 {
		parameters = []string{"name=my_snapshot_test01"}
		if snapshots, _, err := c.SnapshotGetIter(volumes[0].Uuid, parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(snapshots) > 0 {
				if snapshot, _, err := c.SnapshotGet(snapshots[0].GetRef(), []string{}); err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(snapshot, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("SnapshotGet response:\n%s\n", string(responseJSON))
					}
				}
			} else {
				fmt.Println("no snapshosts found")
			}
		}
	} else {
		fmt.Println("no volumes found found")
	}
}
