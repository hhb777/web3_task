// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;


import "forge-std/Test.sol";
import "../src/task3.sol";

contract Task3Test is Test{
    Calculator calculator;

    function setUp() public {
        calculator = new Calculator();
    }

    function testAdd() public view {
        assertEq(calculator.add(2, 3), 5);
        assertEq(calculator.add(0, 0), 0);
        assertEq(calculator.add(1e18, 2e18), 3e18);
    }

    function testSubtract() public view {
        assertEq(calculator.subtract(5, 3), 2);
        assertEq(calculator.subtract(3, 3), 0);
    }

    function testSubtractUnderflow() public {
        vm.expectRevert("Subtraction underflow");
        calculator.subtract(3, 5);
    }

    function testMultiply() public view {
        assertEq(calculator.multiply(2, 3), 6);
        assertEq(calculator.multiply(0, 100), 0);
        assertEq(calculator.multiply(1e9, 1e9), 1e18);
    }

    function testDivide() public view {
        assertEq(calculator.divide(6, 3), 2);
        assertEq(calculator.divide(5, 2), 2); // Integer division
    }

    function testDivideByZero() public {
        vm.expectRevert("Division by zero");
        calculator.divide(5, 0);
    }
}