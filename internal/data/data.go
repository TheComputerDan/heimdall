package data

import (
	"database/sql"
	"fmt"
	"github.com/TheComputerDan/sentinel_server/internal/host"
	_ "github.com/lib/pq"
	"strings"
)

const (
	SentinelSchema = "sentinel"
	AgentTable      = "agents"
	UsersTable     = "users"
)

type agent struct {
	agentID int
	hostname string
	ipaddr string 	// As to accommodate both ipv4 and ipv6
	os string
}


type AgentDatabase struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func (dbi AgentDatabase) Open() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbi.Host, dbi.Port, dbi.User, dbi.Password, dbi.DBName)
	db, err := sql.Open("postgres",psqlInfo)
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func (dbi AgentDatabase) GetAgentIDs() []int{
	var agentIDs []int

	db := dbi.Open()
	agentIDQuery := fmt.Sprintf("SELECT \"agentID\" FROM %s.%s ", SentinelSchema, AgentTable )
	rows, err := db.Query(agentIDQuery)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for rows.Next(){
		var agentID int
		rows.Scan(&agentID)
		agentIDs = append(agentIDs,agentID)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return agentIDs
}
func (dbi AgentDatabase) nextAvailableID() int {

	var startID int = 1

	unavailableIDs := dbi.GetAgentIDs()

	for _, unavailableID := range unavailableIDs {
		if unavailableID == startID{
			startID++
		} else {
			continue
		}
	}

	return startID
}

//func updateAgents(){}

func (dbi AgentDatabase) registerServer() {
	var exists bool

	serverInfo := host.Info{}
	serverInfo.Init()

	id := 1 // The server should always be agentID 1
	ipv4 := serverInfo.IP["ipv4"]
	ip := strings.Split(ipv4,"/")
	hostname := serverInfo.Hostname
	os := serverInfo.OSType

	db := dbi.Open()

	//TODO Better check (More Go friendly...)
	duplicateCheckQuery := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s.%s WHERE \"agentID\"=%d)", SentinelSchema, AgentTable, id)
	duplicateCheck, err := db.Query(duplicateCheckQuery)
	if err != nil {
		db.Close()
		panic(err)
	}

	for duplicateCheck.Next(){
		err = duplicateCheck.Scan(&exists)
		if err != nil {
			panic(err)
		}
	}

	//TODO Figure out if there is a better way to handle the IP Address
	if exists == false {
		insertQuery := fmt.Sprintf("INSERT INTO %s.%s (\"agentID\",ipaddr,hostname,os) VALUES (%d,'%s','%s','%s');", SentinelSchema, AgentTable, id, ip[0], hostname, os)

		_, err := db.Exec(insertQuery)
		if err != nil {
			panic(err)
		}
		defer db.Close()

	}
}

func (dbi AgentDatabase) registerClient(id int, ip string, hostname string, os string) sql.Result {
	//TODO add a client to the database
	db := dbi.Open()
	insertQuery := fmt.Sprintf("INSERT INTO %s.%s (\"agentID\",ipaddr,hostname,os) VALUES (%d,'%s','%s','%s');", SentinelSchema, AgentTable, id, ip,hostname, os)

	// TODO Add check to make sure insert worked.
	result, err := db.Exec(insertQuery)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return result
}

func (dbi AgentDatabase) getAllHosts() []agent{
	//TODO get all hosts from hosts table
	var agents []agent

	db := dbi.Open()

	agentQuery := fmt.Sprintf("select * from %s.%s;", SentinelSchema, AgentTable)

	rows, err := db.Query(agentQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		singeAgent := agent{}

		err := rows.Scan(&singeAgent.agentID, &singeAgent.ipaddr, &singeAgent.hostname, &singeAgent.os)
		if err != nil{
			panic(err)
		}
		agents = append(agents, singeAgent)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return agents
}

//HostCount counts the number of registered hosts in the AgentTable by agentID
func (dbi AgentDatabase) HostCount() int {
	var count int

	db := dbi.Open()
	countQuery := fmt.Sprintf("select count(\"agentID\") from %s.%s", SentinelSchema, AgentTable)

	rows, err := db.Query(countQuery)
	for rows.Next(){
		err:=rows.Scan(&count)
		if err != nil{
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return count
}


func Test(){

	dbi := AgentDatabase{
		Host:     "localhost",
		Port:     5432,
		User:     "username",
		Password: "password",
		DBName:   "sentinel",
	}

	fmt.Printf("Host Count: %d\n", dbi.HostCount())

	//fmt.Printf("Agent IDs: %d\n", dbi.GetAgentIDs())
	//dbi.registerServer()
	//dbi.registerClient(dbi.nextAvailableID(),"10.120.1.57","Dans-Fake-Mac","darwin") // This worked
	//fmt.Println(dbi.getAllHosts())
}
