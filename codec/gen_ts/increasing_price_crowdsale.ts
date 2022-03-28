/* DO NOT EDIT THE BELOW FILE AS THIS IS LIKELY WILL BE GENERATED AGAIN AND REWRITE OVER IT */

// tslint:disable:no-consecutive-blank-lines ordered-imports align trailing-comma enum-naming
// tslint:disable:whitespace no-unbound-method no-trailing-whitespace
// tslint:disable:no-unused-variable

import * as ethers from "ethers"
// eslint-disable-next-line import/named
import {
  assert,
  schemas,
  // eslint-disable-next-line import/named
  SubscriptionManager,
  // eslint-disable-next-line import/named
  BaseContract,
  // eslint-disable-next-line import/named
  EventCallback,
  // eslint-disable-next-line import/named
  IndexedFilterValues,
  // eslint-disable-next-line import/named
  BlockRange,
  // eslint-disable-next-line import/named
  DecodedLogArgs,
  // eslint-disable-next-line import/named
  LogWithDecodedArgs,
  // eslint-disable-next-line import/named
  MethodAbi
} from "vue-blocklink"

import {
  BatchRequest,
  Extension,
  Log,
  PromiEvent,
  provider,
  Providers,
  RLPEncodedTransaction,
  Transaction,
  TransactionConfig,
  TransactionReceipt,
  Common,
  hardfork,
  chain,
  BlockNumber,
  LogsOptions,
  PastLogsOptions
} from "web3-core"

import { Utils, AbiItem } from "web3-utils"
import Web3 from "web3"
import BN from "BN.js"

// tslint:enable:no-unused-variable
export interface ContractInterface {
    buyTokens(beneficiary: string,coin: BN):Promise<void>
    closingTime():Promise<BN>
    finalRate():Promise<BN>
    getCurrentRate():Promise<BN>
    hasClosed():Promise<boolean>
    initialRate():Promise<BN>
    isOpen():Promise<boolean>
    openingTime():Promise<BN>
    rate():Promise<BN>
    token():Promise<string>
    wallet():Promise<string>
    weiRaised():Promise<BN>
}





export enum IncreasingPriceCrowdsaleEvents {
    TimedCrowdsaleExtended = 'TimedCrowdsaleExtended',
    TokensPurchased = 'TokensPurchased',
}

export interface IncreasingPriceCrowdsaleTimedCrowdsaleExtendedEventArgs extends DecodedLogArgs {
    prevClosingTime: BN;
    newClosingTime: BN;
}

export interface IncreasingPriceCrowdsaleTokensPurchasedEventArgs extends DecodedLogArgs {
    purchaser: string;
    beneficiary: string;
    value: BN;
    amount: BN;
}


export type IncreasingPriceCrowdsaleEventArgs =
    | IncreasingPriceCrowdsaleTimedCrowdsaleExtendedEventArgs
    | IncreasingPriceCrowdsaleTokensPurchasedEventArgs;




/* istanbul ignore next */
// tslint:disable:array-type
// tslint:disable:no-parameter-reassignment
// tslint:disable-next-line:class-name
export class IncreasingPriceCrowdsaleContract extends BaseContract implements ContractInterface{
    /**
     * @ignore
     */
public static deployedBytecode: string | undefined;
public static readonly contractName = 'IncreasingPriceCrowdsale';
    private readonly _methodABIIndex: { [name: string]: number } = {};
    //todo: we will come back and fix it later not generic Error @https://www.typescriptlang.org/docs/handbook/2/conditional-types.html
    // @ts-ignore
    private readonly _subscriptionManager: SubscriptionManager<IncreasingPriceCrowdsaleEventArgs, IncreasingPriceCrowdsaleEvents>;


    public static Instance(): (IncreasingPriceCrowdsaleContract | any | boolean) {
        if (window && window.hasOwnProperty("__eth_contract_IncreasingPriceCrowdsale")) {
          // @ts-ignore
          const obj = window.__eth_contract_IncreasingPriceCrowdsale
          if (obj instanceof IncreasingPriceCrowdsaleContract) {
            return (obj) as IncreasingPriceCrowdsaleContract
          } else {
            return (obj) as IncreasingPriceCrowdsaleContract
          }
        } else {
          return false
        }
    }

