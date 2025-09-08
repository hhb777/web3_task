一、部署hardhat环境
npx hardhat init
在README.md文件中先启动npx hardhat node启动本地节点服务

如果没有deploy目录就创建一个deploy目录,然后执行npx hardhat deploy

vscode中执行命令npx  remixd即可连接remixd（若没有安装remixd执行 npm install remixd命令）

二、安装项目需要的依赖
//安装uniswap依赖
npm install @uniswap/v2-core @uniswap/v2-periphery
//安装openzeppelin依赖
npm install @openzeppelin/contracts

三、编写MyMeme.sol合约和部署
编写deploy_mymeme.js来部署合约
MyMeme.js编写测试用例
测试结果展示：
<img width="1225" height="644" alt="image" src="https://github.com/user-attachments/assets/1493f158-3e0f-4ad2-8477-5b1c4395839c" />

<img width="1213" height="820" alt="image" src="https://github.com/user-attachments/assets/71655c81-0fe0-440c-b351-cad99f6b7650" />

四、编写LiquidityManager.sol流动性池集成合约和部署
