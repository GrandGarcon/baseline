import ("fmt" , 
"github.com/consensys/gnark/backend/groth16",
p256  "github.com/ing-bank/zkrp/crypto", 
eddsa_gen "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa",

)


type ExampleCircuit struct {
//  here i will be taking example of an real use case  
 frontend.Variable `zkrp:",public`
 Pubkey eddsa_circuit.Signature
 SecretKey ed


}



func 
//fetchiing the contents of the proofs and id , in order to generate 



// now here we will have to utilise other techniques , like plonk 