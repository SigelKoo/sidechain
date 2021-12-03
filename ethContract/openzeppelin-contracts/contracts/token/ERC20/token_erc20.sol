// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ERC20.sol";

contract token_erc20 is ERC20 {
    constructor(uint256 initialSupply) ERC20("token_erc20", "token_erc20") {
        _mint(msg.sender, initialSupply);
    }
}
