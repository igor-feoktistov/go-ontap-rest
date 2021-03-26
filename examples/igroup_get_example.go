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
	igroups, _, err := c.IgroupGetIter(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(igroups) > 0 {
			for _, igroup := range igroups {
				response, _, err := c.IgroupGet(igroup.GetRef(), []string{})
				if err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("IgroupGet response:\n%s\n", string(responseJSON))
					}
				}
			}
		} else {
			fmt.Println("no igroups found")
		}
	}
}
