1. 初始化模块
go mod init task1
2. 安装ethclient依赖
go get github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/rpc
go get github.com/ethereum/go-ethereum/keystore
3. solcjs工具将合约生成abi文件 
..\..\solidity\hardhat\node_modules\.bin\solcjs --abi ..\mycontracts\countertest.sol
4. solcjs工具将合约生成bin文件
..\..\solidity\hardhat\node_modules\.bin\solcjs --bin ..\mycontracts\countertest.sol
5. abigen工具生成go文件
   D:\goproject\pkg\mod\github.com\ethereum\go-ethereum@v1.16.2\build\bin\cmd\abigen.exe --bin=countertest_sol_CounterTest.bin --abi=countertest_sol_CounterTest.abi --pkg=countertest --out=countertest.go
6.安装keystore
  go get github.com/ethereum/go-ethereum/accounts/keystore@v1.16.2
