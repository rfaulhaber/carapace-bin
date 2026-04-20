package zfs

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPools completes ZFS pool names
//
//	rpool (ONLINE)
//	tank (DEGRADED)
func ActionPools() carapace.Action {
	return carapace.ActionExecCommand("zpool", "list", "-H", "-o", "name,health")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) >= 2 {
				name := fields[0]
				health := fields[1]
				var s string
				switch health {
				case "ONLINE":
					s = style.Carapace.KeywordPositive
				case "DEGRADED":
					s = style.Carapace.KeywordAmbiguous
				case "FAULTED", "OFFLINE", "UNAVAIL", "REMOVED":
					s = style.Carapace.KeywordNegative
				default:
					s = styles.Zfs.Pool
				}
				vals = append(vals, name, health, s)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("pools")
}

// ActionPoolDevices completes device names within a pool
//
//	sda1
//	sdb1
func ActionPoolDevices(pool string) carapace.Action {
	return carapace.ActionExecCommand("zpool", "status", pool)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		re := regexp.MustCompile(`^\t  +(\S+)\s+(ONLINE|DEGRADED|FAULTED|OFFLINE|UNAVAIL|REMOVED)`)

		for _, line := range lines {
			if matches := re.FindStringSubmatch(line); len(matches) >= 3 {
				device := matches[1]
				if device == pool {
					continue
				}
				// Skip vdev types
				switch device {
				case "mirror", "raidz1", "raidz2", "raidz3", "draid", "spare", "log", "cache", "special", "dedup":
					continue
				}
				state := matches[2]
				vals = append(vals, device, state)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("devices")
}

// ActionVdevTypes completes ZFS vdev type keywords
//
//	mirror
//	raidz1
func ActionVdevTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"mirror", "mirror vdev",
		"raidz", "single-parity RAID-Z",
		"raidz1", "single-parity RAID-Z",
		"raidz2", "double-parity RAID-Z",
		"raidz3", "triple-parity RAID-Z",
		"draid", "distributed RAID",
		"spare", "hot spare",
		"log", "ZFS Intent Log",
		"cache", "L2ARC cache device",
		"special", "special allocation class",
		"dedup", "dedup allocation class",
	)
}
