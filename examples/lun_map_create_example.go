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
	parameters = []string{}
	lunId := 1
	lunMap := ontap.LunMap{
		Igroup: &ontap.Igroup{
			Resource: ontap.Resource{
                    		Name: "igroup_my_test01",
                        },
		},
		Lun: &ontap.LunRef{
			Resource: ontap.Resource{
                    		Name: "/vol/my_test_vol01/my_test_lun01",
                        },
		},
		LogicalUnitNumber: &lunId,
	}
	_, err := c.LunMapCreate(&lunMap, parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
