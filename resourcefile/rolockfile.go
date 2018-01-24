package resourcefile

import (
	"fmt"
	"strings"

	"github.com/autopilothq/banks/resourcefile/entities"
)

// ROLockfile is a read-only representation of the .banks/lock
// file for this project. You can use it to intropect metainfo
// about the resources that are in use.
type ROLockfile struct {
	resources       entities.Map
	resourcesByName map[string]*entities.Resource
}

// NewROLockfile returns a empty, read-only lockfile
func NewROLockfile() *ROLockfile {
	return &ROLockfile{
		resources:       make(entities.Map),
		resourcesByName: make(map[string]*entities.Resource),
	}
}

// GetROLockfile returns a read-only lockfile with the resources
// for this project loaded into it.
func GetROLockfile() (*ROLockfile, error) {
	lockfile := NewROLockfile()
	return lockfile, lockfile.restore()
}

// restore loads the resources from the resource data that is
// compiled into the final binary.
func (l *ROLockfile) restore() error {
	l.resources = make(entities.Map)

	raw, err := Asset(".banks/lock")
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(string(raw))) == 0 {
		// it's either empty or full of whitespace. Don't bother parsing it
		return nil
	}

	resources := make(entities.Collection, 0)
	if err := resources.UnmarshalJSON(raw); err != nil {
		return err
	}

	for _, resource := range resources {
		l.resources[resource.ID] = resource
		l.resourcesByName[resource.Name()] = resource
	}

	return nil
}

// Resources is a getter provided to banks for accessing declared resources
func (l *ROLockfile) Resources() (entities.Map, error) {
	return l.resources, nil
}

// Resource is a getter provided to banks for accessing one declared resource
func (l *ROLockfile) Resource(name string) (*entities.Resource, error) {
	decl, exists := l.resourcesByName[name]
	if !exists {
		return nil, fmt.Errorf("There was no resource named `%s`", name)
	}
	return decl, nil
}
