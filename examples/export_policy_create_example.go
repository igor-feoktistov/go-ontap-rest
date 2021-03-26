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
	expPolicy := ontap.ExportPolicy{
		ExportPolicyRef: ontap.ExportPolicyRef{
			Resource: ontap.Resource{
				Name: "my_test",
			},
		},
	}
	_, err := c.ExportPolicyCreate(&expPolicy, parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
