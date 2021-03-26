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
	parameters = []string{}
	lunMaps, _, err := c.LunMapGetIter(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(lunMaps) > 0 {
			for _, lunMap := range lunMaps {
				parameters = []string{}
				response, _, err := c.LunMapGet(lunMap.GetRef(), parameters)
				if err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("LunMapGet response:\n%s\n", string(responseJSON))
					}
				}
			}
		} else {
			fmt.Println("no LUN maps found")
		}
	}
}
