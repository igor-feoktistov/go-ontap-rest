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
	lunPath := "/vol/my_test_vol01/my_test_lun01"
	initiatorSubnet := "192.168.1.192/26"
	if lifs, err := util.DiscoverIscsiLIFs(c, lunPath, initiatorSubnet); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, lif := range lifs {
			fmt.Printf("%s\n", lif.Ip.Address);
		}
	}
}
