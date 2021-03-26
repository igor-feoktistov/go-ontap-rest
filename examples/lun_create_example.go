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
	lunSize := 1024 * 1024 * 1024
	lun := ontap.Lun{
		Resource: ontap.Resource{
			Name: "/vol/my_test_vol01/my_test_lun01",
		},
		Location: &ontap.LunLocation{
			LogicalUnit: "my_test_lun01",
			Volume: &ontap.Resource{
				Name: "my_test_vol01",
			},
		},
		Svm: &ontap.Resource{
			Name: "mytestsvm",
		},
		OsType: "linux",
		Space: &ontap.LunSpace{
			Size: &lunSize,
		},
	}
	_, err := c.LunCreate(&lun, parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
