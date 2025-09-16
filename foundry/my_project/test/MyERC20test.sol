// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

import "forge-std/Test.sol";
import "../src/MyERC20.sol";

contract MyERC20Test is Test {
    MyERC20 token;
    address user = vm.addr(1);

    event Transfer(address indexed from, address indexed to, uint256 value);

    function setUp() public {
        token = new MyERC20("MyToken", "MTK");
    }

    function testMint() public {
        token.mint(user, 100 ether);
        assertEq(token.balanceOf(user), 100 ether);
        // vm.expectEmit(true, true, true, true);
        emit Transfer(address(0), user, 100 ether);
    }  
}