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
	parameters = []string{"name=my_test_vol01"}
	volumes, _, err := c.VolumeGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(volumes) > 0 {
		parameters = []string{}
		volume := ontap.Volume{
			Size: 10 * 1024 * 1024 * 1024,
		}
		_, err := c.VolumeModify(volumes[0].GetRef(), &volume, parameters)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success")
		}
	}
}
