package util

import (
	"fmt"

	"go-ontap-rest/ontap"
)

func GetAggregateMax(c *ontap.Client) (aggregate string, err error) {
	var svms []ontap.Svm
	if svms, _, err = c.SvmGetIter([]string{"fields=aggregates"}); err != nil {
		return
	}
	if len(svms) == 0 {
		err = fmt.Errorf("GetAggregateMax(): no SVMs returned in SvmGetIter()")
		return
	}
	var maxAvailableSize int
	for _, aggr := range svms[0].Aggregates {
		var aggregates []ontap.PrivateCliAggregate
		if aggregates, _, err = c.PrivateCliAggregateGetIter([]string{"aggregate=" + aggr.Name}); err != nil {
			break
		}
		if len(aggregates) > 0 {
			if aggregates[0].State == "online" && *aggregates[0].AvailableSize > maxAvailableSize {
				aggregate = aggregates[0].Name
				maxAvailableSize = *aggregates[0].AvailableSize
			}
		} else {
			err = fmt.Errorf("GetAggregateMax(): no aggregates returned in PrivateCliAggregateGetIter()")
			break
		}
	}
	return
}
