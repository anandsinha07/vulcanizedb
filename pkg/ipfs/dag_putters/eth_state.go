// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package dag_putters

import (
	"fmt"

	"github.com/vulcanize/vulcanizedb/pkg/ipfs"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/ipld"
)

type EthStateDagPutter struct {
	adder *ipfs.IPFS
}

func NewEthStateDagPutter(adder *ipfs.IPFS) *EthStateDagPutter {
	return &EthStateDagPutter{adder: adder}
}

func (erdp *EthStateDagPutter) DagPut(raw interface{}) ([]string, error) {
	stateNodeRLP, ok := raw.([]byte)
	if !ok {
		return nil, fmt.Errorf("EthStateDagPutter expected input type %T got %T", []byte{}, raw)
	}
	node, err := ipld.FromStateTrieRLP(stateNodeRLP)
	if err != nil {
		return nil, err
	}
	if err := erdp.adder.Add(node); err != nil {
		return nil, err
	}
	return []string{node.Cid().String()}, nil
}
