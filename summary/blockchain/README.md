
# 基础工具安装&使用

[安装 Solidity 编译器](https://docs.soliditylang.org/en/latest/installing-solidity.html)

安装 abigen 工具用于根据 ABI（Application Binary Interface）文件生成 Go 文件

```shell
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

如生成 [ERC20 代币](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol) 的 Go 合约文件：

```shell
# 由于该合约文件中包含了对其他 sol 文件的引用，所以需要在 openzeppelin-contracts 仓库的根目录下执行
# 根据智能合约 ERC20.sol 生成 ABI 文件，输出文件的目录为 build 
solc --abi contracts/token/ERC20/ERC20.sol -o build
# 在 ERC20.abi 所在目录下利用 abigen 工具，根据 ABI 文件编译生成 Go 合约文件
abigen --abi=./ERC20.abi --pkg=erc20 --out=ERC20.go
```

如果需要通过 Go 代码来进行合约的部署，则除了 ABI 文件外，还需要生成 bin 文件，而在生成 Go 文件时还要引入 bin 文件，这样 Go 文件中才会生成合约部署方法（Deploy 前缀）

```shell
# 根据只能合约 SimpleAuction.sol 生成 BIN 文件，输出文件的目录为 .
solc --bin ./SimpleAuction.sol -o . 
# 利用 abigen 工具，根据 ABI 文件、bin 文件编译生成 Go 合约文件，指定 Go 文件的包名、结构体名称、位置及文件名
# 由于同时引入了 BIN 文件，所以 Go 合约文件中会包含合约的部署方法（方法名以 Deploy 作为前缀）
abigen --abi=./SimpleAuction.abi --bin=./SimpleAuction.bin --pkg=contract --type=simpleAuction --out=../simpleAuction.go
```

