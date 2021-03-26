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
		if content, _, err := c.FileRead(volumes[0].Uuid, "cloud-init/cloud-init01", 0, 7000); err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Printf("Bytes read:\n%s", string(content))
		}
	} else {
		fmt.Println("no volumes found found")
	}
}
