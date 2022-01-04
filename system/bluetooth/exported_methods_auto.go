// Code generated by "dbusutil-gen em -type SysBluetooth,agent"; DO NOT EDIT.

package bluetooth

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (v *SysBluetooth) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name: "ClearUnpairedDevice",
			Fn:   v.ClearUnpairedDevice,
		},
		{
			Name:   "ConnectDevice",
			Fn:     v.ConnectDevice,
			InArgs: []string{"devPath", "adapterPath"},
		},
		{
			Name:    "DebugInfo",
			Fn:      v.DebugInfo,
			OutArgs: []string{"info"},
		},
		{
			Name: "DisconnectAudioDevices",
			Fn:   v.DisconnectAudioDevices,
		},
		{
			Name:   "DisconnectDevice",
			Fn:     v.DisconnectDevice,
			InArgs: []string{"devPath"},
		},
		{
			Name:    "GetAdapters",
			Fn:      v.GetAdapters,
			OutArgs: []string{"adaptersJSON"},
		},
		{
			Name:    "GetDevices",
			Fn:      v.GetDevices,
			InArgs:  []string{"adapterPath"},
			OutArgs: []string{"devicesJSON"},
		},
		{
			Name:   "RegisterAgent",
			Fn:     v.RegisterAgent,
			InArgs: []string{"agentPath"},
		},
		{
			Name:   "RemoveDevice",
			Fn:     v.RemoveDevice,
			InArgs: []string{"adapterPath", "devPath"},
		},
		{
			Name:   "RequestDiscovery",
			Fn:     v.RequestDiscovery,
			InArgs: []string{"adapterPath"},
		},
		{
			Name:   "SetAdapterAlias",
			Fn:     v.SetAdapterAlias,
			InArgs: []string{"adapterPath", "alias"},
		},
		{
			Name:   "SetAdapterDiscoverable",
			Fn:     v.SetAdapterDiscoverable,
			InArgs: []string{"adapterPath", "discoverable"},
		},
		{
			Name:   "SetAdapterDiscoverableTimeout",
			Fn:     v.SetAdapterDiscoverableTimeout,
			InArgs: []string{"adapterPath", "discoverableTimeout"},
		},
		{
			Name:   "SetAdapterDiscovering",
			Fn:     v.SetAdapterDiscovering,
			InArgs: []string{"adapterPath", "discovering"},
		},
		{
			Name:   "SetAdapterPowered",
			Fn:     v.SetAdapterPowered,
			InArgs: []string{"adapterPath", "powered"},
		},
		{
			Name:   "SetDeviceAlias",
			Fn:     v.SetDeviceAlias,
			InArgs: []string{"device", "alias"},
		},
		{
			Name:   "SetDeviceTrusted",
			Fn:     v.SetDeviceTrusted,
			InArgs: []string{"device", "trusted"},
		},
		{
			Name:   "UnregisterAgent",
			Fn:     v.UnregisterAgent,
			InArgs: []string{"agentPath"},
		},
	}
}
func (v *agent) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:   "AuthorizeService",
			Fn:     v.AuthorizeService,
			InArgs: []string{"device", "uuid"},
		},
		{
			Name: "Cancel",
			Fn:   v.Cancel,
		},
		{
			Name:   "DisplayPasskey",
			Fn:     v.DisplayPasskey,
			InArgs: []string{"device", "passkey", "entered"},
		},
		{
			Name:   "DisplayPinCode",
			Fn:     v.DisplayPinCode,
			InArgs: []string{"device", "pinCode"},
		},
		{
			Name: "Release",
			Fn:   v.Release,
		},
		{
			Name:   "RequestAuthorization",
			Fn:     v.RequestAuthorization,
			InArgs: []string{"device"},
		},
		{
			Name:   "RequestConfirmation",
			Fn:     v.RequestConfirmation,
			InArgs: []string{"device", "passkey"},
		},
		{
			Name:    "RequestPasskey",
			Fn:      v.RequestPasskey,
			InArgs:  []string{"device"},
			OutArgs: []string{"passkey"},
		},
		{
			Name:    "RequestPinCode",
			Fn:      v.RequestPinCode,
			InArgs:  []string{"device"},
			OutArgs: []string{"pinCode"},
		},
	}
}
