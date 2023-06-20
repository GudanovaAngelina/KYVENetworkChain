package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	globalTypes "github.com/KYVENetwork/chain/x/global/types"
	poolTypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*

TEST CASES - abci.go

* inactive pool should not receive inflation funds
* active pool should receive inflation funds
* pool should split inflation funds depending on operating cost
* pools with zero operating cost should receive nothing
* every pool has zero operating cost

*/

var _ = Describe("abci.go", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		s.App().PoolKeeper.SetParams(s.Ctx(), poolTypes.Params{
			ProtocolInflationShare:  sdk.MustNewDecFromStr("0.1"),
			PoolInflationPayoutRate: sdk.MustNewDecFromStr("0.1"),
		})

		s.App().PoolKeeper.AppendPool(s.Ctx(), poolTypes.Pool{
			Name:           "PoolTest",
			MaxBundleSize:  100,
			StartKey:       "0",
			UploadInterval: 60,
			OperatingCost:  1_000_000,
			Protocol: &poolTypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &poolTypes.UpgradePlan{},
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("inactive pool should not receive inflation funds", func() {
		// ARRANGE
		b1, b2 := uint64(0), uint64(0)

		for t := 0; t < 100; t++ {
			// ACT
			pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
			b1 = uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool.GetPoolAccount(), globalTypes.Denom).Amount.Int64())
			s.Commit()
			b2 = uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

			// ASSERT
			Expect(b1).To(BeZero())
			Expect(b2).To(BeZero())
		}
	})

	It("pool should split inflation funds depending on operating cost", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), poolTypes.Pool{
			Name:           "PoolTest2",
			MaxBundleSize:  100,
			StartKey:       "0",
			UploadInterval: 60,
			OperatingCost:  2_000_000,
			Protocol: &poolTypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &poolTypes.UpgradePlan{},
		})

		// ACT
		s.Commit()

		pool1, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		b1 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool1.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		pool2, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 1)
		b2 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool2.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		// ASSERT
		Expect(b1 * 2).To(BeNumerically("~", b2, 1))
	})

	PIt("pools with zero operating cost should receive nothing", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), poolTypes.Pool{
			Name:           "PoolTest2",
			MaxBundleSize:  100,
			StartKey:       "0",
			UploadInterval: 60,
			OperatingCost:  0,
			Protocol: &poolTypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &poolTypes.UpgradePlan{},
		})

		// ACT
		s.Commit()

		pool1, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		b1 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool1.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		pool2, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 1)
		b2 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool2.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		// ASSERT
		Expect(b1).To(BeNumerically(">", b2))
		Expect(b2).To(BeZero())
	})

	It("pools with zero operating cost should receive nothing", func() {
		// ARRANGE
		s.App().PoolKeeper.SetPool(s.Ctx(), poolTypes.Pool{
			Id:             0,
			Name:           "PoolTest",
			MaxBundleSize:  100,
			StartKey:       "0",
			UploadInterval: 60,
			OperatingCost:  0,
			Protocol: &poolTypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &poolTypes.UpgradePlan{},
		})

		s.App().PoolKeeper.AppendPool(s.Ctx(), poolTypes.Pool{
			Name:           "PoolTest2",
			MaxBundleSize:  100,
			StartKey:       "0",
			UploadInterval: 60,
			OperatingCost:  0,
			Protocol: &poolTypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &poolTypes.UpgradePlan{},
		})

		// ACT
		s.Commit()

		pool1, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		b1 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool1.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		pool2, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 1)
		b2 := uint64(s.App().BankKeeper.GetBalance(s.Ctx(), pool2.GetPoolAccount(), globalTypes.Denom).Amount.Int64())

		// ASSERT
		Expect(b1).To(BeZero())
		Expect(b2).To(BeZero())
	})
})