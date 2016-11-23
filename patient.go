package main

import (
	"errors"
	"fmt"
	"time"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	var id int
	var name string // Entities	
	var dob time.Time
	var err error	

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	// Initialize IDs
	const shortForm = "2006-Jan-02"
	id,err = strconv.Atoi(args[0])
	name = args[1]
	dob,err = time.Parse(shortForm,args[2])	

	// Put data into world stage
	err = stub.PutState("Id", []byte(strconv.Itoa(id)))
	if err != nil {
		return nil,err
	}

	err = stub.PutState("Name", []byte(name))
	if err != nil {
		return nil,err
	}
	
	err = stub.PutState("DOB", []byte(dob.String()))
	if err != nil {
		return nil,err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}

	// Read in the name of the appointment to be updated
	if function == "register_with_gp" {
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
		// delete action in Fabric Blockchain
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
	if function == "query_appointment" {//read a variable
		name,_ := stub.GetState("Name")
		doa,_ := stub.GetState("DOA")
		
		fmt.Println("Name: " + string(name)) //error
		fmt.Println("Date of Appointment: " + string(doa))
		return nil, nil;
	}
	fmt.Println("query did not find func: " + function)//error

	return nil, errors.New("Received unknown function query: " + function)
}
