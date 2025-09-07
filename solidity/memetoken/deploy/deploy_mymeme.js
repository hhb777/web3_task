const { deployments, ethers } = require("hardhat");

module.exports = async ({ getNamedAccounts, deployments }) => {
    const { deployer } = await getNamedAccounts();
    const { save } = deployments;

    console.log("部署用户地址：", deployer);
    const Contract = await ethers.getContractFactory("MyMeme");
    const contract = await Contract.deploy("MyMeme", "MEME");
    await contract.waitForDeployment();
    console.log("合约地址：", await contract.getAddress());

    await save("MyMeme", {
        abi: Contract.interface.format("json"),
        address: await contract.getAddress(),
    })

};
module.exports.tags = ["deployMyMeme"];