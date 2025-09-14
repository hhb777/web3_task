require("@nomicfoundation/hardhat-toolbox");
// require("hardhat-deploy");
// require("dotenv").config({ debug:true });
require("@openzeppelin/hardhat-upgrades")

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
      localhost: {
        url: "http://127.0.0.1:8545",

      },
      // sepolia: {
      //   url:`https://sepolia.infura.io/v3/xxx`,
      //   accounts: [`xxxx`]
      // }
  }
};
