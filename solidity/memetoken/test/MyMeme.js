const { ethers, deployments, getNamedAccounts } = require("hardhat");
const { expect } = require("chai");

describe("MyMeme", async function (){
    it("Should be ok", async function (){
        await main();
    });
})

async function main(){

        const { user1 } = await getNamedAccounts();
        const { user2 } = await getNamedAccounts();
        // 获取签名者，部署合约的账户地址
        const [siger] = await ethers.getSigners();
        console.log("siger", siger.address);
        // 在测试环境中重置并重新部署合约
        await deployments.fixture(["deployMyMeme"]);
        // 获取已经部署的合约
        const contractobj = await deployments.get("MyMeme");
        // 通过工厂和地址获取合约对象
        const contract = await ethers.getContractAt(
            "MyMeme",
            contractobj.address
        );
        // 铸造一些代币到部署合约的账户
        await contract.mint(siger.address, 100_000_000_000_000);

        const onwerblance = await contract.balanceOf(siger.address);
        console.log("onwerblance", onwerblance.toString());

        // 设置税额池地址
        await contract.setTaxPool(user1);
        // 设置最大交易额
        await contract.setMaxTxAmount(1000_000_000_000_000);
        // 设置每日交易限额
        await contract.setDailyTxLimit(1000);

        // const taxTier = await contract.getTaxTiers();
        // console.log("taxTier", taxTier);
        // 设置税率档位
        await contract.addTaxTier(10000, 1); // 0.01%
        await contract.addTaxTier(50000, 2); // 0.02%
        await contract.addTaxTier(100000, 3); // 0.3%
        await contract.addTaxTier(500000, 5); // 0.05%

        // 每种档位税额测试
        await contract.metransfer(siger.address, user2, 20000);
        onwerblance2 = await contract.balanceOf(siger.address);
        console.log("onwerblance2", onwerblance2.toString());
        reciverblance = await contract.balanceOf(user2);
        console.log("reciverblance", reciverblance.toString());
        taxblance = await contract.balanceOf(user1);
        console.log("taxblance", taxblance.toString());
        expect(taxblance).to.equal(2);

        await contract.metransfer(siger.address, user2, 60000);
        onwerblance2 = await contract.balanceOf(siger.address);
        console.log("onwerblance2", onwerblance2.toString());
        reciverblance = await contract.balanceOf(user2);
        console.log("reciverblance", reciverblance.toString());
        taxblance = await contract.balanceOf(user1);
        console.log("taxblance", taxblance.toString());

        await contract.metransfer(siger.address, user2, 200000);
        onwerblance2 = await contract.balanceOf(siger.address);
        console.log("onwerblance2", onwerblance2.toString());
        reciverblance = await contract.balanceOf(user2);
        console.log("reciverblance", reciverblance.toString());
        taxblance = await contract.balanceOf(user1);
        console.log("taxblance", taxblance.toString());

        await contract.metransfer(siger.address, user2, 1000000);
        onwerblance2 = await contract.balanceOf(siger.address);
        console.log("onwerblance2", onwerblance2.toString());
        reciverblance = await contract.balanceOf(user2);
        console.log("reciverblance", reciverblance.toString());
        taxblance = await contract.balanceOf(user1);
        console.log("taxblance", taxblance.toString());


}

