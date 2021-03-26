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
	//parameters = []string{"fields=name,uuid,state"}
	svms, _, err := c.SvmGetIter(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(svms)
		if len(svms) > 0 {
			//parameters = []string{"fields=ip_interfaces"}
			response, _, err := c.SvmGet(svms[0].GetRef(), parameters)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(response)
			}
		}
	}
}
