package example2units

import (
	"context"

	"github.com/starter-go/afs"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type CaseTryDirMan struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	DirMan units.DirManager //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryDirMan) ListRegistrations(list []*units.Registration) []*units.Registration {
	r1 := inst.getMyInfo()
	list = append(list, r1)
	return list
}

func (inst *CaseTryDirMan) getMyInfo() *units.Registration {
	return &units.Registration{
		Name:     "case-try-dir-manager",
		Enabled:  true,
		Priority: 0,
		Do:       inst.run,
	}
}

func (inst *CaseTryDirMan) run(ctx context.Context) error {

	// info := inst.getMyInfo()

	holder := &units.DirHolder{
		// Unit:    info,

		Context: ctx,
		Scope:   units.DirScopeRuntime,
		Key:     units.DirKeyOutput,
	}

	holder, err := inst.DirMan.GetDir(holder)
	if err != nil {
		return err
	}

	dir := holder.Path
	vlog.Info("  units.dir_man.path = %s", dir.GetPath())

	// mkdir

	om := new(afs.OptionsMaker)
	om.SetMode(7, 5, 5)
	opt := om.Options()

	if !dir.Exists() {
		err := dir.Mkdirs(&opt)
		if err != nil {
			return err
		}
	}

	//create file

	file := dir.GetChild("demo.txt")
	text := "hello, units_dir_man!\n"

	om.WriteOnly().Create().Append()
	om.SetMode(6, 4, 4)

	opt = om.Options()
	err = file.GetIO().WriteText(text, &opt)
	if err != nil {
		return err
	}

	return nil
}

func (inst *CaseTryDirMan) _impl() units.Unit {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
