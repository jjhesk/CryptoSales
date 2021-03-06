pragma solidity ^0.6.0;

import "../Crowdsale.sol";
import "../../vendor/v0.6/WhitelistedRole.sol";
import "../../vendor/v0.6/IERC20.sol";


/**
 * @title WhitelistCrowdsale
 * @dev Crowdsale in which only whitelisted users can contribute.
 */
contract WhitelistCrowdsale is WhitelistedRole, Crowdsale {
    constructor (uint256 rate, address payable wallet, IERC20 token)
    Crowdsale(rate, wallet, token)
    public {

    }
    /**
     * @dev Extend parent behavior requiring beneficiary to be whitelisted. Note that no
     * restriction is imposed on the account sending the transaction.
     * @param _beneficiary Token beneficiary
     * @param _weiAmount Amount of wei contributed
     */
    function _preValidatePurchase(address _beneficiary, uint256 _weiAmount) internal override view {
        require(isWhitelisted(_beneficiary), "WhitelistCrowdsale: beneficiary doesn't have the Whitelisted role");
        super._preValidatePurchase(_beneficiary, _weiAmount);
    }

}
