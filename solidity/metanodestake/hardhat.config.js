require("@nomicfoundation/hardhat-toolbox");
// require("hardhat-deploy");
// require("dotenv").config({ debug:true });
require("@openzeppelin/hardhat-upgrades")

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
      sepolia: {
        url:`https://sepolia.infura.io/v3/xxx`,
        accounts: [`xxxx`]
      }
  }
};
