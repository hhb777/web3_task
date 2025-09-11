const { ethers,deployments,upgrades } = require("hardhat")
const { expect } = require("chai")

describe("Test MetaNodeStake", async function() {
    it("Should be ok", async function () {
        await main();
    });
})

async function main(){
    
    const [siger] = await ethers.getSigners()
    //先部署token合约
    const MetaNodeTokenContract = await ethers.getContractFactory("MetaNodeToken");
    const MetaNodeToken = await MetaNodeTokenContract.deploy();
    await MetaNodeToken.waitForDeployment();
    const MetaNodeTokenAddress = await MetaNodeToken.getAddress()
    console.log("部署MetaNodeToken合约地址:", MetaNodeTokenAddress);


    //部署stake合约
    const MetaNodeStakeContract = await ethers.getContractFactory("MetaNodeStake");
     // 质押起始区块高度,可以去sepolia上面读取最新的区块高度
    const startBlock = await ethers.provider.getBlockNumber();
    // 质押结束的区块高度,sepolia 出块时间是12s,想要质押合约运行x秒,那么endBlock = startBlock+x/12
    const endBlok = startBlock + 300000;
    // 每个区块奖励的MetaNode token的数量0.02ETH
    const MetaNodePerBlock = "2000000000000000";
    const MetaNodeStake = await upgrades.deployProxy(
        MetaNodeStakeContract,
        [MetaNodeTokenAddress, startBlock, endBlok, MetaNodePerBlock],
        { initializer: "initialize" }
    );
    await MetaNodeStake.waitForDeployment();
    console.log("部署MetaNodeStake合约地址:", await MetaNodeStake.getAddress());

    const tokenAmount = await MetaNodeToken.balanceOf(siger.address)
    //质押所有MetaNodeToken代币
    let tx = await MetaNodeToken.connect(siger).transfer(MetaNodeStake.getAddress(), tokenAmount)
    await tx.wait()
    const lastblocknum = await ethers.provider.getBlockNumber();
 
    // await deployments.fixture(["deployMetaNode"]);

    // // const MetaNodeTokenObj = await deployments.get("MetaNodeToken");
    // const MetaNodeTokenAddress = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788";
    // const MetaNodeStakeAddress = "0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e";
    // const MetaNodeToken = await ethers.getContractAt(
    //     "MetaNodeToken",
    //     MetaNodeTokenAddress
    // );

    // // const MetaNodeStakeObj = await deployments.get("MetaNodeStake");
    // const MetaNodeStake = await ethers.getContractAt(
    //     "MetaNodeStake",
    //     MetaNodeStakeAddress
    // );

    /**
     * 添加池测试
     * 参数一、token合约地址
     * 参数二、池权重
     * 参数三、最小质押额度
     * 参数四、unstake的时候需要等待的块数
     * 参数五、是否更新池
     */
    console.log("start addpool");
    await MetaNodeStake.addPool(ethers.ZeroAddress, 500, 100, 20, true);
    console.log("success for addpool");
    const cpool =  await MetaNodeStake.pool(0);
    expect(cpool.stTokenAddress).to.equal(ethers.ZeroAddress);
    expect(cpool.poolWeight).to.equal(500);
    // expect(cpool.lastRewardBlock).to.equal(startBlock);
    expect(cpool.accMetaNodePerST).to.equal(0);
    expect(cpool.stTokenAmount).to.equal(0);
    expect(cpool.minDepositAmount).to.equal(100);
    expect(cpool.unstakeLockedBlocks).to.equal(20);

    
    /**
     * 更新池测试
     */
    await MetaNodeStake.massUpdatePools();
    const pool = await MetaNodeStake.pool(0);

    expect(pool.minDepositAmount).to.equal(100);
    expect(pool.unstakeLockedBlocks).to.equal(20);
    // expect(pool.lastRewardBlock).to.equal(startBlock);

    
    for (let i = lastblocknum; i < lastblocknum + 1; i++) {
            await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.massUpdatePools();
    const updatedpool = await MetaNodeStake.pool(0);
    // expect(updatedpool.lastRewardBlock).to.equal(1000)

    /**
     * 设置池权重
    */
    const preTotalPoolWeight = await MetaNodeStake.totalPoolWeight();
    console.log("preTotalPoolWeight:", preTotalPoolWeight);
    await MetaNodeStake.setPoolWeight(0, 200, false);
    const pool2 = await MetaNodeStake.pool(0);
    const totalPoolWeight = await MetaNodeStake.totalPoolWeight();
    console.log("totalPoolWeight:",totalPoolWeight);
    expect(pool2.poolWeight).to.equal(200);
    expect(totalPoolWeight).to.equal(preTotalPoolWeight - BigInt(500) + BigInt(200));
    console.log("success setPoolWeight");

    /**
     * 质押native currency
     */
    const pool3 = await MetaNodeStake.pool(0);
    const prePoolStTokenAmount = pool3.stTokenAmount;

    // await siger.sendTransaction({ to: MetaNodeStake.address, value: 1000_000_000 });
    console.log("start depositETH");
    await MetaNodeStake.connect(siger).depositETH({value: 1000_000_000});
    console.log("end depositETH");
    const updatedPool = await MetaNodeStake.pool(0);
    const userInfo = await MetaNodeStake.user(0, siger.address);
    console.log("userInfo:", userInfo);
    expect(updatedPool.stTokenAmount).to.equal(prePoolStTokenAmount + BigInt(1000_000_000 ));

    // await siger.sendTransaction({
    //     to: MetaNodeStake.address,
    //     value: 2000_000_000, 
    // });
    await MetaNodeStake.connect(siger).depositETH({value: 2000_000_000});

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
            await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.unstake(0, 100);

    
    // Repeat for further deposits and unstaking
    await MetaNodeStake.unstake(0, 100);
    for (let i = 3; i <= 7; i++) {
        // await siger.sendTransaction({
        //     to: MetaNodeStake.address,
        //     value: 1000_000_000, 
        // });
        await MetaNodeStake.connect(siger).depositETH({value: 1000_000_000});
        for (let j = lastblocknum; j < lastblocknum + i; j++) {
            await ethers.provider.send("evm_mine", []);
        }
        await MetaNodeStake.unstake(0, 100);
    }

    await MetaNodeStake.withdraw(0);
    console.log("success deposit");

    /**
     * claim after deposit
     */
    // await MetaNodeToken.transfer(MetaNodeStake.getAddress(), 100000000000);
    const preUserMetaNodeBalance = await MetaNodeToken.balanceOf(siger.address);

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
        await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.claim(0);

    const postUserMetaNodeBalance = await MetaNodeToken.balanceOf(siger.address);
    expect(postUserMetaNodeBalance).to.be.gt(preUserMetaNodeBalance);

    /**
     * unstake测试
     */

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
        await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.unstake(0, 100);

    const user = await MetaNodeStake.user(0, siger.address);
    // expect(user.stAmount).to.equal(0);
    // expect(user.finishedMetaNode).to.equal(0);
    console.log("unstake user:", user);
    expect(user.pendingMetaNode).to.be.gt(0);

    const pool4 = await MetaNodeStake.pool(0);
    // expect(pool4.stTokenAmount).to.equal(0);
    console.log("unstake pool:", pool4);

    /**
     * claim after unstake
     */
    // await MetaNodeToken.transfer(MetaNodeStake.getAddress(), 100000000000);
    const preUserMetaNodeBalance2 = await MetaNodeToken.balanceOf(siger.address);

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
        await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.claim(0);

    const postUserMetaNodeBalance2 = await MetaNodeToken.balanceOf(siger.address);
    expect(postUserMetaNodeBalance2).to.be.gt(preUserMetaNodeBalance2);

    /**
     * withdraw测试
     */
    const preContractBalance = await ethers.provider.getBalance(MetaNodeStake.getAddress());
    const preUserBalance = await ethers.provider.getBalance(siger.address);

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
        await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.withdraw(0);

    const postContractBalance = await ethers.provider.getBalance(MetaNodeStake.getAddress());
    const postUserBalance = await ethers.provider.getBalance(siger.address);
    expect(postContractBalance).to.be.lt(preContractBalance);
    // expect(postUserBalance).to.be.gt(preUserBalance);???

    /**
     * claim after withdraw
     */
    // await MetaNodeToken.transfer(MetaNodeStake.getAddress(), 100000000000);
    const preUserMetaNodeBalance3 = await MetaNodeToken.balanceOf(siger.address);

    for (let i = lastblocknum; i < lastblocknum + 2; i++) {
        await ethers.provider.send("evm_mine", []);
    }
    await MetaNodeStake.claim(0);

    const postUserMetaNodeBalance3 = await MetaNodeToken.balanceOf(siger.address);
    expect(postUserMetaNodeBalance3).to.be.gt(preUserMetaNodeBalance3);
}