package main

import (
	"fmt"
	"github.com/iden3/go-rapidsnark/prover"
	"os"
)

func main() {
	zkeyBytes, err := os.ReadFile("./testdata/prove.zkey")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read zkey file: %v\n", err)
		os.Exit(1)
	}

	wtnsBytes, err := os.ReadFile("./testdata/witness.wtns")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read witness file: %v\n", err)
		os.Exit(1)
	}

	proof, publicInputs, err := prover.Groth16ProverRaw(zkeyBytes, wtnsBytes)
	if err != nil {
		panic(err)
	}

	var proofFName = "proof.json"
	var publicInputsFName = "public.json"

	err = os.WriteFile(proofFName, []byte(proof), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(publicInputsFName, []byte(publicInputs), 0644)
	if err != nil {
		panic(err)
	}
}
