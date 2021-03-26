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
		return
	}
	if len(expPolicies) > 0 {
		parameters = []string{"clients.match=192.168.1.2"}
		if rules, _, err := c.ExportPolicyRuleGetIter(expPolicies[0].GetRef(), parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(rules) > 0 {
				if responseJSON, err := json.MarshalIndent(rules, "", "  "); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("ExportPolicyRuleGet response:\n%s\n", string(responseJSON))
				}
			} else {
				fmt.Println("no export policy rules found")
			}
	    }
	} else {
		fmt.Println("export policy not found")
	}
}
