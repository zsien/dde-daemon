/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package bluetooth

import (
	"dbus/org/bluez"
	"fmt"
	"pkg.deepin.io/lib/dbus"
	"time"
)

type adapter struct {
	bluezAdapter *bluez.Adapter1

	Path                dbus.ObjectPath
	adddress            string
	Name                string
	Alias               string
	Powered             bool
	Discovering         bool
	Discoverable        bool
	DiscoverableTimeout uint32
}

func newAdapter(apath dbus.ObjectPath) (a *adapter) {
	a = &adapter{Path: apath}
	a.bluezAdapter, _ = bluezNewAdapter(apath)
	a.connectProperties()
	a.adddress = a.bluezAdapter.Address.Get()
	a.Alias = a.bluezAdapter.Alias.Get()
	a.Name = a.bluezAdapter.Name.Get()
	a.Powered = a.bluezAdapter.Powered.Get()
	a.Discovering = a.bluezAdapter.Discovering.Get()
	a.Discoverable = a.bluezAdapter.Discoverable.Get()
	a.DiscoverableTimeout = a.bluezAdapter.DiscoverableTimeout.Get()
	return
}
func destroyAdapter(a *adapter) {
	bluezDestroyAdapter(a.bluezAdapter)
}

func (a *adapter) notifyAdapterAdded() {
	logger.Info("AdapterAdded", marshalJSON(a))
	dbus.Emit(bluetooth, "AdapterAdded", marshalJSON(a))
	bluetooth.setPropState()
}
func (a *adapter) notifyAdapterRemoved() {
	logger.Info("AdapterRemoved", marshalJSON(a))
	dbus.Emit(bluetooth, "AdapterRemoved", marshalJSON(a))
	bluetooth.setPropState()
}
func (a *adapter) notifyPropertiesChanged() {
	logger.Debug("AdapterPropertiesChanged", marshalJSON(a))
	dbus.Emit(bluetooth, "AdapterPropertiesChanged", marshalJSON(a))
	bluetooth.setPropState()
}
func (a *adapter) connectProperties() {
	a.bluezAdapter.Name.ConnectChanged(func() {
		a.Name = a.bluezAdapter.Name.Get()
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
	a.bluezAdapter.Alias.ConnectChanged(func() {
		a.Alias = a.bluezAdapter.Alias.Get()
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
	a.bluezAdapter.Powered.ConnectChanged(func() {
		a.Powered = a.bluezAdapter.Powered.Get()
		logger.Infof("adapter powered changed %#v", a)
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
	a.bluezAdapter.Discovering.ConnectChanged(func() {
		a.Discovering = a.bluezAdapter.Discovering.Get()
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
	a.bluezAdapter.Discoverable.ConnectChanged(func() {
		a.Discoverable = a.bluezAdapter.Discoverable.Get()
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
	a.bluezAdapter.DiscoverableTimeout.ConnectChanged(func() {
		a.DiscoverableTimeout = a.bluezAdapter.DiscoverableTimeout.Get()
		a.notifyPropertiesChanged()
		bluetooth.setPropAdapters()
	})
}

func (b *Bluetooth) addAdapter(apath dbus.ObjectPath) {
	if b.isAdapterExists(apath) {
		logger.Error("repeat add adapter", apath)
		return
	}

	// initialize adapter power state
	b.config.addAdapterConfig(bluezGetAdapterAddress(apath))
	oldPowered := b.config.getAdapterConfigPowered(bluezGetAdapterAddress(apath))
	b.SetAdapterPowered(apath, oldPowered)
	b.SetAdapterDiscoverable(apath, false)
	if oldPowered {
		b.RequestDiscovery(apath)
	}

	b.adaptersLock.Lock()
	defer b.adaptersLock.Unlock()
	a := newAdapter(apath)
	b.adapters = append(b.adapters, a)
	a.notifyAdapterAdded()
	b.setPropAdapters()
}
func (b *Bluetooth) removeAdapter(apath dbus.ObjectPath) {
	i := b.getAdapterIndex(apath)
	if i < 0 {
		logger.Error("repeat remove adapter", apath)
		return
	}

	b.adaptersLock.Lock()
	defer b.adaptersLock.Unlock()
	b.doRemoveAdapter(i)
	b.setPropAdapters()
}
func (b *Bluetooth) doRemoveAdapter(i int) {
	removeAdapter := b.adapters[i]
	copy(b.adapters[i:], b.adapters[i+1:])
	b.adapters[len(b.adapters)-1] = nil
	b.adapters = b.adapters[:len(b.adapters)-1]

	removeAdapter.notifyAdapterRemoved()
	destroyAdapter(removeAdapter)
}

func (b *Bluetooth) getAdapter(apath dbus.ObjectPath) (a *adapter, err error) {
	i := b.getAdapterIndex(apath)
	if i < 0 {
		err = fmt.Errorf("adapter not exists %s", apath)
		logger.Error(err)
		return
	}

	b.adaptersLock.Lock()
	defer b.adaptersLock.Unlock()
	a = b.adapters[i]
	return
}
func (b *Bluetooth) isAdapterExists(apath dbus.ObjectPath) bool {
	if b.getAdapterIndex(apath) >= 0 {
		return true
	}
	return false
}
func (b *Bluetooth) getAdapterIndex(apath dbus.ObjectPath) int {
	b.adaptersLock.Lock()
	defer b.adaptersLock.Unlock()
	for i, a := range b.adapters {
		if a.Path == apath {
			return i
		}
	}
	return -1
}

// GetAdapters return all adapter objects that marshaled by json.
func (b *Bluetooth) GetAdapters() (adaptersJSON string, err error) {
	adaptersJSON = marshalJSON(b.adapters)
	return
}

func (b *Bluetooth) RequestDiscovery(apath dbus.ObjectPath) (err error) {
	// if adapter is discovering now, just ignore
	if bluezGetAdapterDiscovering(apath) {
		return
	}

	err = bluezStartDiscovery(apath)
	go func() {
		// adapter is not ready, retry again
		if err != nil {
			time.Sleep(3 * time.Second)
			err = bluezStartDiscovery(apath)
		}

		//	if err == nil {
		//		time.Sleep(20 * time.Second)
		//		if b.isAdapterExists(apath) {
		//			bluezStopDiscovery(apath)
		//		}
		//	}
	}()
	return
}

func (b *Bluetooth) SetAdapterPowered(apath dbus.ObjectPath, powered bool) (err error) {
	err = bluezSetAdapterPowered(apath, powered)
	if err == nil {
		// save the powered state
		b.config.setAdapterConfigPowered(bluezGetAdapterAddress(apath), powered)
	}
	return
}
func (b *Bluetooth) SetAdapterAlias(apath dbus.ObjectPath, alias string) (err error) {
	err = bluezSetAdapterAlias(apath, alias)
	return
}
func (b *Bluetooth) SetAdapterDiscoverable(apath dbus.ObjectPath, discoverable bool) (err error) {
	err = bluezSetAdapterDiscoverable(apath, discoverable)
	return
}
func (b *Bluetooth) SetAdapterDiscoverableTimeout(apath dbus.ObjectPath, discoverableTimeout uint32) (err error) {
	err = bluezSetAdapterDiscoverableTimeout(apath, discoverableTimeout)
	return
}
