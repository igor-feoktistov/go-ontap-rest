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
	volumeName := "my_test_vol01"
	filePath := "/cloud-init/cloud-init01"
	if content, err := util.DownloadFileAPI(c, volumeName, filePath); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("File content:\n--\n%s--\n", content);
	}
}
