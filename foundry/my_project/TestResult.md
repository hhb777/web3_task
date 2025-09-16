# 一、部署环境：

## 1.初始化项目：
forge init my_project
cd my_project

## 2.安装合约需要的依赖
forge install OpenZeppelin/openzeppelin-contracts

## 3.写完合约执行build编译
forge build src/task3.sol

## 4.执行测试代码
forge test test/task3test.sol

## 5.gas记录
forge test test/task3test.sol --gas-report

# 二、测试结果展示：
## 1.第一次测试
<img width="1065" height="570" alt="image" src="https://github.com/user-attachments/assets/dba879c9-5da3-4d14-9d40-8b23179342d7" />

## 2.第二次添加unchecked优化测试
<img width="990" height="671" alt="image" src="https://github.com/user-attachments/assets/6a082717-c0fd-49bb-94c5-6d748a55d260" />

## 3.第三次内联汇编优化测试
<img width="995" height="710" alt="image" src="https://github.com/user-attachments/assets/d5357e83-f2ff-47df-a441-c4555f9d17fc" />

