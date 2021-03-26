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
	volume := ontap.Volume{
		Resource: ontap.Resource{
			Name: "my_test_vol01",
		},
		Svm: &ontap.Resource{
			Name: "mytestsvm",
		},
		Aggregates: []ontap.Resource{
			ontap.Resource{
				Name: "aggr_01",
			},
		},
		Encryption: &ontap.Encryption{
			Enabled: false,
		},
		Guarantee: &ontap.VolumeSpaceGuarantee{
			Type: "volume",
		},
		Nas: &ontap.Nas{
			ExportPolicy: &ontap.ExportPolicyRef{
				Resource: ontap.Resource{
					Name: "my_test_exp01",
				},
			},
			Path: "/my_test_vol01",
		},
		Size: 1024 * 1024 * 1024,
	}
	_, err := c.VolumeCreate(&volume, parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
