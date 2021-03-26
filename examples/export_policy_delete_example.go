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
	parameters = []string{"name=my_test"}
	expPolicies, _, err := c.ExportPolicyGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(expPolicies) > 0 {
		if _, err := c.ExportPolicyDelete(expPolicies[0].GetRef()); err != nil {
			fmt.Print(err)
		} else {
			fmt.Println("success")
		}
	} else {
		fmt.Println("export policy not found")
	}
}
