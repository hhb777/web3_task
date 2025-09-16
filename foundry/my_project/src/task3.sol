// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

contract Calculator {
    function add(uint256 a, uint256 b) public pure returns (uint256 result) {
        // return a + b; // 971 gas
        
        // unchecked {
        //     return a + b; // 792 gas
        // }

        assembly {
            result := add(a, b) // 792 gas
        }
    }

    function subtract(uint256 a, uint256 b) public pure returns (uint256 result) {
        require(b <= a, "Subtraction underflow");
        // return a - b; // 880 gas

        // unchecked {
        //     return a - b; // 796 gas
        // }

        assembly {
            result := sub(a, b) // 796 gas
        }
    }

    function multiply(uint256 a, uint256 b) public pure returns (uint256 result) {
        // return a * b; // 990 gas

        // unchecked {
        //     return a * b; // 750 gas
        // }

        assembly {
            result := mul(a, b) // 750 gas
        }
    }

    function divide(uint256 a, uint256 b) public pure returns (uint256 result) {
        require(b > 0, "Division by zero");
        // return a / b; // 921 gas

        // unchecked {
        //     return a / b; // 856 gas
        // }

        assembly {
            result := div(a, b) // 839 gas
        }
        
    }
}