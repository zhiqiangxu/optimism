// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20Upgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/// @title SoulGasToken
/// @notice The SoulGasToken is a soul-bounded ERC20 contract which cab be used to pay gas on L2.
///         It has 2 modes:
///             1. when _isSoulQKC, the token can be mint by anyone depositing native token into the contract.
///             2. when !_isSoulQKC, the token can only be mint by whitelist minters specified by
///                contract owner.
contract SoulGasToken is ERC20Upgradeable, OwnableUpgradeable {
    /// @custom:storage-location erc7201:openzeppelin.storage.SoulGasToken
    struct SoulGasTokenStorage {
        bool _isSoulQKC;
        // used when _isSoulQKC, will be a hardcoded dead address called from op-geth to reduce gas
        address _burner;
        // used when !_isSoulQKC
        // _minters will be whitelist EOA addresses
        // _burners can either be whitelist EOA addresses or a hardcoded dead address called from op-geth to reduce gas
        mapping(address => bool) _minters;
        mapping(address => bool) _burners;
    }

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.SoulGasToken")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant _SOULGASTOKEN_STORAGE_LOCATION =
        0x135c38e215d95c59dcdd8fe622dccc30d04cacb8c88c332e4e7441bac172dd00;

    function _getSoulGasTokenStorage() private pure returns (SoulGasTokenStorage storage $) {
        assembly {
            $.slot := _SOULGASTOKEN_STORAGE_LOCATION
        }
    }

    constructor() {
        _disableInitializers();
    }

    /// @custom:legacy
    /// @notice initialize is used to initialize SoulGasToken contract.
    function initialize(
        bool isSoulQKC_,
        string calldata name_,
        string calldata symbol_,
        address calldata burner_,
        address owner_,
        address[] calldata minters_,
        address[] calldata burners_
    )
        external
        initializer
    {
        if (isSoulQKC_) {
            require(owner_ == address(0), "owner_ should be zero when isSoulQKC_");
            require(minters_.length == 0, "minters_ should be empty when isSoulQKC_");
            require(burners_.length == 0, "burners_ should be empty when isSoulQKC_");
        } else {
            require(burner_ == address(0), "burners_ should be empty when !isSoulQKC_");
            require(owner_ != address(0), "owner_ should not be zero when !isSoulQKC_");
            require(minters_.length != 0, "minters_ should not be empty when !isSoulQKC_");
            require(burners_.length != 0, "burners_ should not be empty when !isSoulQKC_");
        }

        // even though owner is only used when _isSoulQKC, we always initialize the inherited
        // OwnableUpgradeable
        __Ownable_init();
        transferOwnership(owner_);
        // initialize the inherited ERC20Upgradeable
        __ERC20_init(name_, symbol_);

        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        $._isSoulQKC = isSoulQKC_;
        $._burner = burner_;
        uint256 i;
        for (i = 0; i < minters_.length; i++) {
            $._minters[minters_[i]] = true;
        }
        for (i = 0; i < burners_.length; i++) {
            $._burners[burners_[i]] = true;
        }
    }

    /// @custom:legacy
    /// @notice deposit is used by anyone to deposit native token for SoulGasToken when _isSoulQKC.
    function deposit() external payable {
        _mint(_msgSender(), msg.value);
    }

    /// @custom:legacy
    /// @notice batchDeposit is used by anyone to deposit native token for SoulGasToken in batch when _isSoulQKC.
    function batchDeposit(address[] calldata accounts, uint256[] calldata values) external payable {
        require(accounts.length == values.length, "invalid arguments");

        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        require($._isSoulQKC, "batchDeposit should only be called when _isSoulQKC");

        uint256 totalValue = 0;
        for (uint256 i = 0; i < accounts.length; i++) {
            _mint(accounts[i], values[i]);
            totalValue += values[i];
        }
        require(msg.value == totalValue, "unexpected msg.value");
    }

    /// @custom:legacy
    /// @notice batchMint is used by minters to mint SoulGasToken in batch when !_isSoulQKC.
    function batchMint(address[] calldata accounts, uint256[] calldata values) external {
        require(accounts.length == values.length, "invalid arguments");

        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        require(!$._isSoulQKC, "batchMint should only be called when !_isSoulQKC");
        require($._minters[_msgSender()], "not a minter");

        for (uint256 i = 0; i < accounts.length; i++) {
            _mint(accounts[i], values[i]);
        }
    }

    /// @custom:legacy
    /// @notice addMinters is used by owner to add minters when !_isSoulQKC.
    function addMinters(address[] calldata minters_) external onlyOwner {
        // addMinters should only be called when !_isSoulQKC, but we don't check it explicitly since it's ensured by
        // onlyOwner
        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        uint256 i;
        for (i = 0; i < minters_.length; i++) {
            $._minters[minters_[i]] = true;
        }
    }

    /// @custom:legacy
    /// @notice delMinters is used by owner to delete minters when !_isSoulQKC.
    function delMinters(address[] calldata minters_) external onlyOwner {
        // delMinters should only be called when !_isSoulQKC, but we don't check it explicitly since it's ensured by
        // onlyOwner
        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        uint256 i;
        for (i = 0; i < minters_.length; i++) {
            delete $._minters[minters_[i]];
        }
    }

    /// @custom:legacy
    /// @notice addBurners is used by owner to add burners when !_isSoulQKC.
    function addBurners(address[] calldata burners_) external onlyOwner {
        // addBurners should only be called when !_isSoulQKC, but we don't check it explicitly since it's ensured by
        // onlyOwner
        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        uint256 i;
        for (i = 0; i < burners_.length; i++) {
            $._burners[burners_[i]] = true;
        }
    }

    /// @custom:legacy
    /// @notice delBurners is used by owner to delete burners when !_isSoulQKC.
    function delBurners(address[] calldata burners_) external onlyOwner {
        // delBurners should only be called when !_isSoulQKC, but we don't check it explicitly since it's ensured by
        // onlyOwner
        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        uint256 i;
        for (i = 0; i < burners_.length; i++) {
            delete $._burners[burners_[i]];
        }
    }

    /// @custom:legacy
    /// @notice burnFrom is used by the burner to burn SoulGasToken.
    function burnFrom(address account, uint256 value) external {
        SoulGasTokenStorage storage $ = _getSoulGasTokenStorage();
        require(
            ($._isSoulQKC && $._burner == _msgSender()) || (!$._isSoulQKC && $._burners[_msgSender()]), "not the burner"
        );
        _burn(account, value);
    }

    /// @custom:legacy
    /// @notice transferFrom is disabled for SoulGasToken.
    function transfer(address, uint256) public virtual override returns (bool) {
        revert("transfer is disabled for SoulGasToken");
    }

    /// @custom:legacy
    /// @notice transferFrom is disabled for SoulGasToken.
    function transferFrom(address, address, uint256) public virtual override returns (bool) {
        revert("transferFrom is disabled for SoulGasToken");
    }

    /// @custom:legacy
    /// @notice approve is disabled for SoulGasToken.
    function approve(address, uint256) public virtual override returns (bool) {
        revert("approve is disabled for SoulGasToken");
    }
}
