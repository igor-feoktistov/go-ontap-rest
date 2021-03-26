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
	parameters = []string{"fields=name,enabled,state,ip,services,svm,location", "enabled=true", "state=up", "services=data_iscsi"}
	ipInterfaces, _, err := c.IpInterfaceGetIter(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(ipInterfaces)
	}
}
