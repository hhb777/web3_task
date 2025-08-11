// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Car{
    uint speed;
    constructor(uint _speed){
        speed = _speed;
    }
    function drive() public virtual returns(uint,uint){}
}

contract ElectricCar is Car{
    uint batterLevel;
    constructor() Car(10) {
        batterLevel = 2;
    }
    function drive() public view override returns(uint, uint){
        return (batterLevel, Car.speed);
    }
}

contract Person{
    function test() public pure virtual returns(string memory){
        return "Person";
    }
}

contract Employee{
    function test() public pure virtual returns(string memory){
        return "Employee";
    }
}

contract Manager is Person,Employee{
    function test() public pure override (Person,Employee) returns(string memory){
        return "Manager";
    }
}

abstract contract Shape{
    function area(uint,uint) virtual pure public returns(uint);
}

contract Square is Shape{
    function area(uint x,uint y) public pure override returns(uint){
        return x*y;
    }
}

contract Circle is Shape{
    function area(uint x,uint y) public pure override returns(uint){
        return 3*x;
    }
}