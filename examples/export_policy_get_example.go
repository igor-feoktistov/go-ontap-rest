package main

import (
	"fmt"
	"time"
	"encoding/json"

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
	} else {
		fmt.Print(expPolicies)
		if len(expPolicies) > 0 {
			for _, expPolicy := range expPolicies {
				response, _, err := c.ExportPolicyGet(expPolicy.GetRef(), parameters)
				if err != nil {
					fmt.Println(err)
				} else {
					if responseJSON, err := json.MarshalIndent(response, "", "  "); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("ExportPolicyGet response:\n%s\n", string(responseJSON))
					}
				}
			}
		}
	}
}
