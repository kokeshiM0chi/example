package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Circuit defines a simple circuit
// x**3 + x + 5 == y
type Circuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

type basicCircuit struct {
	X frontend.Variable
}

func (c *basicCircuit) Define(api frontend.API) error {
	cmt, err := api.(frontend.Committer).Commit(c.X)
	if err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	api.AssertIsEqual(cmt, "0xaabbcc")
	return nil
}

type constantHash struct{}

func (h constantHash) Write(p []byte) (n int, err error) { return len(p), nil }
func (h constantHash) Sum(b []byte) []byte               { return []byte{0xaa, 0xbb, 0xcc} }
func (h constantHash) Reset()                            {}
func (h constantHash) Size() int                         { return 3 }
func (h constantHash) BlockSize() int                    { return 32 }

func main() {
	{
		curve := ecc.BN254
		assignment := &basicCircuit{X: 1}
		ccs, _ := frontend.Compile(curve.ScalarField(), r1cs.NewBuilder, &basicCircuit{})
		pk, vk, _ := groth16.Setup(ccs)
		witness, _ := frontend.NewWitness(assignment, curve.ScalarField())
		proof, _ := groth16.Prove(ccs, pk, witness, backend.WithProverHashToFieldFunction(constantHash{}))
		pubWitness, _ := witness.Public()
		_ = groth16.Verify(proof, vk, pubWitness, backend.WithVerifierHashToFieldFunction(constantHash{}))
	}

	{
		curve := ecc.BLS12_381
		assignment := &basicCircuit{X: 1}
		ccs, _ := frontend.Compile(curve.ScalarField(), r1cs.NewBuilder, &basicCircuit{})
		pk, vk, _ := groth16.Setup(ccs)
		witness, _ := frontend.NewWitness(assignment, curve.ScalarField())
		proof, _ := groth16.Prove(ccs, pk, witness, backend.WithProverHashToFieldFunction(constantHash{}))
		pubWitness, _ := witness.Public()
		_ = groth16.Verify(proof, vk, pubWitness, backend.WithVerifierHashToFieldFunction(constantHash{}))
	}

}
