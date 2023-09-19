// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package localnet_test

import (
	"context"
	"fmt"
	"math/big"

	localnet "pkg.berachain.dev/polaris/e2e/localnet/network"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainerizedNode", func() {
	var c localnet.ContainerizedNode

	BeforeEach(func() {
		var err error
		c, err = localnet.NewContainerizedNode(
			"polard/localnet",
			"v0.0.0",
			"goodcontainer",
			"8545/tcp",
			"8546/tcp",
			[]string{
				"GO_VERSION=1.21.1",
				"BASE_IMAGE=polard/base:v0.0.0",
			},
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	AfterEach(func() {
		if !CurrentSpecReport().Failure.IsZero() {
			logs, err := c.DumpLogs()
			Expect(err).ToNot(HaveOccurred())
			GinkgoWriter.Println(logs)
		}
		Expect(c.Stop()).To(Succeed())
		Expect(c.Remove()).To(Succeed())
	})

	It("should wait for a certain block height", func() {
		Expect(c.WaitForBlock(1)).To(MatchError("block height already passed"))
		Expect(c.WaitForBlock(12)).To(Succeed())
		Expect(c.WaitForNextBlock()).To(Succeed())
	})

	It("should get recent blocks with websockets", func() {
		wsclient := c.EthWsClient()
		headers := make(chan *coretypes.Header)
		sub, err := wsclient.SubscribeNewHead(context.Background(), headers)
		Expect(err).ToNot(HaveOccurred())
		GinkgoWriter.Println("Listening for blocks...")
		select {
		case err = <-sub.Err():
			Fail(fmt.Sprintf("Error in subscription for recent blocks: %v", err))
		case header := <-headers:
			GinkgoWriter.Printf("New block: %v", header.Number.Uint64())
			_, err = wsclient.BlockByNumber(
				context.Background(), big.NewInt(header.Number.Int64()),
			)
			Expect(err).ToNot(HaveOccurred())
		}
	})
})
