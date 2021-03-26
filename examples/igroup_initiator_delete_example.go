package main

import (
	"fmt"
	"time"

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
		parameters = []string{"name=iqn.2005-02.com.open-iscsi:initiator01"}
		if initiators, _, err := c.IgroupInitiatorGetIter(igroups[0].GetRef(), parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(initiators) > 0 {
				if _, err := c.IgroupInitiatorDelete(initiators[0].GetRef()); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("success")
				}
			} else {
				fmt.Println("initiator not found")
			}
	    }
	} else {
		fmt.Println("no igroup found")
	}
}
