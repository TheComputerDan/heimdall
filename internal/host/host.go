package host

import (
	_ "fmt"
	"github.com/TheComputerDan/sentinel_server/internal/config"
	"log"
	"net"
	"os"
	"os/user"
	"runtime"
)

// Info defines basic information on the host being collected
// by the agent to uniquely identify the machine in the inventory.
type Info struct {
	IP       map[string]string
	OSType   string
	Hostname string
}

//runAsUser determines the user that sentinel is running as.
func runAsUser() (string, error){
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

//getInterfaceName takes agent.yml and loads its values for runtime.
func getInterfaceName() string {
	loaded := config.Load()
	interfaceName := loaded.GetString("interface_name")

	return interfaceName
}

func (hostInfo *Info) Init() {
	hostInfo.IP = IPAddr(getInterfaceName())
	hostInfo.OSType = runtime.GOOS
	hostInfo.Hostname = hostname()
}

//hostname gets the hostname of the device and returns the string
func hostname() string {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return hostname
}

//IPAddr Prints off the IP addresses of each interface
func IPAddr(iface string) map[string]string {

	var addrs map[string]string

	netIface, err := net.InterfaceByName(iface)
	if err != nil {
		log.Println(err)
		return nil
	}

	addresses, err := netIface.Addrs()

	osType := runtime.GOOS
	if osType == "linux" {
		addrs = map[string]string{"ipv6": addresses[1].String(), "ipv4": addresses[0].String()}
	} else if osType == "darwin" {
		addrs = map[string]string{"ipv6": addresses[0].String(), "ipv4": addresses[1].String()}
	} else {
		addrs = map[string]string{"ipv6": addresses[0].String(), "ipv4": addresses[1].String()}
	}

	return addrs
}
