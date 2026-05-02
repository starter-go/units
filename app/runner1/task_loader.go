package runner1

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/units"
)

type innerTaskLoader struct {
	namelist []string
	tasks    []*innerTask
}

func (inst *innerTaskLoader) load(ulist []units.Unit, namelist string) ([]*innerTask, error) {

	inst.namelist = nil
	inst.tasks = nil
	var err error

	err = inst.innerLoadNameList(namelist)
	if err != nil {
		return nil, err
	}

	err = inst.innerLoadUnits(ulist)
	if err != nil {
		return nil, err
	}

	err = inst.innerSelect()
	if err != nil {
		return nil, err
	}

	inst.innerSort()

	return inst.tasks, nil
}

func (inst *innerTaskLoader) innerLoadUnits(src []units.Unit) error {

	inst.tasks = nil

	dst := inst.tasks
	tmp := make([]*units.Registration, 0)

	for _, it := range src {
		tmp = it.ListRegistrations(tmp)
	}

	for _, it1 := range tmp {
		if it1 == nil {
			continue
		}
		if it1.Do == nil {
			continue
		}
		it2 := new(innerTask)
		it2.Info = *it1
		it2.State = units.TaskStateInit
		dst = append(dst, it2)
	}

	inst.tasks = dst
	return nil
}

func (inst *innerTaskLoader) innerLoadNameList(namelistString string) error {

	src := namelistString
	tmp := strings.Split(src, ",")
	dst := []string{}

	for _, item := range tmp {
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}
		dst = append(dst, item)
	}

	inst.namelist = dst
	return nil
}

func (inst *innerTaskLoader) innerSort() {
	sort.Sort(inst)
}

func (inst *innerTaskLoader) innerPutTaskToTable(table map[string]*innerTask, name string, task *innerTask) error {

	if table == nil || task == nil {
		return nil
	}

	if name == "" {
		return nil
	}

	older := table[name]
	if older != nil {
		return fmt.Errorf("unit.name(alias): '%s' is duplicate", name)
	}

	table[name] = task
	return nil
}

func (inst *innerTaskLoader) innerSelect() error {

	table := make(map[string]*innerTask)
	namelist := inst.namelist
	tasks := inst.tasks

	for _, item := range tasks {

		info := &item.Info
		if info.Alias == info.Name {
			info.Alias = ""
		}

		err1 := inst.innerPutTaskToTable(table, info.Alias, item)
		err2 := inst.innerPutTaskToTable(table, info.Name, item)

		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
	}

	for _, name := range namelist {
		item := table[name]
		if item == nil {
			return fmt.Errorf("no test case with name (alias) of '%s'", name)
		}
		item.Selected = true
	}

	return nil
}

func (inst *innerTaskLoader) acceptItem(item *innerTask) bool {

	if item == nil {
		return false
	}

	info := &item.Info

	if info.Do == nil {
		return false
	}

	return info.Enabled
}

func (inst *innerTaskLoader) Len() int {
	return len(inst.tasks)
}
func (inst *innerTaskLoader) Less(i1, i2 int) bool {

	o1 := inst.tasks[i1]
	o2 := inst.tasks[i2]

	// by (index, priority)

	if o1.Index != o2.Index {
		return o1.Index > o2.Index
	}
	return o1.Info.Priority < o2.Info.Priority
}
func (inst *innerTaskLoader) Swap(i1, i2 int) {
	l := inst.tasks
	l[i1], l[i2] = l[i2], l[i1]
}
