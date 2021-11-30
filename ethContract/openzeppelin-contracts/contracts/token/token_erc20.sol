pragma solidity ^0.8.0;

import "./ERC20/ERC20.sol";

contract token_erc20 is ERC20 {
    constructor(uint256 initialSupply) public ERC20("token_erc20", "token_erc20") {
        _mint(msg.sender, initialSupply);
    }
}
