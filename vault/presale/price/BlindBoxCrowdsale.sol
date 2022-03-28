pragma solidity ^0.6.0;

import "../../vendor/v0.6/SafeMath.sol";
import "../validation/TimedCrowdsale.sol";

contract BlindBoxCrowdsale is TimedCrowdsale {
    using SafeMath for uint256;

    uint256 private _upperRate;
    uint256 private _lowerRate;


    /**
     * @dev Constructor, takes initial and final rates of tokens received per wei contributed.
     * @param minRate Number of tokens a buyer gets per wei at the start of the crowdsale
     * @param maxRate Number of tokens a buyer gets per wei at the end of the crowdsale
     * @param startTime Number of tokens a buyer gets per wei at the end of the crowdsale
     * @param endTime Number of tokens a buyer gets per wei at the end of the crowdsale
     * @param rate Number of tokens a buyer gets per wei at the end of the crowdsale
     * @param wallet Number of tokens a buyer gets per wei at the end of the crowdsale
     * @param token Number of tokens a buyer gets per wei at the end of the crowdsale
     */
    constructor (
        uint256 minRate,
        uint256 maxRate,
        uint256 startTime,
        uint256 endTime,
        uint256 rate,
        address payable wallet,
        IERC20 token
    ) public TimedCrowdsale(startTime, endTime) Crowdsale(rate, wallet, token) {
        require(minRate > 0 && maxRate > 0 && maxRate > minRate, "BlindBoxSale: needs to be the corrected rate");
        _upperRate = maxRate;
        _lowerRate = minRate;
    }

    /**
     * @return the max rate of the crowdsale.
     */
    function ceilingRate() public view returns (uint256) {
        return _upperRate;
    }

    /**
     * @return the min rate of the crowdsale.
     */
    function floorRate() public view returns (uint256) {
        return _lowerRate;
    }


    /**
   * @dev Returns the rate of tokens per wei at the present time.
   * Note that, as price _increases_ with time, the rate _decreases_.
   * @return The number of tokens a buyer gets per wei at a given time
   */
    function getCurrentRate() public view returns (uint256) {
        if (!isOpen()) {
            return 0;
        }

        // solhint-disable-next-line not-rely-on-time
        //uint256 elapsedTime = block.timestamp.sub(openingTime());
        //uint256 timeRange = closingTime().sub(openingTime());
        //uint256 rateRange = _upperRate.sub(_lowerRate);

        return getRandomNumber(_upperRate, _lowerRate, _msgSender());
        //  return _initialRate.sub(elapsedTime.mul(rateRange).div(timeRange));
    }

    function getRandomNumber(uint256 maxRandom, uint256 min, address privateAddress) public view returns (uint8) {
        uint256 genNum = uint256(blockhash(block.number - 1)) + uint256(privateAddress);
        return uint8(genNum % (maxRandom - min + 1) + min);
    }


    /**
     * @dev Override to extend the way in which ether is converted to tokens.
     * @param weiAmount Value in wei to be converted into tokens
     * @return Number of tokens that can be purchased with the specified _weiAmount
     */
    function _getTokenAmount(uint256 weiAmount) internal override view returns (uint256) {
        uint256 currentRate = getCurrentRate();
        return currentRate.mul(weiAmount);
    }


}
