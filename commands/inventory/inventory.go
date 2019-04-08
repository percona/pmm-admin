// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package inventory provides inventory commands.
package inventory

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// register command
var (
	InventoryC    = kingpin.Command("inventory", "Inventory commands.")
	InventoryAddC = InventoryC.Command("add", "Add to inventory commands.")
)
