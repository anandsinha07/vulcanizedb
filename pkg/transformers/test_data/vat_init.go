// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test_data

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/transformers/vat_init"
)

var VatAddress = "0x239E6f0AB02713f1F8AA90ebeDeD9FC66Dc96CD6"

var EthVatInitLog = types.Log{
	Address: common.HexToAddress(VatAddress),
	Topics: []common.Hash{
		common.HexToHash("0x3b66319500000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x66616b6520696c6b000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000243b66319566616b6520696c6b000000000000000000000000000000000000000000000000"),
	BlockNumber: 24,
	TxHash:      common.HexToHash("0xe8f39fbb7fea3621f543868f19b1114e305aff6a063a30d32835ff1012526f91"),
	TxIndex:     7,
	BlockHash:   common.HexToHash("0xe3dd2e05bd8b92833e20ed83e2171bbc06a9ec823232eca1730a807bd8f5edc0"),
	Index:       0,
	Removed:     false,
}

var rawVatInitLog, _ = json.Marshal(EthVatInitLog)
var VatInitModel = vat_init.VatInitModel{
	Ilk:              "fake ilk",
	TransactionIndex: EthVatInitLog.TxIndex,
	Raw:              rawVatInitLog,
}
