// VulcanizeDB
// Copyright © 2018 Vulcanize

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

package shared

import (
	"fmt"
)

type ErrContractNotFound struct {
	Contract string
}

func (e ErrContractNotFound) Error() string {
	return fmt.Sprintf("transformer not found for contract: %s", e.Contract)
}

type ErrHeaderMismatch struct {
	BlockHeight int
	DbHash      string
	DiffHash    string
}

func (e ErrHeaderMismatch) Error() string {
	return fmt.Sprintf("header hash in row does not match db at height %d - row: %s, db: %s", e.BlockHeight, e.DbHash, e.DiffHash)
}

type ErrMetadataMalformed struct {
	MissingData Key
}

func (e ErrMetadataMalformed) Error() string {
	return fmt.Sprintf("storage metadata malformed: missing %s", e.MissingData)
}

type ErrRowMalformed struct {
	Length int
}

func (e ErrRowMalformed) Error() string {
	return fmt.Sprintf("storage row malformed: length %d, expected %d", e.Length, ExpectedRowLength)
}

type ErrStorageKeyNotFound struct {
	Key string
}

func (e ErrStorageKeyNotFound) Error() string {
	return fmt.Sprintf("unknown storage key: %s", e.Key)
}

type ErrTypeNotFound struct {
	Type int
}

func (e ErrTypeNotFound) Error() string {
	return fmt.Sprintf("no decoder for type: %d", e.Type)
}