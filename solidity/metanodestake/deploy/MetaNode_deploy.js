const { ethers,deployments,upgrades } = require("hardhat");

// module.exports = async ({ deployments }) => {
async function main() {
    const [siger] = await ethers.getSigners();
    const { save } = deployments;
    
    //先部署token合约
    const tokenContract = await ethers.getContractFactory("MetaNodeToken");
    const tokencontract = await tokenContract.deploy();
    await tokencontract.waitForDeployment();
    const tokenaddress = await tokencontract.getAddress()
    console.log("部署MetaNodeToken合约地址:", tokenaddress);


    //部署stake合约
    const stakeContract = await ethers.getContractFactory("MetaNodeStake");
     // 质押起始区块高度,可以去sepolia上面读取最新的区块高度
    const startBlock = await ethers.provider.getBlockNumber();
    // 质押结束的区块高度,sepolia 出块时间是12s,想要质押合约运行x秒,那么endBlock = startBlock+x/12
    const endBlok = startBlock + 300000;
    // 每个区块奖励的MetaNode token的数量
    const MetaNodePerBlock = "2000000000000000";
    const stakecontract = await upgrades.deployProxy(
        stakeContract,
        [tokenaddress, startBlock, endBlok, MetaNodePerBlock],
        { initializer: "initialize" }
    );
    await stakecontract.waitForDeployment();
    console.log("部署MetaNodeStake合约地址:", await stakecontract.getAddress());

    const tokenAmount = await tokencontract.balanceOf(siger.address)
    //质押所有MetaNodeToken代币
    let tx = await tokencontract.connect(siger).transfer(stakecontract.getAddress(), tokenAmount)
    await tx.wait()
    
    const tokenArtifact = await deployments.getArtifact("MetaNodeToken");
    await save("MetaNodeToken", {
        abi: tokenArtifact,
        address: tokenaddress,
    })

    const stakeArtifact = await deployments.getArtifact("MetaNodeStake");
    await save("MetaNodeStake", {
        abi: stakeArtifact,
        address: await stakecontract.getAddress(),
    })

    console.log("tokenaddress:", tokenaddress);
    console.log("stakeaddress:", await stakecontract.getAddress());
};

// module.exports.tag = ["deployMetaNode"];
main();