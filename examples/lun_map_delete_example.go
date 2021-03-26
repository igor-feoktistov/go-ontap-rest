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
	parameters = []string{"name=/vol/my_test_vol01/my_test_lun01"}
	luns, _, err := c.LunGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(luns) == 0 {
		fmt.Println("no LUNs found")
		return
	}
	parameters = []string{"name=igroup_my_test01"}
	igroups, _, err := c.IgroupGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(igroups) == 0 {
		fmt.Println("no igroups found")
		return
	}
	if _, err := c.LunMapDelete(luns[0].Uuid, igroups[0].Uuid); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
