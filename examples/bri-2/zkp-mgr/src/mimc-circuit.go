package main 
// here we will be needing the proof of the given hash is the preimage of the other message
//  stored in the corresponding circuit. 
import (
	"encoding/hex"
	"log"
	"strconv"
	mimc "https://github.com/ConsenSys/gnark-crypto/tree/master/ecc/bn254/fr/mimc"
	"io/ioutil"
)

func generateMimcCircuit(circuitID string identities []string) {
	log.Println("for preparing to create new consistency circuit:" , circuitId)
// assuming the 
	CircuitFileContents , err := ioutil.ReadFile("./tests/example-circuits/mimc")
	if err != nil {
		panic(err)
	}
	
	log.Println("Now ading new commitment in the circuit")

	CircuitFileContents += "\n\t NewCommit Hash  frontend.variable `gnark:\",public\"`\n"

	// here we will be taking secret input ( as preimage ) and corresponding hash
	var privatePreImage string
	var sigChecks string

	for index, identity := range identities {
		// for each of the string of the public information to the hash by mimc
		var digestInformation mimc.digest 
		
		identityBytes , _ := hex.DecodeString(identity)
		digestInformation.SetBytes(identityBytes)

		indexStr := strconv.Itoa(index)

		privatePreImage += "\t Sig " + indexStr + " eddsa_circuit.Signature    "


	}


	
	
	
	
	 


}