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
	var dob Time // date of birth and date of appointment
	var err error	

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	// Initialize IDs
	id,err = strconv.Atoi(args[0])
	name = args[1]
	dob,err = Parse(args[2])	

	// Put data into world stage
	err = stub.PutState("Id", []byte(id))
	if err != nil {
		return nil,err
	}

	err = stub.PutState("Name", []byte(name))
	if err != nil {
		return nil,err
	}
	
	err = stub.PutState("DOB", []byte(name))
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

	if function == "book_appointment" { //book an appointment
		var doa Time
		doa = Parse(args[0])
		stub.putState("DOA",[]byte(doa)) // date of appointment
	}

	if function == "cancel_appointment" { // cancel an appointment
		var doa Time
		doa = Parse(args[0])
		stub.putState("DOA",[]byte(doa))
	}
	
	fmt.Println("invoke did not find func: " + function)//error

	return nil, errors.New("Received unknown function invocation: " + function)
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" {//read a variable
		fmt.Println("hi there " + function)//error
		return nil, nil;
	}
	fmt.Println("query did not find func: " + function)//error

	return nil, errors.New("Received unknown function query: " + function)
}




