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
		parameters = []string{"type=directory","return_metadata=true"}
		if files, response, err := c.FileGetIter(volumes[0].Uuid, "repo", parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(files) > 0 {
				for _, file := range files {
					if responseJSON, err := json.MarshalIndent(file, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("FileGetIter response:\n%s\n", string(responseJSON))
					}
				}
			} else {
				fmt.Println("no files found")
			}
		}
	} else {
		fmt.Println("no volumes found found")
	}
}
