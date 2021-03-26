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
	parameters = []string{"name=igroup_my_test01"}
	igroups, _, err := c.IgroupGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(igroups) > 0 {
		parameters = []string{}
		if initiators, _, err := c.IgroupInitiatorGetIter(igroups[0].GetRef(), parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(initiators) > 0 {
				if responseJSON, err := json.MarshalIndent(initiators, "", "  "); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("IgroupInitiatorGet response:\n%s\n", string(responseJSON))
				}
			} else {
				fmt.Println("no igroup initiators found")
			}
	    }
	} else {
		fmt.Println("no igroup found")
	}
}
