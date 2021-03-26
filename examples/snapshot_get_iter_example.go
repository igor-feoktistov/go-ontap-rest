package main

import (
	"fmt"
	"time"
	"encoding/json"

	"go-ontap-rest/ontap"
)

func main() {
	var parameters []string
	parameters = []string{"name=my_test_vol01"}
	volumes, _, err := c.VolumeGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(volumes) > 0 {
		parameters = []string{}
		if snapshots, _, err := c.SnapshotGetIter(volumes[0].Uuid, parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(snapshots) > 0 {
				for _, snapshot := range snapshots {
					if responseJSON, err := json.MarshalIndent(snapshot, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("SnapshotGetIter response:\n%s\n", string(responseJSON))
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
