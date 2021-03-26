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
		parameters = []string{}
		initiator := ontap.IgroupInitiator{
			IgroupInitiators: &[]ontap.Resource{
                    		ontap.Resource{
					Name: "iqn.2005-02.com.open-iscsi:initiator01",
                    		},
            		},

		}
		if _, err := c.IgroupInitiatorCreate(igroups[0].GetRef(), &initiator); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	} else {
		fmt.Println("no igroup found")
	}
}
