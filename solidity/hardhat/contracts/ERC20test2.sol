// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ERC20test2 is ERC20 {
    constructor(uint256 initialsupply) ERC20("TEST2", "TST2") {
        _mint(msg.sender, initialsupply);
    }

}