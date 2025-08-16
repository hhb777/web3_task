const { ethers, deployments } = require("hardhat")
const { expect } = require("chai")


describe("Test auction", async function () {
    it("Should be ok", async function () {
        await main();
    });
})

async function main() {
    const [signer, buyer] = await ethers.getSigners()
    await deployments.fixture(["depolyNftAuction"]);
    
    const nftAuctionProxy = await deployments.get("NftAuctionProxy");
    const nftAuction = await ethers.getContractAt(
        "Auction",
        nftAuctionProxy.address
    );

    const TestERC20 = await ethers.getContractFactory("TestERC20");
    const testERC20 = await TestERC20.deploy();
    await testERC20.waitForDeployment();
    const UsdcAddress = await testERC20.getAddress();
    
    let tx = await testERC20.connect(signer).transfer(buyer, ethers.parseEther("100"))
    await tx.wait()
    // nftAuctionProxy.setPriceFeed()
    const aggreagatorV3 = await ethers.getContractFactory("AggreagatorV3")
    const priceFeedEthDeploy = await aggreagatorV3.deploy(ethers.parseEther("1000"))
    const priceFeedEth = await priceFeedEthDeploy.waitForDeployment()
    const priceFeedEthAddress = await priceFeedEth.getAddress()
    console.log("ethFeed: ", priceFeedEthAddress)
    const priceFeedUSDCDeploy = await aggreagatorV3.deploy(ethers.parseEther("1"))
    const priceFeedUSDC = await priceFeedUSDCDeploy.waitForDeployment()
    const priceFeedUSDCAddress = await priceFeedUSDC.getAddress()
    console.log("usdcFeed: ", await priceFeedUSDCAddress)

    const token2Usd = [{
        token: ethers.ZeroAddress,
        priceFeed: priceFeedEthAddress
    }, {
        token: UsdcAddress,
        priceFeed: priceFeedUSDCAddress
    }]

    for (let i = 0; i < token2Usd.length; i++) {
        const { token, priceFeed } = token2Usd[i];
        await nftAuction.setPriceFeed(token, priceFeed);
    }
    // 1. 部署 ERC721 合约
    const TestERC721 = await ethers.getContractFactory("MyNFT");
    const testERC721 = await TestERC721.deploy();
    await testERC721.waitForDeployment();
    const testERC721Address = await testERC721.getAddress();
    console.log("testERC721Address::", testERC721Address);

    // mint 10个 NFT
    for (let i = 0; i < 10; i++) {
        await testERC721.mintNFT(signer.address, i + 1);
    }

    const tokenId = 1;    

    console.log("mintNFT success::", tokenId);
    // 给代理合约授权
    await testERC721.connect(signer).setApprovalForAll(nftAuctionProxy.address, true);
    console.log("=============setapproval success::");

    await nftAuction.createAuction(
        testERC721Address,
        tokenId,
        ethers.parseEther("0.01"),
        10,  
    );
    console.log("=============createAuction success::");
    const auction = await nftAuction.bids(0);

    console.log("创建拍卖成功：：", auction);

    // 3. 购买者参与拍卖
    // await testERC721.connect(buyer).approve(nftAuctionProxy.address, tokenId);
    // const res = await nftAuction.getLastPrice(ethers.ZeroAddress);
    // console.log("getLastPrice ETH success::", res);
    // ETH参与竞价
    tx = await nftAuction.connect(buyer).startbid(0, 0, ethers.ZeroAddress, { value: ethers.parseEther("0.1") });
    await tx.wait()
    console.log("ETH success::");
    // USDC参与竞价
    tx = await testERC20.connect(buyer).approve(nftAuctionProxy.address, ethers.MaxUint256)
    await tx.wait()
    console.log("approve success::");
    tx = await nftAuction.connect(buyer).startbid(0, ethers.parseEther("10"), UsdcAddress);
    await tx.wait()
    console.log("USDC success::");

    // 4. 结束拍卖
    // 等待 10 s
    await new Promise((resolve) => setTimeout(resolve, 10 * 1000));

    await nftAuction.connect(signer).endbid(0);

    // 验证结果
    const auctionResult = await nftAuction.bids(0);
    console.log("结束拍卖后读取拍卖成功：：", auctionResult);
    expect(auctionResult.highestBidder).to.equal(buyer.address);
    expect(auctionResult.highestBid).to.equal(ethers.parseEther("10"));

    // 验证 NFT 所有权
    const owner = await testERC721.ownerOf(tokenId);
    console.log("owner::", owner);
    expect(owner).to.equal(buyer.address);
}

main()