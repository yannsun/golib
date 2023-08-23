// +build linux

package runtime

import (
	"math"

	cg "selfergit.com/selfer/golib/automaxprocs/internal/cgroups"
)

// CPUQuotaToGOMAXPROCS converts the CPU quota applied to the calling process
// to a valid GOMAXPROCS value.
func CPUQuotaToGOMAXPROCS(minValue int) (int, CPUQuotaStatus, error) {
	cgroups, err := cg.NewCGroupsForCurrentProcess()
	if err != nil {
		return -1, CPUQuotaUndefined, err
	}

	quota, defined, err := cgroups.CPUQuota()
	if !defined || err != nil {
		return -1, CPUQuotaUndefined, err
	}

	maxProcs := int(math.Floor(quota))
	if minValue > 0 && maxProcs < minValue {
		return minValue, CPUQuotaMinUsed, nil
	}
	return maxProcs, CPUQuotaUsed, nil
}
