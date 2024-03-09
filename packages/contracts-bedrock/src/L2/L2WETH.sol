// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "src/libraries/Predeploys.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";


/// @custom:proxied
/// @title L2WETH
/// @notice L2WETH holds ERC20 balance for ETH crossed over from L1
contract L2WETH is OptimismMintableERC20 {
    /// @notice Initializes the contract as an Optimism Mintable ERC20.
    constructor() OptimismMintableERC20(Predeploys.L2_STANDARD_BRIDGE, address(0), "Wrapped Ether", "WETH", 18) { }
}