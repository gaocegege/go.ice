// Code generated by gommon from ice/db/adapters/sqlite/gommon.yml DO NOT EDIT.
package sqlite

import dlog "github.com/dyweb/gommon/log"

func (a *Adapter) SetLogger(logger *dlog.Logger) {
	a.log = logger
}

func (a *Adapter) GetLogger() *dlog.Logger {
	return a.log
}

func (a *Adapter) LoggerIdentity(justCallMe func() *dlog.Identity) *dlog.Identity {
	return justCallMe()
}
