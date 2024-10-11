// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {L2CustomERC20Gateway} from "../l2/gateways/L2CustomERC20Gateway.sol";
import {L2GatewayRouter} from "../l2/gateways/L2GatewayRouter.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {L2ToL1MessagePasser} from "../l2/system/L2ToL1MessagePasser.sol";
import {L2CrossDomainMessenger} from "../l2/L2CrossDomainMessenger.sol";
import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";
import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";

contract L2CustomERC20GatewayTest is L2GatewayBaseTest {
    L2CustomERC20Gateway private gateway;
    L2GatewayRouter private router;
    L2CrossDomainMessenger private l2Messenger;

    address private counterpartGateway;
    address private feeVault;
    address private l1Messenger;

    MockERC20 private l1Token;
    MockERC20 private l2Token;

    function setUp() public override {
        super.setUp();
        _deployCustomERC20();
        gateway = l2CustomERC20Gateway;
        router = l2GatewayRouter;
        counterpartGateway = gateway.counterpart();
        l2Messenger = l2CrossDomainMessenger;
        feeVault = l2FeeVault;
        l1Messenger = address(NON_ZERO_ADDRESS);

        // transfer ownership
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        // Deploy tokens
        l1Token = new MockERC20("Mock L1", "ML1", 18);
        l2Token = new MockERC20("Mock L2", "ML2", 18);

        // Prepare token balances
        l2Token.mint(address(this), type(uint128).max);
        l2Token.approve(address(gateway), type(uint256).max);
    }

    function test_initialize_initializer_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1), address(1));
    }
    function test_initialize_zeroAddress_reverts() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for L2CustomERC20Gateway.
        TransparentUpgradeableProxy l2CustomERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2CustomERC20Gateway contract.
        L2CustomERC20Gateway l2CustomERC20GatewayImplTemp = new L2CustomERC20Gateway();
        
        // Expect revert due to zero router address.
        hevm.expectRevert("zero router address");
        ITransparentUpgradeableProxy(address(l2CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l2CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L2CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );
        
        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l2CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l2CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L2CustomERC20Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(NON_ZERO_ADDRESS), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );
        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l2CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l2CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L2CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(NON_ZERO_ADDRESS), // _router
                    address(0) // _messenger
                )
            )
        );
        hevm.stopPrank();
    }
    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for the L2CustomERC20Gateway.
        TransparentUpgradeableProxy l2CustomERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2CustomERC20Gateway contract.
        L2CustomERC20Gateway l2CustomERC20GatewayImplTemp = new L2CustomERC20Gateway();
        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l2CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L2CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2GatewayRouter), // _router
                    address(l2CrossDomainMessenger) // _messenger
                )
            )
        );
        // Cast the proxy contract address to the L2CustomERC20Gateway contract type to call its methods.
        L2CustomERC20Gateway l2CustomERC20GatewayTemp = L2CustomERC20Gateway((address(l2CustomERC20GatewayProxyTemp)));
        hevm.stopPrank();
        
        // Verify the counterpart, router and messenger are initialized successfully.
        assertEq(l2CustomERC20GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l2CustomERC20GatewayTemp.router(), address(l2GatewayRouter));
        assertEq(l2CustomERC20GatewayTemp.messenger(), address(l2CrossDomainMessenger));
    }

    function test_updateTokenMapping_onlyOwner_fails(address token2) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token2, token2);
        hevm.stopPrank();

        // l1 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token2, address(0));
    }

    function test_updateTokenMapping_succeeds(address token1, address token2) public {
        hevm.assume(token1 != address(0));
        
        // Expect the UpdateTokenMapping event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit L2CustomERC20Gateway.UpdateTokenMapping(token2, address(0), token1);

        assertEq(gateway.getL1ERC20Address(token2), address(0));
        gateway.updateTokenMapping(token2, token1);
        assertEq(gateway.getL1ERC20Address(token2), token1);
    }

    function test_withdrawERC20_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _withdrawERC20(false, amount, gasLimit, feePerGas);
    }

    function test_withdrawERC20WithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipient(false, amount, recipient, gasLimit, feePerGas);
    }

    function test_withdrawERC20WithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipientAndCalldata(false, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_finalizeDepositERC20_mocking_fails(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, 100000);

        // revert when caller is not messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositERC20(address(l1Token), address(l2Token), sender, recipient, amount, dataToCall);

        MockCrossDomainMessenger mockMessenger = new MockCrossDomainMessenger();
        hevm.store(address(gateway), bytes32(eth_erc20_messenger_slot), bytes32(abi.encode(address(mockMessenger))));

        // only call by counterpart
        hevm.expectRevert("only call by counterpart");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeCall(
                gateway.finalizeDepositERC20,
                (address(l1Token), address(l2Token), sender, recipient, amount, dataToCall)
            )
        );

        mockMessenger.setXDomainMessageSender(address(counterpartGateway));

        // msg.value mismatch
        hevm.expectRevert("nonzero msg.value");
        mockMessenger.callTarget{value: 1}(
            address(gateway),
            abi.encodeCall(
                gateway.finalizeDepositERC20,
                (address(l1Token), address(l2Token), sender, recipient, amount, dataToCall)
            )
        );

        // l1 token mismatch
        hevm.expectRevert("l1 token mismatch");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeCall(
                gateway.finalizeDepositERC20,
                (address(l1Token), address(l2Token), sender, recipient, amount, dataToCall)
            )
        );
    }

    function test_finalizeDepositERC20_counterError_fails(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));

        gateway.updateTokenMapping(address(l2Token), address(l1Token));

        amount = bound(amount, 1, l2Token.balanceOf(address(this)));

        // do finalize deposit token
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), sender, recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        // counterpart is not L1CustomERC20Gateway
        // emit FailedRelayedMessage from L2CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
        uint256 recipientBalance = l2Token.balanceOf(recipient);
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
        hevm.startPrank(AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger)));
        l2Messenger.relayMessage(address(uint160(address(counterpartGateway)) + 1), address(gateway), 0, 0, message);
        hevm.stopPrank();
        assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l2Token.balanceOf(recipient));
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
    }

    function test_finalizeDepositERC20_succeeds(address sender, uint256 amount, bytes memory dataToCall) public {
        address recipient = address(2048);

        gateway.updateTokenMapping(address(l2Token), address(l1Token));

        amount = bound(amount, 1, l2Token.balanceOf(address(this)));

        // do finalize deposit token
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), sender, address(recipient), amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );

        // emit FinalizeDepositERC20 from L2CustomERC20Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL2ERC20Gateway.FinalizeDepositERC20(
                address(l1Token),
                address(l2Token),
                sender,
                address(recipient),
                amount,
                dataToCall
            );
        }

        // emit RelayedMessage from L2CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
        uint256 recipientBalance = l2Token.balanceOf(address(recipient));
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
        hevm.startPrank(AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger)));
        l2Messenger.relayMessage(address(counterpartGateway), address(gateway), 0, 0, message);
        hevm.stopPrank();
        assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
        assertEq(recipientBalance + amount, l2Token.balanceOf(address(recipient)));
        assertBoolEq(true, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
    }

    function _withdrawERC20(bool useRouter, uint256 amount, uint256 gasLimit, uint256 feePerGas) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), address(this), address(this), amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        } else {
            gateway.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
            } else {
                gateway.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
            }
        } else {
            // emit AppendMessage from L2MessageQueue
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL2ERC20Gateway.WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
            } else {
                gateway.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
            }
            assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _withdrawERC20WithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), address(this), recipient, amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        } else {
            gateway.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(address(l2Token), recipient, amount, gasLimit);
            } else {
                gateway.withdrawERC20{value: feeToPay}(address(l2Token), recipient, amount, gasLimit);
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL2ERC20Gateway.WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(address(l2Token), recipient, amount, gasLimit);
            } else {
                gateway.withdrawERC20{value: feeToPay}(address(l2Token), recipient, amount, gasLimit);
            }
            assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _withdrawERC20WithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), address(this), recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        } else {
            gateway.withdrawERC20{value: feeToPay}(address(l2Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(address(l2Token), recipient, amount, dataToCall, gasLimit);
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL2ERC20Gateway.WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                dataToCall,
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(address(l2Token), recipient, amount, dataToCall, gasLimit);
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }
}
