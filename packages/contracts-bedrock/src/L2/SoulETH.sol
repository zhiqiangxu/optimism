// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {ERC20Upgradeable} from '@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol';
import {OwnableUpgradeable} from '@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol';

contract SoulETH is ERC20Upgradeable, OwnableUpgradeable {
  /// @custom:storage-location erc7201:openzeppelin.storage.SoulETH
  struct SoulETHStorage {
    mapping(address => bool) _minters;
    mapping(address => bool) _burners;
  }

  // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.SoulETH")) - 1)) & ~bytes32(uint256(0xff))
  bytes32 private constant _SOULETH_STORAGE_LOCATION =
    0x6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc00;

  function _getSoulETHStorage()
    private
    pure
    returns (SoulETHStorage storage $)
  {
    assembly {
      $.slot := _SOULETH_STORAGE_LOCATION
    }
  }

  constructor() {
    _disableInitializers();
  }

  /// @custom:legacy
  /// @notice initialize is used to initialize SoulETH contract.
  function initialize(
    string calldata name_,
    string calldata symbol_,
    address owner_,
    address[] calldata minters_,
    address[] calldata burners_
  ) external initializer {
    require(owner_ != address(0), 'owner is zero');

    __ERC20_init(name_, symbol_);
    __Ownable_init();
    transferOwnership(owner_);

    SoulETHStorage storage $ = _getSoulETHStorage();
    uint i;
    for (i = 0; i < minters_.length; i++) {
      $._minters[minters_[i]] = true;
    }
    for (i = 0; i < burners_.length; i++) {
      $._burners[burners_[i]] = true;
    }
  }

  /// @custom:legacy
  /// @notice addMinters is used by owner to add minters.
  function addMinters(address[] calldata minters_) external onlyOwner {
    SoulETHStorage storage $ = _getSoulETHStorage();
    uint i;
    for (i = 0; i < minters_.length; i++) {
      $._minters[minters_[i]] = true;
    }
  }

  /// @custom:legacy
  /// @notice delMinters is used by owner to delete minters.
  function delMinters(address[] calldata minters_) external onlyOwner {
    SoulETHStorage storage $ = _getSoulETHStorage();
    uint i;
    for (i = 0; i < minters_.length; i++) {
      delete $._minters[minters_[i]];
    }
  }

  /// @custom:legacy
  /// @notice addBurners is used by owner to add burners.
  function addBurners(address[] calldata burners_) external onlyOwner {
    SoulETHStorage storage $ = _getSoulETHStorage();
    uint i;
    for (i = 0; i < burners_.length; i++) {
      $._burners[burners_[i]] = true;
    }
  }

  /// @custom:legacy
  /// @notice delBurners is used by owner to delete burners.
  function delBurners(address[] calldata burners_) external onlyOwner {
    SoulETHStorage storage $ = _getSoulETHStorage();
    uint i;
    for (i = 0; i < burners_.length; i++) {
      delete $._burners[burners_[i]];
    }
  }

  /// @custom:legacy
  /// @notice batchMint is used by minters to mint SoulETH in batch.
  function batchMint(
    address[] calldata accounts,
    uint256[] calldata values
  ) external {
    require(accounts.length == values.length, 'invalid arguments');

    SoulETHStorage storage $ = _getSoulETHStorage();
    require($._minters[_msgSender()], 'not a minter');

    for (uint i = 0; i < accounts.length; i++) {
      _mint(accounts[i], values[i]);
    }
  }

  /// @custom:legacy
  /// @notice burnFrom is used by burners to burn SoulETH.
  function burnFrom(address account, uint256 value) external {
    SoulETHStorage storage $ = _getSoulETHStorage();
    require($._burners[_msgSender()], 'not a burner');
    _burn(account, value);
  }

  /// @custom:legacy
  /// @notice transferFrom is disabled for SoulETH.
  function transfer(address, uint256) public virtual override returns (bool) {
    revert('transfer is disabled for SoulETH');
  }

  /// @custom:legacy
  /// @notice transferFrom is disabled for SoulETH.
  function transferFrom(
    address,
    address,
    uint256
  ) public virtual override returns (bool) {
    revert('transferFrom is disabled for SoulETH');
  }

  /// @custom:legacy
  /// @notice approve is disabled for SoulETH.
  function approve(address, uint256) public virtual override returns (bool) {
    revert('approve is disabled for SoulETH');
  }
}
