package host

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"os"
	"runtime"
)

type Info struct {
	IPAddr   string
	OSType   string
	Hostname string
}

func loadConfig() string {
	viper.SetConfigName("heimdall.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("docker_agent/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	iface := viper.GetString("interface_name")
	return iface
}

func HostInfo() Info {
	ifaceName := loadConfig()

	hInfo := Info{
		IPAddr:   IPAddr(ifaceName),
		OSType:   OperatingSystem(),
		Hostname: Hostname(),
	}
	return hInfo
}

// OperatingSystem simply returns `runtime.GOOS`
func OperatingSystem() string {
	return runtime.GOOS
}

func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

//IPAddr Prints off the IP addresses of each interface
func IPAddr(iface string) string {
	netIface, err := net.InterfaceByName(iface)
	// ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	// fmt.Println(netIface)

	addresses, err := netIface.Addrs()

	// for k, v := range addresses {
	// 	fmt.Printf("Interface Address #%v : %v\n", k, v.String())
	// }
	return addresses[1].String()
}
