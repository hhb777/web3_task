const hre = require("hardhat");
const { expect } = require("chai");

describe("ERC20test2", async() => {

    const { ethers } = hre;
    const initialSupply = 10000
    let mycontract;
    let account1,account2;
    beforeEach(async() => {
        // console.log("wait 2s");
        // await new Promise((reslove) => {
        //     setTimeout(() => {
        //         reslove();
        //     }, 2000);
        // })
        [account1, account2] = await ethers.getSigners();
        console.log("=====accounts=======:", account1, account2);
        const ERC20test2 = await ethers.getContractFactory("ERC20test2");

        mycontract = await ERC20test2.connect(account2).deploy(initialSupply)
        mycontract.waitForDeployment();
        const contractaddress = await mycontract.getAddress()
        expect(contractaddress).to.length.greaterThan(0);
        // console.log(contractaddress, "==contract address==")
    })
    it("验证合约的name/symbol/decimal", async() => {
        const name = await mycontract.name();
        const symbol = await mycontract.symbol();
        const decimal = await mycontract.decimals();
        expect(name).to.equal("TEST2");
        expect(symbol).to.equal("TST2");
        expect(decimal).to.equal(18);
        console.log(name, symbol, decimal);
        // console.log("I am test1");
    })
    it("测试转账", async() => {
        // console.log("I am test2");
        // const balanceofaccount1 = await mycontract.balanceOf(account1);
        // expect(balanceofaccount1).to.equal(initialSupply);
        const resp = await mycontract.transfer(account1.address, initialSupply / 2)
        console.log("transfer resp:", resp);
        const balanceofaccount1 = await mycontract.balanceOf(account1.address);
        expect(balanceofaccount1).to.equal(initialSupply / 2);
        const balanceofaccount2 = await mycontract.balanceOf(account2.address);
        expect(balanceofaccount2).to.equal(initialSupply / 2);
    })
})