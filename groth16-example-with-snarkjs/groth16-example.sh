#!/bin/bash
function hello () {
    echo hello
}



## generate r1cs circuit
circom multiplier2.circom --r1cs --wasm --sym --c

## calculate witness
cp input.json multiplier2_js/input.json
cd multiplier2_js && node generate_witness.js multiplier2.wasm input.json witness.wtns

# trusted setup(2 phase setup)
## phase1: power of tau

snarkjs powersoftau new bn128 12 pot12_0000.ptau -v
snarkjs powersoftau contribute pot12_0000.ptau pot12_0001.ptau --name="First contribution" -v

## phase2: specified circuit
snarkjs powersoftau prepare phase2 pot12_0001.ptau pot12_final.ptau -v
snarkjs groth16 setup ../multiplier2.r1cs pot12_final.ptau multiplier2_0000.zkey

# gen key
snarkjs zkey contribute multiplier2_0000.zkey multiplier2_0001.zkey --name="1st Contributor Name" -v

# export key
snarkjs zkey export verificationkey multiplier2_0001.zkey verification_key.json

# generate proof
snarkjs groth16 prove multiplier2_0001.zkey witness.wtns proof.json public.json

# verify proof
snarkjs groth16 verify verification_key.json public.json proof.json