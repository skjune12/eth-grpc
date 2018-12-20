pragma solidity ^0.5.1;

contract ExampleContract{
  uint32 public number;

  function setNumber(uint32 _number) public {
    number = _number;
  }
}
