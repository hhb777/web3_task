const hre = require("hardhat");
const { expect } = require("chai");

describe("ERC721test2", async() => {
    const {ethers} = hre;
    let mycontract;
    beforeEach(async() => {
        const ERC721test2 = await ethers.getContractFactory("Mytoken");
        mycontract = await ERC721test2.deploy("TEST721", "TST721");
        mycontract.waitForDeployment();
    })
    it("验证合约的name/symbol", async() => {
        const name = await mycontract.name();
        const symbol = await mycontract.symbol();
        expect(name).to.equal("TEST721");
        expect(symbol).to.equal("TST721");
    })
    it("测试mint", async() => {
        const [account1, account2] = await ethers.getSigners();
        console.log("=====accounts=======:", account1, account2);
        const resp = await mycontract.connect(account1).mintNFT(account1.address, "https://example.com/metadata");
        console.log("mint resp:", resp);
        const balanceofaccount1 = await mycontract.balanceOf(account1.address);
        expect(balanceofaccount1).to.equal(1);
        const ownerofnft1 = await mycontract.ownerOf(1);
        expect(ownerofnft1).to.equal(account1.address);
        const tokenuri1 = await mycontract.tokenURI(1);
        expect(tokenuri1).to.equal("https://example.com/metadata");
        console.log("tokenuri1:", tokenuri1);
    })

})