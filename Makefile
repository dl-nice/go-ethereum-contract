
.PHONY: abi
abi:
	cd contract && solc --abi *.sol -o ./abi && solc --bin *.sol -o ./abi

.PHONY: erc20
erc20:
	cd contract && mkdir "erc20" && abigen --bin=./abi/ERC20Token.bin --abi=./abi/ERC20Token.abi --pkg=erc20 --out=../contract/erc20/erc20.go

.PHONY: erc721
erc721:
	cd contract && mkdir "erc721" &&  abigen --bin=./abi/ERC721Token.bin --abi=./abi/ERC721Token.abi --pkg=erc721 --out=../contract/erc721/erc721.go
