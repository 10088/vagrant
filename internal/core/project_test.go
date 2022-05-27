package core

import (
	"fmt"
	"testing"

	"github.com/hashicorp/vagrant/internal/server/proto/vagrant_server"
	"github.com/stretchr/testify/require"
)

func projectTargets(t *testing.T, project *Project, numTargets int) (targets []*Target) {
	targets = make([]*Target, numTargets)
	for i := 0; i < numTargets; i++ {
		tt, err := TestTarget(t, project, &vagrant_server.Target{
			ResourceId: fmt.Sprintf("id-%d", i),
			Name:       fmt.Sprintf("target-%d", i),
			Uuid:       fmt.Sprintf("uuid-%d", i),
		})
		if err != nil {
			t.Error(err)
		}
		targets = append(targets, tt)
	}
	return
}

func TestNewProject(t *testing.T) {
	tp := TestMinimalProject(t)
	vn := tp.Ref()
	if vn == nil {
		t.Errorf("Creating project failed")
	}
}

func TestProjectGetTarget(t *testing.T) {
	tp := TestMinimalProject(t)
	// Add targets to project
	targetOne, err := TestTarget(t, tp, &vagrant_server.Target{ResourceId: "id-one", Name: "target-one"})
	if err != nil {
		t.Error(err)
	}
	targetTwo, err := TestTarget(t, tp, &vagrant_server.Target{ResourceId: "id-two", Name: "target-two"})
	if err != nil {
		t.Error(err)
	}

	// Get by id
	one, err := tp.Target("id-one", "")
	require.NoError(t, err)
	require.Equal(t, targetOne, one)

	// Get by name
	two, err := tp.Target("target-two", "")
	require.NoError(t, err)
	require.Equal(t, targetTwo, two)

	// Get target that doesn't exist
	noexist, err := tp.Target("ohnooooo", "")
	require.Error(t, err)
	require.Nil(t, noexist)
}

func TestProjectGetTargetNames(t *testing.T) {
	tp := TestMinimalProject(t)

	// No targets added
	names, err := tp.TargetNames()
	require.NoError(t, err)
	require.Len(t, names, 0)

	// Add targets to project
	projectTargets(t, tp, 3)

	names, err = tp.TargetNames()
	require.NoError(t, err)
	require.Len(t, names, 3)
	require.Contains(t, names, "target-0")
	require.Contains(t, names, "target-1")
	require.Contains(t, names, "target-2")
}

func TestProjectGetTargetIds(t *testing.T) {
	tp := TestMinimalProject(t)

	// No targets added
	ids, err := tp.TargetIds()
	require.NoError(t, err)
	require.Len(t, ids, 0)

	// Add targets to project
	projectTargets(t, tp, 3)

	ids, err = tp.TargetIds()
	require.NoError(t, err)
	require.Len(t, ids, 3)
	require.Contains(t, ids, "id-0")
	require.Contains(t, ids, "id-1")
	require.Contains(t, ids, "id-2")
}

func TestProjectGetTargets(t *testing.T) {
	tp := TestMinimalProject(t)

	// No targets added
	targets, err := tp.Targets()
	require.NoError(t, err)
	require.Len(t, targets, 0)

	// Add targets to project
	projectTargets(t, tp, 3)

	targets, err = tp.Targets()
	require.NoError(t, err)
	require.Len(t, targets, 3)
}
