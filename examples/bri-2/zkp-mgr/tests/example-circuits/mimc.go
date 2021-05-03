package mimc
// reference : https://github.com/ConsenSys/gnark/tree/master/examples.
import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)
type Circuit struct {
	// struct tag on a variable is optional
	// default uses variable name and secret visibility.
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

func (circuit *Circuit) Define(curveID ecc.ID , cs *frontend.ConstraintSystem ) error {
	mimc, _ := mimc.NewMiMC("seed", curveID)
	// specify constraints
	// mimc(preImage) == hash
	cs.AssertIsEqual(circuit.Hash, mimc.Hash(cs, circuit.PreImage))
	return nil
}
