// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package distribution_test

import (
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	bbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/distribution"
	sbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	utils "pkg.berachain.dev/polaris/cosmos/testing/e2e"
	network "pkg.berachain.dev/polaris/e2e/localnet/network"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/e2e/localnet/utils"
)

func TestDistributionPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/e2e/precompile/distribution")
}

var _ = Describe("Distribution Precompile", func() {
	var (
		tf                *network.TestFixture
		precompile        *bindings.DistributionModule
		stakingPrecompile *sbindings.StakingModule
		bankPrecompile    *bbindings.BankModule
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
		// Setup the governance precompile.
		precompile, _ = bindings.NewDistributionModule(
			common.HexToAddress("0x69"),
			tf.EthClient(),
		)
		// Setup the staking precompile.
		stakingPrecompile, _ = sbindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient())
		bankPrecompile, _ = bbindings.NewBankModule(
			common.BytesToAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
			tf.EthClient(),
		)
	})

	AfterEach(func() {
		// Dump logs and stop the containter here.
		if !CurrentSpecReport().Failure.IsZero() {
			logs, err := tf.DumpLogs()
			Expect(err).ToNot(HaveOccurred())
			GinkgoWriter.Println(logs)
		}
		Expect(tf.Teardown()).To(Succeed())
	})

	It("should be able to get if withdraw address is enabled", func() {
		res, err := precompile.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(BeTrue())
	})

	It("should be able to set withdraw address with cosmos address", func() {
		addr := sdk.AccAddress("addr")
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.SetWithdrawAddress(txr, common.BytesToAddress(addr))
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("should be able to set withdraw address with ethereum address", func() {
		addr := sdk.AccAddress("addr")
		ethAddr := common.BytesToAddress(addr)
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.SetWithdrawAddress(txr, ethAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("should be able to get delegator reward", func() {
		// Delegate some tokens to an active validator.
		validators, _, err := stakingPrecompile.GetBondedValidators(nil, sbindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		val := validators[0]
		delegateAmt := big.NewInt(123450000000)
		txr := tf.GenerateTransactOpts("alice")
		txr.Value = delegateAmt
		tx, err := stakingPrecompile.Delegate(txr, val.OperatorAddr, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for 2 blocks to be produced, to make sure there are rewards.
		for i := 0; i < 2; i++ {
			Expect(tf.WaitForNextBlock()).To(Succeed())
		}

		// Preview the withdraw rewards.
		rewards, err := precompile.GetTotalDelegatorReward(nil, tf.Address("alice"))
		Expect(err).ToNot(HaveOccurred())
		Expect(rewards).ToNot(BeNil())
		for _, reward := range rewards {
			Expect(reward.Amount.Cmp(new(big.Int))).To(Equal(1))
		}

		// Withdraw the rewards.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = precompile.WithdrawDelegatorReward(txr, tf.Address("alice"), val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("Should be able to call the precompile via the contract", func() {
		// Deploy the contract.
		contractAddress, tx, contract, err := tbindings.DeployDistributionWrapper(
			tf.GenerateTransactOpts("alice"),
			tf.EthClient(),
			common.HexToAddress("0x69"),
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		txr := tf.GenerateTransactOpts("alice")
		amt := big.NewInt(123450000000)
		tx, err = bankPrecompile.Send(txr, contractAddress, []bbindings.CosmosCoin{
			{
				Denom:  "abera",
				Amount: amt,
			},
		})
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Delegate some tokens to a validator.
		validators, _, err := stakingPrecompile.GetBondedValidators(nil, sbindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		val := validators[0]
		amt = big.NewInt(123450000000)
		txr = tf.GenerateTransactOpts("alice")
		txr.Value = amt
		tx, err = contract.Delegate(txr, val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for the 2 block to be produced, to make sure there are rewards.
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// Withdraw the rewards.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = contract.WithdrawRewards(txr, contractAddress, val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Get withdraw address enabled.
		res, err := contract.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(BeTrue())

		// Set the withdraw address.
		addr := sdk.AccAddress("addr")
		ethAddr := common.BytesToAddress(addr)
		txr = tf.GenerateTransactOpts("alice")
		tx, err = contract.SetWithdrawAddress(txr, ethAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})
})
