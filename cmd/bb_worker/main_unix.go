// +build darwin freebsd linux

package main

import (
	"path/filepath"
	"syscall"

	"github.com/buildbarn/bb-storage/pkg/util"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func clearUmask() {
	syscall.Umask(0)
}

func getInputRootCharacterDevices(names []string) (map[string]int, error) {
	inputRootCharacterDevices := map[string]int{}
	for _, device := range names {
		var stat unix.Stat_t
		devicePath := filepath.Join("/dev", device)
		if err := unix.Stat(devicePath, &stat); err != nil {
			return nil, util.StatusWrapf(err, "Unable to stat character device %#v", devicePath)
		}
		if stat.Mode&syscall.S_IFMT != syscall.S_IFCHR {
			return nil, status.Errorf(codes.InvalidArgument, "The specified device %#v is not a character device", devicePath)
		}
		inputRootCharacterDevices[device] = int(stat.Rdev)
	}
	return inputRootCharacterDevices, nil
}
