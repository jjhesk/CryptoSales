pragma solidity ^0.5.0;

import "../Crowdsale.sol";
import "../../vendor/v0.6/WhitelistedRole.sol";

/**
 * @title PausableCrowdsale
 * @dev Extension of Crowdsale contract where purchases can be paused and unpaused by the pauser role.
 */
contract PausableCrowdsale is Crowdsale, WhitelistedRole {
    /**
     * @dev Validation of an incoming purchase. Use require statements to revert state when conditions are not met.
     * Use super to concatenate validations.
     * Adds the validation that the crowdsale must not be paused.
     * @param _beneficiary Address performing the token purchase
     * @param _weiAmount Value in wei involved in the purchase
     */
    function _preValidatePurchase(address _beneficiary, uint256 _weiAmount) internal override view pauseCheck {
        return super._preValidatePurchase(_beneficiary, _weiAmount);
    }
}
