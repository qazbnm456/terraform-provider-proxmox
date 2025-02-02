/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vms

import (
	"testing"

	"github.com/stretchr/testify/require"

	types2 "github.com/bpg/terraform-provider-proxmox/proxmox/types"
)

func TestCustomStorageDevice_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	ds8gig := types2.DiskSizeFromGigabytes(8)
	tests := []struct {
		name    string
		line    string
		want    *CustomStorageDevice
		wantErr bool
	}{
		{
			name: "simple volume",
			line: `"local-lvm:vm-2041-disk-0,discard=on,ssd=1,iothread=1,size=8G,cache=writeback"`,
			want: &CustomStorageDevice{
				Cache:      types2.StrPtr("writeback"),
				Discard:    types2.StrPtr("on"),
				Enabled:    true,
				FileVolume: "local-lvm:vm-2041-disk-0",
				IOThread:   types2.BoolPtr(true),
				Size:       &ds8gig,
				SSD:        types2.BoolPtr(true),
			},
		},
		{
			name: "raw volume type",
			line: `"nfs:2041/vm-2041-disk-0.raw,discard=ignore,ssd=1,iothread=1,size=8G"`,
			want: &CustomStorageDevice{
				Discard:    types2.StrPtr("ignore"),
				Enabled:    true,
				FileVolume: "nfs:2041/vm-2041-disk-0.raw",
				Format:     types2.StrPtr("raw"),
				IOThread:   types2.BoolPtr(true),
				Size:       &ds8gig,
				SSD:        types2.BoolPtr(true),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &CustomStorageDevice{}
			if err := r.UnmarshalJSON([]byte(tt.line)); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, r)
		})
	}
}

func TestCustomPCIDevice_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		line    string
		want    *CustomPCIDevice
		wantErr bool
	}{
		{
			name: "id only pci device",
			line: `"0000:81:00.2"`,
			want: &CustomPCIDevice{
				DeviceIDs:  &[]string{"0000:81:00.2"},
				MDev:       nil,
				PCIExpress: types2.BoolPtr(false),
				ROMBAR:     types2.BoolPtr(true),
				ROMFile:    nil,
				XVGA:       types2.BoolPtr(false),
			},
		},
		{
			name: "pci device with more details",
			line: `"host=81:00.4,pcie=0,rombar=1,x-vga=0"`,
			want: &CustomPCIDevice{
				DeviceIDs:  &[]string{"81:00.4"},
				MDev:       nil,
				PCIExpress: types2.BoolPtr(false),
				ROMBAR:     types2.BoolPtr(true),
				ROMFile:    nil,
				XVGA:       types2.BoolPtr(false),
			},
		},
		{
			name: "pci device with mapping",
			line: `"mapping=mappeddevice,pcie=0,rombar=1,x-vga=0"`,
			want: &CustomPCIDevice{
				DeviceIDs:  nil,
				Mapping:    types2.StrPtr("mappeddevice"),
				MDev:       nil,
				PCIExpress: types2.BoolPtr(false),
				ROMBAR:     types2.BoolPtr(true),
				ROMFile:    nil,
				XVGA:       types2.BoolPtr(false),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &CustomPCIDevice{}
			if err := r.UnmarshalJSON([]byte(tt.line)); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
