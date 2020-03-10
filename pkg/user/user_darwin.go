// +build darwin

package user

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ListUsers() (*services.ListUsersResponse, error) {
	out, err := exec.Command("dscacheutil", "-q", "user").Output()
	if err != nil {
		return nil, err
	}

	users := &services.ListUsersResponse{Users: []*resources.User{}}
	u := &resources.User{}
	i := 0
	for _, l := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(l, ": ", 2)
		switch kv[0] {
		case "name":
			u.Username = kv[1]
			i++
		case "password":
			i++
			continue
		case "uid":
			signed, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.UidSigned = signed
			i++
		case "gid":
			signed, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.GidSigned = signed
			i++
		case "dir":
			u.Directory = kv[1]
			i++
		case "shell":
			u.Shell = kv[1]
			i++
		case "gecos":
			u.Name = kv[1]
			i++
		case "":
			if i >= 7 {
				users.Users = append(users.Users, u)
			}
			u = &resources.User{}
			i = 0
		default:
			return nil, fmt.Errorf("unknown key: %s", kv[0])
		}
	}

	return users, nil
}
