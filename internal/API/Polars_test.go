/*
 * Minecraft Server Tool(MCST) is a command-line utility making Minecraft server creation quick and easy for beginners.
 * Copyright (c) 2024-2024 Arama.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package api_test

import (
	"encoding/json"
	"testing"

	api "github.com/Arama0517/MCST/internal/API"
	"github.com/Arama0517/MCST/internal/configs"
)

func TestPolars(t *testing.T) {
	if err := configs.InitData(); err != nil {
		t.Fatal(err)
	}
	data, err := api.GetPolarsData()
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonData))
}

func TestPolarsCore(t *testing.T) {
	if err := configs.InitData(); err != nil {
		t.Fatal(err)
	}
	data, err := api.GetPolarsCoresData(16)
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonData))
}
