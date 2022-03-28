pragma solidity >=0.6.0;

import "../Crowdsale.sol";
import "../../vendor/v0.6/IERC20.sol";
import "../../vendor/v0.6/SafeERC20.sol";
import "../../vendor/v0.6/SafeMath.sol";


library Math {
    /**
     * @dev Returns the largest of two numbers.
     */
    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a >= b ? a : b;
    }

    /**
     * @dev Returns the smallest of two numbers.
     */
    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two numbers. The result is rounded towards
     * zero.
     */
    function average(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b) / 2 can overflow, so we distribute
        return (a / 2) + (b / 2) + (((a % 2) + (b % 2)) / 2);
    }
}

/**
 * @title AllowanceCrowdsale
 * @dev Extension of Crowdsale where tokens are held by a wallet, which approves an allowance to the crowdsale.
 */
contract AllowanceCrowdsale is Crowdsale {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    address private _tokenWallet;

    /**
     * @dev Constructor, takes token wallet address.
     * @param tokenWallet Address holding the tokens, which has approved allowance to the crowdsale.
     */
    constructor (address tokenWallet) public {
        require(tokenWallet != address(0), "AllowanceCrowdsale: token wallet is the zero address");
        _tokenWallet = tokenWallet;
    }

    /**
     * @return the address of the wallet that will hold the tokens.
     */
    function tokenWallet() public view returns (address) {
        return _tokenWallet;
    }

    /**
     * @dev Checks the amount of tokens left in the allowance.
     * @return Amount of tokens left in the allowance
     */
    function remainingTokens() public view returns (uint256) {
        return Math.min(token().balanceOf(_tokenWallet), token().allowance(_tokenWallet, address(this)));
    }

    /**
     * @dev Overrides parent behavior by transferring tokens from wallet.
     * @param beneficiary Token purchaser
     * @param tokenAmount Amount of tokens purchased
     */
    function _deliverTokens(address beneficiary, uint256 tokenAmount) internal {
        token().safeTransferFrom(_tokenWallet, beneficiary, tokenAmount);
    }
}
