package util

import (
	"fmt"
	"net"
	"strconv"

	"github.com/igor-feoktistov/go-ontap-rest/ontap"
)

func DiscoverIscsiLIFs(c *ontap.Client, lunPath string, initiatorSubnet string) (lifs []ontap.IpInterface, err error) {
	var lun *ontap.Lun
	if lun, _, err = c.LunGetByPath(lunPath, []string{"fields=location,svm"}); err != nil {
	        return
	}
	var ipInterfaces []ontap.IpInterface
        if ipInterfaces, _, err = c.IpInterfaceGetIter([]string{"fields=ip,location,svm","enabled=true","state=up","services=data_iscsi"}); err != nil {
    		return
    	}
    	if len(ipInterfaces) == 0 {
		err = fmt.Errorf("DiscoverIscsiLIFs(): no IP interfaces found")
    		return
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Svm.Name == lun.Svm.Name && ipInterface.Location.HomeNode.Name == lun.Location.Node.Name {
    			var netmask int
    			if netmask, err = strconv.Atoi(ipInterface.Ip.Netmask); err != nil {
    				return
    			}
			if fmt.Sprintf("%s/%d", net.ParseIP(ipInterface.Ip.Address).Mask(net.CIDRMask(netmask, 32)), netmask) == initiatorSubnet {
    				lifs = append(lifs, ipInterface)
    				break
    			}
    		}
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Location.HomeNode.Name != lun.Location.Node.Name {
    			var netmask int
    			if netmask, err = strconv.Atoi(ipInterface.Ip.Netmask); err != nil {
    				return
    			}
			if fmt.Sprintf("%s/%d", net.ParseIP(ipInterface.Ip.Address).Mask(net.CIDRMask(netmask, 32)), netmask) == initiatorSubnet {
    				lifs = append(lifs, ipInterface)
    				break
    			}
    		}
    	}
	return
}

func DiscoverNfsLIFs(c *ontap.Client, volumeName string) (lifs []ontap.IpInterface, err error) {
	var volumeNode string
	if volumeNode, _, err = c.PrivateCliVolumeGetNode(volumeName); err != nil {
		return
	}
	var ipInterfaces []ontap.IpInterface
        if ipInterfaces, _, err = c.IpInterfaceGetIter([]string{"fields=ip,location","enabled=true","state=up","services=data_nfs"}); err != nil {
    		return
    	}
    	if len(ipInterfaces) == 0 {
		err = fmt.Errorf("DiscoverNfsLIFs(): no IP interfaces found")
    		return
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Location.HomeNode.Name == volumeNode {
    			lifs = append(lifs, ipInterface)
    		}
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Location.HomeNode.Name != volumeNode {
    			lifs = append(lifs, ipInterface)
    		}
    	}
	return lifs, err
}

func DiscoverNvmeLIFs(c *ontap.Client, namespacePath string, hostSubnet string) (lifs []ontap.IpInterface, err error) {
	var namespace *ontap.NvmeNamespace
	if namespace, _, err = c.NvmeNamespaceGetByPath(namespacePath, []string{"fields=location,svm"}); err != nil {
	        return
	}
	var ipInterfaces []ontap.IpInterface
        if ipInterfaces, _, err = c.IpInterfaceGetIter([]string{"fields=ip,location,svm","enabled=true","state=up","services=data_nvme_tcp"}); err != nil {
    		return
    	}
    	if len(ipInterfaces) == 0 {
		err = fmt.Errorf("DiscoverNvmeLIFs(): no IP interfaces found")
    		return
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Svm.Name == namespace.Svm.Name && ipInterface.Location.HomeNode.Name == namespace.Location.Node.Name {
    			var netmask int
    			if netmask, err = strconv.Atoi(ipInterface.Ip.Netmask); err != nil {
    				return
    			}
			if fmt.Sprintf("%s/%d", net.ParseIP(ipInterface.Ip.Address).Mask(net.CIDRMask(netmask, 32)), netmask) == hostSubnet {
    				lifs = append(lifs, ipInterface)
    				break
    			}
    		}
    	}
    	for _, ipInterface := range ipInterfaces {
    		if ipInterface.Location.HomeNode.Name != namespace.Location.Node.Name {
    			var netmask int
    			if netmask, err = strconv.Atoi(ipInterface.Ip.Netmask); err != nil {
    				return
    			}
			if fmt.Sprintf("%s/%d", net.ParseIP(ipInterface.Ip.Address).Mask(net.CIDRMask(netmask, 32)), netmask) == hostSubnet {
    				lifs = append(lifs, ipInterface)
    				break
    			}
    		}
    	}
	return
}
