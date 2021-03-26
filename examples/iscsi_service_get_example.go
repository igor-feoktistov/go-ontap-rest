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
	iscsiServices, _, err := c.IscsiServiceGetIter([]string{"enabled=true","fields=target"})
	if err != nil {
		fmt.Println(err)
	} else {
		if len(iscsiServices) > 0 {
			fmt.Printf("Target name=%s\n", iscsiServices[0].Target.Name)
		} else {
			fmt.Println("no iscsi service found")
		}
	}
}
