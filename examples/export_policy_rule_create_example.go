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
		rule := ontap.ExportPolicyRule{
			AnonymousUser: "root",
			Protocols: []string{"nfs"},
			RoRule: []string{"any"},
			RwRule: []string{"any"},
			Superuser: []string{"any"},
			Clients: []ontap.ExportRuleClient{
				ontap.ExportRuleClient{
					Match: "192.168.20.10",
				},
			},
		}
		if _, err := c.ExportPolicyRuleCreate(expPolicies[0].GetRef(), &rule); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	} else {
		fmt.Println("export policy not found")
	}
}
