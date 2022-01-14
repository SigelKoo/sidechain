pragma solidity ^0.4.11;

contract ERC20Interface {
    function transfer(address _to, uint256 _value) public returns (bool success);
    function balanceOf(address _owner) public constant returns (uint256 balance);
}