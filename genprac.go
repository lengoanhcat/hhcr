package main

import (
	"errors"
	"fmt"
	"time"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// gpcc: General Practicer Chaincode implementation
type gpcc struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(gpcc))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *gpcc) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	var id, name, address string // Entities	

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	// Initialize IDs
	// Put data into world stage
	err = stub.PutState("id", []byte(args[0]))
	if err != nil {
		return nil,err
	}

	err = stub.PutState("name", []byte(args[1]))
	if err != nil {
		return nil,err
	}
	
	err = stub.PutState("address", []byte(args[2]))
	if err != nil {
		return nil,err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *gpcc) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}

	// Read in the name of the appointment to be updated
	if function == "diagnose" {
		gp_id := args[0]
		stub.PutState("GP",[]byte(gp_id))		
	}
		
	// request consultation and book an appointment
	if function == "request_gp_consultation" { 
		var doa time.Time		
		const RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
		doa,_ = time.Parse(RFC850,args[0])
		stub.PutState("DOA",[]byte(doa.String())) // date of appointment
	}

	if function == "cancel_appointment" { // cancel an appointment		
		stub.PutState("DOA",[]byte(""))
	}	

	if function == "" { //
	}
	
	fmt.Println("invoke did not find func: " + function)//error

	return nil, errors.New("Received unknown function invocation: " + function)
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions	
	if function == "query_doa" {//read a variable		
		name,_ := stub.GetState("Name")
		doa,_ := stub.GetState("DOA")
		
		fmt.Println("Name: " + string(name)) //error
		fmt.Println("Date of Appointment: " + string(doa))
		return nil, nil;
	}
	fmt.Println("query did not find func: " + function)//error

	return nil, errors.New("Received unknown function query: " + function)
}
