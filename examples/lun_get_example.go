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
	luns, _, err := c.LunGetIter(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(luns) > 0 {
			for _, lun := range luns {
				//parameters = []string{"fields=location,movement"}
				response, _, err := c.LunGet(lun.GetRef(), parameters)
				if err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("LunGet response:\n%s\n", string(responseJSON))
					}
				}
			}
		} else {
			fmt.Println("no LUNs found")
		}
	}
}
