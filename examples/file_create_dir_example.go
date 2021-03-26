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
	unixPermissions := 755
	if len(volumes) > 0 {
		file := ontap.FileInfo{
			Type: "directory",
			UnixPermissions: &unixPermissions,
		}
		if response, _, err := c.FileCreate(volumes[0].Uuid, "repo", &file); err != nil {
			fmt.Println(err)
		} else {
			if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("FileCreate response:\n%s\n", string(responseJSON))
			}
		}
	} else {
		fmt.Println("no volumes found found")
	}
}
