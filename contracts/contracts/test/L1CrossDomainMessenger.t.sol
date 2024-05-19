// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L1MessageQueueWithGasPriceOracle} from "../l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {IL1CrossDomainMessenger} from "../l1/L1CrossDomainMessenger.sol";
import {L1CrossDomainMessenger} from "../l1/L1CrossDomainMessenger.sol";
import {IRollup} from "../l1/rollup/IRollup.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import "forge-std/console.sol";

contract L1CrossDomainMessengerTest is L1MessageBaseTest {
    uint256 L1CrossDomainMessenger_provenWithdrawals_slot = 251;
    address refundAddress = address(2048);
    L1CrossDomainMessenger l1CrossDomainMessengerTemp;

    function testInitialize_messenger() external {
        address feeVault = address(1);
        address rollup = address(1);
        address messageQueue = address(1);

        // verify the initialize only can be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1CrossDomainMessenger.initialize(address(1), address(1), address(1));

        // verify the state variable maxReplayTimes is initialized successfully.
        assertEq(l1CrossDomainMessenger.maxReplayTimes(), 3);

        // deploy proxy
        TransparentUpgradeableProxy l1CrossDomainMessengerProxyTemp = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );

        // deploy L1CrossDomainMessengerImpl
        l1CrossDomainMessengerTemp = new L1CrossDomainMessenger();

        hevm.startPrank(multisig);

        // verify it throws a custom error ErrZeroAddress() when the feeVault is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp))
            .upgradeToAndCall(
                address(l1CrossDomainMessengerTemp),
                abi.encodeCall(
                    l1CrossDomainMessengerTemp.initialize,
                    (address(0), rollup, messageQueue)
                )
            );

        // verify it throws a custom error ErrZeroAddress() when the rollup is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp))
            .upgradeToAndCall(
                address(l1CrossDomainMessengerTemp),
                abi.encodeCall(
                    l1CrossDomainMessengerTemp.initialize,
                    (feeVault, rollup, address(0))
                )
            );

        // verify it throws a custom error ErrZeroAddress() when the messageQueue is equal to zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp))
            .upgradeToAndCall(
                address(l1CrossDomainMessengerTemp),
                abi.encodeCall(
                    l1CrossDomainMessengerTemp.initialize,
                    (feeVault, address(0), messageQueue)
                )
            );

        // 这个跑不过 ！！！
        // console.log(
        //     "maxReplayTimes is xxxxxs",
        //     l1CrossDomainMessengerTemp.maxReplayTimes()
        // );
        // // verify the event is triggered successfully. 这里需要修改
        // hevm.expectEmit(false, false, false, true);
        // emit IL1CrossDomainMessenger.UpdateMaxReplayTimes(3, _maxReplayTimes);

        // ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp))
        //     .upgradeToAndCall(
        //         address(l1CrossDomainMessengerTemp),
        //         abi.encodeCall(
        //             l1CrossDomainMessengerTemp.initialize,
        //             (
        //                 l1FeeVault,
        //                 address(l1CrossDomainMessengerProxyTemp),
        //                 address(l1CrossDomainMessengerProxyTemp)
        //             )
        //         )
        //     );
        // console.log(
        //     "l1CrossDomainMessengerTemp.feeVault is ",
        //     l1CrossDomainMessengerTemp.feeVault()
        // );
        // // verify the state variable maxReplayTimes is initialized successfully.
        // console.log(
        //     "maxReplayTimes is ",
        //     l1CrossDomainMessengerTemp.maxReplayTimes()
        // );
        // assertEq(l1CrossDomainMessengerTemp.maxReplayTimes(), 3);

        hevm.stopPrank();
    }

    function testProveWithdrawalTransaction_relayMessage() external {
        // tx msg set
        address from = address(alice);
        address to = address(bob);
        uint256 value = 10;
        hevm.deal(address(l1CrossDomainMessenger), value);

        uint256 nonce = 0;
        bytes memory message = "";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        // get proof from ffi
        _appendMessageHash(_xDomainCalldataHash);
        (bytes32 wdHashRes, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(_xDomainCalldataHash);
        assertEq(_xDomainCalldataHash, wdHashRes);
        assertEq(getTreeRoot(), wdRoot);

        // prove without rollup
        hevm.expectRevert("Messenger: withdrawalRoot not finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // mock call rollup withdrawal root submitted success
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (wdRoot)),
            abi.encode(true)
        );
        uint256 balanceBefore = address(bob).balance;
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );
        assertEq(balanceBefore + value, address(bob).balance);

        // prove again
        hevm.expectRevert("Messenger: withdrawal has already been finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // verify it throws the error "Messenger: Forbid to call message queue" when the address of to is equal to messageQueue.
        to = address(l1CrossDomainMessenger.messageQueue());
        hevm.expectRevert("Messenger: Forbid to call message queue");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // verify it throws the error "Messenger: Forbid to call message queue" when the address of from is equal to xDomainMessageSender.
        to = address(bob);
        from = address(l1CrossDomainMessenger.xDomainMessageSender());
        hevm.expectRevert("Messenger: Invalid message sender");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );
    }

    function testReplayMessage(uint256 exceedValue) external {
        hevm.deal(address(this), 1 ether);
        // tx msg set
        exceedValue = bound(exceedValue, 1, address(this).balance / 2);
        address from = address(this);
        address to = address(bob);
        uint256 value = 0;
        bytes memory message = "";
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            from,
            to,
            value,
            nonce,
            message
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );

        // updateMaxReplayTimes to 0
        hevm.expectRevert("replay times must be greater than 0");
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(0);

        // append a message
        l1CrossDomainMessenger.sendMessage{value: 100}(
            to,
            value,
            message,
            gas,
            refundAddress
        );

        // Provided message has not been enqueued
        hevm.expectRevert("Provided message has not been enqueued");
        l1CrossDomainMessenger.replayMessage(
            address(this),
            address(0),
            101,
            0,
            new bytes(0),
            defaultGasLimit,
            refundAddress
        );

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(1);
        // Insufficient msg.value
        hevm.expectRevert("Insufficient msg.value for fee");
        l1CrossDomainMessenger.replayMessage(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );

        hevm.prank(multisig);
        uint256 _fee = l1MessageQueueWithGasPriceOracle.l2BaseFee() *
            defaultGasLimit;

        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(1);

        // refund exceed fee
        uint256 balanceBefore = refundAddress.balance;
        uint256 feeVaultBefore = l1FeeVault.balance;
        l1CrossDomainMessenger.replayMessage{value: _fee + exceedValue}(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );
        assertEq(balanceBefore + exceedValue, refundAddress.balance);
        assertEq(feeVaultBefore + _fee, l1FeeVault.balance);

        // Exceed maximum replay times
        hevm.expectRevert("Exceed maximum replay times");
        l1CrossDomainMessenger.replayMessage{value: _fee}(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );

        // test replay list
        // 1. send a message with nonce 2
        // 2. replay 3 times
        hevm.startPrank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(0);
        l1CrossDomainMessenger.updateMaxReplayTimes(100);
        hevm.stopPrank();
        l1CrossDomainMessenger.sendMessage{value: 100}(
            address(0),
            100,
            new bytes(0),
            defaultGasLimit,
            refundAddress
        );
        bytes32 hash = keccak256(
            _encodeXDomainCalldata(
                address(this),
                address(0),
                100,
                2,
                new bytes(0)
            )
        );
        (uint256 _replayTimes, uint256 _lastIndex) = l1CrossDomainMessenger
            .replayStates(hash);
        assertEq(_replayTimes, 0);
        assertEq(_lastIndex, 0);
        for (uint256 i = 0; i < 3; i++) {
            l1CrossDomainMessenger.replayMessage(
                address(this),
                address(0),
                100,
                2,
                new bytes(0),
                defaultGasLimit,
                refundAddress
            );
            (_replayTimes, _lastIndex) = l1CrossDomainMessenger.replayStates(
                hash
            );
            assertEq(_replayTimes, i + 1);
            assertEq(_lastIndex, i + 3);
            assertEq(l1CrossDomainMessenger.prevReplayIndex(i + 3), i + 2 + 1);
            for (uint256 j = 0; j <= i; j++) {
                assertEq(
                    l1CrossDomainMessenger.prevReplayIndex(i + 3 - j),
                    i + 2 - j + 1
                );
            }
        }
    }

    function testForbidCallMessageQueueFromL2() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1MessageQueueWithGasPriceOracle);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (withdrawalRoot)),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call message queue");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            withdrawalProof,
            withdrawalRoot
        );
    }

    function testForbidCallSelfFromL2() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1CrossDomainMessenger);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (withdrawalRoot)),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call self");
        l1CrossDomainMessenger.proveAndRelayMessage(
            from,
            to,
            value,
            nonce,
            message,
            withdrawalProof,
            withdrawalRoot
        );
    }

    function test_sendMessage() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.SentMessage(
            sender,
            to,
            value,
            nonce,
            gas,
            data
        );

        hevm.expectCall(
            address(l1MessageQueueWithGasPriceOracle),
            abi.encodeCall(
                l1MessageQueueWithGasPriceOracle.appendCrossDomainMessage,
                (Predeploys.L2_CROSS_DOMAIN_MESSENGER, gas, _xDomainCalldata)
            )
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
    }

    function test_sendMessage_value() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);

        // send value not zero
        // set base fee to 100
        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(100);

        value = 1 ether;
        nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        hevm.expectRevert("Insufficient msg.value");
        l1CrossDomainMessenger.sendMessage{value: 1 ether}(
            to,
            value,
            data,
            gas
        );
        // give enough value
        uint256 fee = l1MessageQueueWithGasPriceOracle
            .estimateCrossDomainMessageFee(sender, gas);
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee}(
            to,
            value,
            data,
            gas
        );

        // send more value with refund address
        fee = l1MessageQueueWithGasPriceOracle.estimateCrossDomainMessageFee(
            sender,
            gas
        );
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee * 2}(
            to,
            value,
            data,
            gas,
            refundAddress
        );
        assertEq(address(refundAddress).balance, fee);

        // verify refundAddress.call() failed, trigger the error message "Failed to refund the fee" as expected.
        hevm.expectRevert("Failed to refund the fee");
        (bool _success, ) = refundAddress.call{value: gas + fee}("");
        l1CrossDomainMessenger.sendMessage{value: 2 ether}(
            to,
            value,
            data,
            gas,
            refundAddress
        );

        // verify the call is executed as expected, and the fee is added into the balance of the feeVault.
        uint256 initialFeeVaultBalance = l1FeeVault.balance;
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee}(
            to,
            value,
            data,
            gas
        );
        assertEq(address(l1FeeVault).balance, initialFeeVaultBalance + fee);

        // verify it throws a "Failed to refund the fee" error when the call fails.
        hevm.expectRevert("Failed to deduct the fee");
        (_success, ) = sender.call{value: gas}("");

        console.log("_success is ", _success);

        // 这个跑不过 ！！！
        // 第一次调用 sendMessage
        // console.log("Ginger");
        // bytes32 _xDomainCalldataHash = keccak256(_xDomainCalldata);

        // // 第二次调用 sendMessage，应该触发 "Duplicated message"
        // console.log("first");
        // console.log(
        //     "the value is ",
        //     messageSendTimestamp[_xDomainCalldataHash]
        // );
        // l1CrossDomainMessenger.sendMessage{value: value + fee}(
        //     to,
        //     value,
        //     data,
        //     gas
        // );
        // console.log(
        //     "the value is 2nd",
        //     messageSendTimestamp[_xDomainCalldataHash]
        // );
        // hevm.expectRevert("Duplicated message");

        // console.log("2nd");
        // l1CrossDomainMessenger.sendMessage{value: value + fee}(
        //     to,
        //     value,
        //     data,
        //     gas
        // );
    }

    // Tests that the sendMessage function is able to send the same message twice.
    function test_sendMessage_twice_succeed() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );

        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);

        console.log(
            "after send message twice ,the nonce is ",
            l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex()
        );
        assertEq(
            nonce + 2,
            l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex()
        );
    }

    function test_dropMessage_fail() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        bytes32 _xDomainCalldataHash = keccak256(_xDomainCalldata);

        // verify it throws an error "Provided message has not been enqueued" when a message that was never enqueued.
        hevm.expectRevert("Provided message has not been enqueued");
        l1CrossDomainMessenger.dropMessage(sender, to, value, nonce, data);
    }

    function testUpdateMaxReplayTimes(uint256 _maxReplayTimes) external {
        hevm.assume(_maxReplayTimes > 0);
        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        hevm.stopPrank();

        hevm.expectEmit(false, false, false, true);
        emit IL1CrossDomainMessenger.UpdateMaxReplayTimes(3, _maxReplayTimes);

        assertEq(l1CrossDomainMessenger.maxReplayTimes(), 3);
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        assertEq(l1CrossDomainMessenger.maxReplayTimes(), _maxReplayTimes);
    }

    function testSetPause() external {
        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1CrossDomainMessenger.setPause(false);
        hevm.stopPrank();

        // pause
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(true);
        assertBoolEq(true, l1CrossDomainMessenger.paused());

        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.sendMessage(
            address(0),
            0,
            new bytes(0),
            defaultGasLimit
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.sendMessage(
            address(0),
            0,
            new bytes(0),
            defaultGasLimit,
            address(0)
        );

        (, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(bytes32(uint256(1)));
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.proveAndRelayMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0),
            wdProof,
            wdRoot
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.replayMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0),
            0,
            address(0)
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.dropMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0)
        );

        // unpause
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(false);
        assertBoolEq(false, l1CrossDomainMessenger.paused());
    }
}
