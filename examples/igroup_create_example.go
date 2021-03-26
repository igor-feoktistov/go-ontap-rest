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
	igroup := ontap.Igroup{
		Resource: ontap.Resource{
			Name: "igroup_my_test01",
		},
		OsType: "linux",
		Protocol: "iscsi",
		Initiators: []ontap.IgroupInitiator{
            		ontap.IgroupInitiator{
                    		Resource: ontap.Resource{
					Name: "iqn.2005-02.com.open-iscsi:initiator01",
                    		},
                    	},
                },
	}
	_, err := c.IgroupCreate(&igroup, []string{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
