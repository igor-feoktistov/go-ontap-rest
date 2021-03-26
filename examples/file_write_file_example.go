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
	buf := []byte("Just test\n")
	n := len(buf)
	if len(volumes) > 0 {
		parameters = []string{}
		if bytesWritten, _, err := c.FileWrite("POST", volumes[0].Uuid, "cloud-init01", parameters, buf[:n]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Bytes written: %d\n", bytesWritten)
		}
	} else {
		fmt.Println("no volumes found found")
	}
}
