// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20Upgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

contract SoulQKC is ERC20Upgradeable {
    /// @custom:storage-location erc7201:openzeppelin.storage.SoulQKC
    struct SoulQKCStorage {
        address _burner;
    }

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.SoulQKC")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _SOULQKC_STORAGE_LOCATION =
        0xd407e56e94f51d7d335ac9f230437fa91d6ad9417764e42539ab040755498300;

    function _getSoulQKCStorage() private pure returns (SoulQKCStorage storage $) {
        assembly {
            $.slot := _SOULQKC_STORAGE_LOCATION
        }
    }

    constructor() {
        _disableInitializers();
    }

    /// @custom:legacy
    /// @notice initialize is used to initialize SoulQKC contract.
    function initialize(
        string calldata name_,
        string calldata symbol_,
        address calldata burner_
    )
        external
        initializer
    {
        __ERC20_init(name_, symbol_);
        SoulQKCStorage storage $ = _getSoulQKCStorage();
        $._burner = burner_;
    }

    /// @custom:legacy
    /// @notice batchMint is used by anyone to mint SoulQKC in batch.
    function batchMint(address[] calldata accounts, uint256[] calldata values) external payable {
        require(accounts.length == values.length, "invalid arguments");

        uint256 totalValue = 0;
        for (uint256 i = 0; i < accounts.length; i++) {
            _mint(accounts[i], values[i]);
            totalValue += values[i];
        }
        require(msg.value == totalValue, "unexpected msg.value");
    }

    /// @custom:legacy
    /// @notice burnFrom is used by the burner to burn SoulETH.
    function burnFrom(address account, uint256 value) external {
        SoulETHStorage storage $ = _getSoulQKCStorage();
        require($._burner == _msgSender(), "not the burner");
        _burn(account, value);
    }

    /// @custom:legacy
    /// @notice transferFrom is disabled for SoulQKC.
    function transfer(address, uint256) public virtual override returns (bool) {
        revert("transfer is disabled for SoulQKC");
    }

    /// @custom:legacy
    /// @notice transferFrom is disabled for SoulQKC.
    function transferFrom(address, address, uint256) public virtual override returns (bool) {
        revert("transferFrom is disabled for SoulQKC");
    }

    /// @custom:legacy
    /// @notice approve is disabled for SoulQKC.
    function approve(address, uint256) public virtual override returns (bool) {
        revert("approve is disabled for SoulQKC");
    }
}
