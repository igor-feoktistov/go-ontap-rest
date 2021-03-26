package main

import (
	"fmt"
	"time"
	"go-ontap-rest/ontap"
	"go-ontap-rest/util"
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
	lunSrcPath := "/vol/my_test_vol01/my_test_lun01"
	lunDstPath := "/vol/my_test_vol01/my_test_lun02"
	if err := util.LunCopy(c, lunSrcPath, lunDstPath); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
