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
	volumes, _, err := c.VolumeGetIter([]string{})
	if err != nil {
		fmt.Println(err)
	} else {
		if len(volumes) > 0 {
			for _, volume := range volumes {
				parameters := []string{}
				response, _, err := c.VolumeGet(volume.GetRef(), parameters)
				if err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
						fmt.Printf("%s\n", err)
					} else {
						fmt.Printf("VolumeGet response:\n%s\n", string(responseJSON))
					}
				}
			}
		}
	}
}
