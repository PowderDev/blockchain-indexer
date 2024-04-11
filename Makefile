
generate:
	@abigen --abi abi/MakNFT.json --pkg abi_gen --type MakNFT --out abi_gen/MakNFT.go
	@abigen --abi abi/MSCEngine.json --pkg abi_gen --type MSCEngine --out abi_gen/MSCEngine.go