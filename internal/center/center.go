package center

import(
	"fmt"
	"github.com/TheComputerDan/sentinel_server/internal/data"
)

type centerInfo struct {
	hosts int
	up int
}

func hostStatus(agentID int) bool {
	fmt.Println(agentID)
	//TODO Add status checks function to allow us to check on registered hosts
	return true
}

func hostsUp() int{
	//TODO Add functionality for multiple clients
	dbi := data.AgentDatabase{
		Host:     "localhost",
		Port:     5432,
		User:     "username",
		Password: "password",
		DBName:   "sentinel",
	}
	agentIDs := dbi.GetAgentIDs()
	for _, agentID := range agentIDs{
		if hostStatus(agentID) == true{
			return 1
		} else {
			return 0
		}
	}
	return 0
}


// TODO think about moving the ones that use data.AgentDatabase
// TODO Also think about making them private again and instantiating an empty one and use set and get methods to manage if posisble
func GetCenter() []centerInfo{
	var centerInfoList []centerInfo
	//TODO Add functionality for multiple clients
	dbi := data.AgentDatabase{
		Host:     "localhost",
		Port:     5432,
		User:     "username",
		Password: "password",
		DBName:   "sentinel",
	}

	centerInfoList = append(centerInfoList,centerInfo{hosts: dbi.HostCount(), up: hostsUp()})

	return centerInfoList
}
