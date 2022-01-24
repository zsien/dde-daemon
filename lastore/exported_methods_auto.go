// Code generated by "dbusutil-gen em -type Lastore,Agent"; DO NOT EDIT.

package lastore

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (v *Agent) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:   "CloseNotification",
			Fn:     v.CloseNotification,
			InArgs: []string{"id"},
		},
		{
			Name:    "GetManualProxy",
			Fn:      v.GetManualProxy,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:   "ReportLog",
			Fn:     v.ReportLog,
			InArgs: []string{"msg"},
		},
		{
			Name:    "SendNotify",
			Fn:      v.SendNotify,
			InArgs:  []string{"appName", "replacesId", "appIcon", "summary", "body", "actions", "hints", "expireTimeout"},
			OutArgs: []string{"outArg0"},
		},
	}
}
func (v *Lastore) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "IsDiskSpaceSufficient",
			Fn:      v.IsDiskSpaceSufficient,
			OutArgs: []string{"result"},
		},
	}
}