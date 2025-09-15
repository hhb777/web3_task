const { ethers, deployments, getNamedAccounts } = require("hardhat");
const { expect } = require("chai");

describe("MyMeme", async function (){
    it("Should be ok", async function (){
        await main();
    });
})

async function main(){

        // 获取签名者，部署合约的账户地址
        const [siger, user1] = await ethers.getSigners();
        console.log("siger", siger.address);
        const Contract = await ethers.getContractFactory("MyMeme");
        const contract = await Contract.connect(user1).deploy("MyMeme", "MEME");
        await contract.waitForDeployment();
        console.log("合约地址：", await contract.getAddress());
        // 铸造一些代币到部署合约的账户1
        await contract.mint(user1.address, ethers.parseEther("100"));

        const user1blance = await contract.balanceOf(user1.address);
        console.log("user1blance", user1blance.toString());

        const Contract2 = await ethers.getContractFactory("MyMeme");
        const contract2 = await Contract2.connect(user1).deploy("MyMeme2", "MEME2");
        await contract2.waitForDeployment();
        console.log("合约地址：", await contract2.getAddress());
        // 铸造一些代币到部署合约的账户1
        await contract2.mint(user1.address, ethers.parseEther("200"));

        const user2blance = await contract2.balanceOf(user1.address);
        console.log("user2blance", user2blance.toString());

        //部署流动性合约
        const LiquidityManager = await ethers.getContractFactory("LiquidityManager");
        const liquidityManager = await LiquidityManager.deploy(contract.getAddress(), contract2.getAddress());
        await liquidityManager.waitForDeployment();
        console.log("流动性合约地址：", await liquidityManager.getAddress());

        //设置tick
        await liquidityManager.addTick(ethers.parseEther("1"), ethers.parseEther("1"));
        await liquidityManager.addTick(ethers.parseEther("2"), ethers.parseEther("2"));

        //更新tick
        await liquidityManager.updateTick(1, ethers.parseEther("3"));
    
        //授权流动性合约可以操作用户的代币
        await contract.connect(user1).approve(liquidityManager.getAddress(), ethers.parseEther("100"));
        await contract2.connect(user1).approve(liquidityManager.getAddress(), ethers.parseEther("200"));
        //添加流动性
        await liquidityManager.connect(user1).addLiquidity(ethers.parseEther("50"), ethers.parseEther("100"));
        const lmblance1 = await contract.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币1余额", lmblance1.toString());
        const lmblance2 = await contract2.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币2余额", lmblance2.toString());
        expect(lmblance1).to.equal(ethers.parseEther("50"));
        expect(lmblance2).to.equal(ethers.parseEther("100"));

        //user1 tokenA => user1 tokenB 代币交易
        await liquidityManager.connect(user1).swap(ethers.parseEther("10"), true);
        const user1blance3 = await contract.balanceOf(user1.address);
        console.log("user1blance", user1blance3.toString());
        expect(user1blance3).to.equal(ethers.parseEther("40"));
        const user2blance3 = await contract2.balanceOf(user1.address);
        console.log("user2blance", user2blance3.toString());
        //3%的手续费
        expect(user2blance3).to.equal(ethers.parseEther("109.7"));

        //user1 tokenB => user1 tokenA 代币交易
        await liquidityManager.connect(user1).swap(ethers.parseEther("20"), false);
        const user1blance4 = await contract.balanceOf(user1.address);
        console.log("user1blance", user1blance4.toString());
        expect(user1blance4).to.equal(ethers.parseEther("59.4"));
        const user2blance4 = await contract2.balanceOf(user1.address);
        console.log("user2blance", user2blance4.toString());
        expect(user2blance4).to.equal(ethers.parseEther("89.7"));

        //移除流动性
        const tuser1blance = await contract.balanceOf(user1.address);
        console.log("user1blance", tuser1blance.toString());
        const tuser2blance = await contract2.balanceOf(user1.address);
        console.log("user1blance", tuser2blance.toString());
        const tlmblance1 = await contract.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币1余额", tlmblance1.toString());
        const tlmblance2 = await contract2.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币2余额", tlmblance2.toString());
        await liquidityManager.connect(user1).removeLiquidity(ethers.parseEther("40"), ethers.parseEther("100"));
        const lmblance11 = await contract.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币1余额", lmblance11.toString());
        const lmblance22 = await contract2.balanceOf(liquidityManager.getAddress());
        console.log("流动性合约代币2余额", lmblance22.toString());
        // 奖励3%的手续费余额
        expect(lmblance11).to.equal(ethers.parseEther("0.3"));
        expect(lmblance22).to.equal(ethers.parseEther("9.7"));

        const user1blance2 = await contract.balanceOf(user1.address);
        console.log("user1blance", user1blance2.toString());
        expect(user1blance2).to.equal(ethers.parseEther("99.4"));
        const user2blance2 = await contract2.balanceOf(user1.address);
        console.log("user2blance", user2blance2.toString());
        expect(user2blance2).to.equal(ethers.parseEther("189.7"));

}

