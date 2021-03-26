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
	} else {
		if len(luns) > 0 {
			if _, err := c.LunDelete(luns[0].GetRef()); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
			}
		} else {
			fmt.Println("no LUNs found")
		}
	}
}
