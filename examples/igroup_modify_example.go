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
	} else {
		if len(igroups) > 0 {
			igroup := ontap.Igroup{
				Resource: ontap.Resource{
					Name: "igroup_my_test02",
				},
			}
			if _, err := c.IgroupModify(igroups[0].GetRef(), &igroup); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
			}
		}
	}
}
