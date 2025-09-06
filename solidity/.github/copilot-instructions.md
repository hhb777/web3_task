# Copilot Instructions for AI Agents

## Project Overview
This codebase is a multi-task Solidity smart contract workspace using Hardhat for development, testing, and deployment. It is organized by tasks and learning modules, with each task or module in its own directory. The main focus is on learning, experimenting, and testing various Solidity patterns and contract types.

## Key Directories & Files
- `hardhat/` – Core Hardhat project setup, config, and sample contracts/tests. Use this as a reference for project structure and Hardhat usage.
- `task1/`, `task2/`, `task3/` – Each contains contracts, tests, and sometimes deployment scripts for specific assignments or experiments.
- `soliditybase/`, `homework/` – Collections of example and practice contracts for learning Solidity features.
- `contracts/` (in each task or in `hardhat/`) – Solidity contract sources.
- `test/` (in each task or in `hardhat/`) – JavaScript tests for contracts, using Hardhat and ethers.js.
- `deploy/` (in some tasks) – Deployment scripts for contracts.

## Developer Workflows
- **Install dependencies:**
  ```shell
  npm install
  ```
- **Run all tests:**
  ```shell
  npx hardhat test
  ```
- **Deploy contracts (example):**
  ```shell
  npx hardhat ignition deploy ./ignition/modules/Lock.js
  ```
- **Start local node:**
  ```shell
  npx hardhat node
  ```
- **Gas reporting:**
  ```shell
  REPORT_GAS=true npx hardhat test
  ```

## Project-Specific Patterns & Conventions
- Each task is self-contained: contracts, tests, and deployment scripts are grouped by task.
- Use Hardhat Ignition for deployment modules (see `ignition/modules/`).
- Test files are in JavaScript and use ethers.js for contract interaction.
- Artifacts and build info are auto-generated in `artifacts/` and `build-info/`.
- Naming: Contract and test files are named to match their purpose (e.g., `ERC20test2.sol`, `auction.js`).
- Some tasks include deployment results or screenshots in their `README.md`.

## Integration & External Dependencies
- Uses OpenZeppelin contracts (see `artifacts/@openzeppelin/` and imports in contracts).
- Chainlink contracts are used in some tasks (see `artifacts/@chainlink/`).
- Ethers.js is used for testing and scripting.

## Examples
- To test the `auction` contract in `task3`, run:
  ```shell
  npx hardhat test test/auction.js
  ```
- To deploy a contract using Ignition:
  ```shell
  npx hardhat ignition deploy ./ignition/modules/Lock.js
  ```

## References
- See each task's `README.md` for task-specific notes, results, or screenshots.
- Use `hardhat.config.js` in each project root for configuration details.

---
_If any conventions or workflows are unclear, please ask for clarification or check the relevant `README.md` files in each task directory._
