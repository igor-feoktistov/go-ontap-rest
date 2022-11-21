package util

import (
	"fmt"

	"github.com/igor-feoktistov/go-ontap-rest/ontap"
)

func GetAggregateMax(c *ontap.Client, svmName string) (aggrName string, spaceAvailable int64, err error) {
	var svms []ontap.Svm
	if svms, _, err = c.SvmGetIter([]string{"name=" + svmName, "fields=aggregates"}); err != nil {
		return
	}
	if len(svms) == 0 {
		err = fmt.Errorf("GetAggregateMax(): no SVMs returned in SvmGetIter()")
		return
	}
	for _, aggr := range svms[0].Aggregates {
                if aggr.State == "online" && aggr.AvailableSize > spaceAvailable {
		        aggrName = aggr.Name
			spaceAvailable = aggr.AvailableSize
		}
	}
	if len(aggrName) == 0 {
	        err = fmt.Errorf("GetAggregateMax(): no aggregates assigned to SVM \"%s\" found", svms[0].Name)
        }
	return
}