    static async init(
        contract_address: string,
        supportedProvider: provider,
        ww3: Web3
        ):Promise<IncreasingPriceCrowdsaleContract>
    {
        const contractInstance = await new IncreasingPriceCrowdsaleContract(
            contract_address,
            supportedProvider,
            ww3
        );

        contractInstance.constructorArgs = [/** "initialRate",
"finalRate",
"startTime",
"endTime",
"rate",
"wallet",
"token"
 **/];

        if (window && !window.hasOwnProperty("__eth_contract_IncreasingPriceCrowdsale")) {
            // @ts-ignore
            window.__eth_contract_IncreasingPriceCrowdsale = contractInstance
        }

        return contractInstance
    }

    /**
     * @returns The contract ABI
     */
    public static ABI(): AbiItem[] {
        const abi = [
            { 
                inputs: [
                    {
                        name: 'initialRate',
                        type: 'uint256',
                    },
                    {
                        name: 'finalRate',
                        type: 'uint256',
                    },
                    {
                        name: 'startTime',
                        type: 'uint256',
                    },
                    {
                        name: 'endTime',
                        type: 'uint256',
                    },
                    {
                        name: 'rate',
                        type: 'uint256',
                    },
                    {
                        name: 'wallet',
                        type: 'address',
                    },
                    {
                        name: 'token',
                        type: 'address',
                    },
                ],
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'constructor',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'prevClosingTime',
                        type: 'uint256',
                        indexed: false,
                    },
                    {
                        name: 'newClosingTime',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'TimedCrowdsaleExtended',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'purchaser',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'beneficiary',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'value',
                        type: 'uint256',
                        indexed: false,
                    },
                    {
                        name: 'amount',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'TokensPurchased',
                outputs: [
                ],
                type: 'event',
            },
            { 
                inputs: [
                ],
                outputs: [
                ],
                stateMutability: 'payable',
                type: 'fallback',
            },
            { 
                inputs: [
                    {
                        name: 'beneficiary',
                        type: 'address',
                    },
                ],
                name: 'buyTokens',
                outputs: [
                ],
                stateMutability: 'payable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'closingTime',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'finalRate',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'getCurrentRate',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'hasClosed',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'initialRate',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'isOpen',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'openingTime',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'rate',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'token',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'wallet',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'weiRaised',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
        ] as AbiItem[];
        return abi;
    }

    /**
     the listed content for the contract functions
    **/

    public async buyTokens(beneficiary: string,valCoin: BN): Promise<void> {
        const self = this as any as IncreasingPriceCrowdsaleContract;

            assert.isString('beneficiary', beneficiary);

        const promizz = self._contract.methods.buyTokens(
            beneficiary,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice, value: valCoin
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async buyTokensGas(beneficiary: string,): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.buyTokens(beneficiary,).estimateGas();
        return gasAmount;
    };


    public async closingTime(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.closingTime(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async closingTimeGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.closingTime().estimateGas();
        return gasAmount;
    };


    public async finalRate(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.finalRate(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async finalRateGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.finalRate().estimateGas();
        return gasAmount;
    };


    public async getCurrentRate(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.getCurrentRate(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getCurrentRateGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.getCurrentRate().estimateGas();
        return gasAmount;
    };


    public async hasClosed(): Promise<boolean> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.hasClosed(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async hasClosedGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.hasClosed().estimateGas();
        return gasAmount;
    };


    public async initialRate(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.initialRate(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async initialRateGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.initialRate().estimateGas();
        return gasAmount;
    };


    public async isOpen(): Promise<boolean> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.isOpen(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isOpenGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.isOpen().estimateGas();
        return gasAmount;
    };


    public async openingTime(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.openingTime(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async openingTimeGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.openingTime().estimateGas();
        return gasAmount;
    };


    public async rate(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.rate(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async rateGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.rate().estimateGas();
        return gasAmount;
    };


    public async token(): Promise<string> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.token(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async tokenGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.token().estimateGas();
        return gasAmount;
    };


    public async wallet(): Promise<string> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.wallet(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async walletGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.wallet().estimateGas();
        return gasAmount;
    };


    public async weiRaised(): Promise<BN> {
        const self = this as any as IncreasingPriceCrowdsaleContract;


        const promizz = self._contract.methods.weiRaised(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async weiRaisedGas(): Promise<number>{
        const self = this as any as IncreasingPriceCrowdsaleContract;
        const gasAmount = await self._contract.methods.weiRaised().estimateGas();
        return gasAmount;
    };


    /**
     * Subscribe to an event type emitted by the IncreasingPriceCrowdsale contract.
     * @param eventName The IncreasingPriceCrowdsale contract event you would like to subscribe to.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{maker: aUserAddressHex}`
     * @param callback Callback that gets called when a log is added/removed
     * @param isVerbose Enable verbose subscription warnings (e.g recoverable network issues encountered)
     * @return Subscription token used later to unsubscribe
     */
    public subscribe<ArgsType extends IncreasingPriceCrowdsaleEventArgs>(
        eventName: IncreasingPriceCrowdsaleEvents,
        indexFilterValues: IndexedFilterValues,
        callback: EventCallback<ArgsType>,
        isVerbose: boolean = false,
        blockPollingIntervalMs?: number,
    ): string {
        assert.doesBelongToStringEnum('eventName', eventName, IncreasingPriceCrowdsaleEvents);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        assert.isFunction('callback', callback);
        const subscriptionToken = this._subscriptionManager.subscribe<ArgsType>(
            this._address,
            eventName,
            indexFilterValues,
            IncreasingPriceCrowdsaleContract.ABI(),
            callback,
            isVerbose,
            blockPollingIntervalMs,
        );
        return subscriptionToken;
    }

    /**
     * Cancel a subscription
     * @param subscriptionToken Subscription token returned by `subscribe()`
     */
    public unsubscribe(subscriptionToken: string): void {
        this._subscriptionManager.unsubscribe(subscriptionToken);
    }

    /**
     * Cancels all existing subscriptions
     */
    public unsubscribeAll(): void {
        this._subscriptionManager.unsubscribeAll();
    }


    /**
     * Gets historical logs without creating a subscription
     * @param eventName The IncreasingPriceCrowdsale contract event you would like to subscribe to.
     * @param blockRange Block range to get logs from.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{_from: aUserAddressHex}`
     * @return Array of logs that match the parameters
     */
    public async getLogsAsync<ArgsType extends IncreasingPriceCrowdsaleEventArgs>(
        eventName: IncreasingPriceCrowdsaleEvents,
        blockRange: BlockRange,
        indexFilterValues: IndexedFilterValues,
    ): Promise<Array<LogWithDecodedArgs<ArgsType>>> {
        assert.doesBelongToStringEnum('eventName', eventName, IncreasingPriceCrowdsaleEvents);
        assert.doesConformToSchema('blockRange', blockRange, schemas.blockRangeSchema);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        const logs = await this._subscriptionManager.getLogsAsync<ArgsType>(
            this._address,
            eventName,
            blockRange,
            indexFilterValues,
            IncreasingPriceCrowdsaleContract.ABI(),
        );
        return logs;
    }

    constructor(
        address: string,
        supportedProvider: provider,
        ww3: Web3
    ) {
        super('IncreasingPriceCrowdsale', IncreasingPriceCrowdsaleContract.ABI(), address, supportedProvider,ww3);
        this._subscriptionManager = new SubscriptionManager(
            IncreasingPriceCrowdsaleContract.ABI(),
            supportedProvider,
        );

        IncreasingPriceCrowdsaleContract.ABI().forEach((item, index) => {
            if (item.type === 'function') {
                const methodAbi = item as MethodAbi;
                this._methodABIIndex[methodAbi.name] = index;
            }
        });


    }
}

// tslint:disable:max-file-line-count
// tslint:enable:no-unbound-method no-parameter-reassignment no-consecutive-blank-lines ordered-imports align
// tslint:enable:trailing-comma whitespace no-trailing-whitespace
