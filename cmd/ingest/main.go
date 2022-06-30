package main

import(
	"fmt"
)


/*
	Launches Agent Ingest, which is responsible for consuming agent event data. 
	Agent Ingest will also be responsible for routing data to Agent Accountant and Agent Brain. 
	All of this will be enabled via gRPC. 
*/

func main() {
	fmt.Println("Hello Agent")
}