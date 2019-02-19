// +build !windows

package fileutil

import (
	"log"
	"os"
	"os/user"
	"strconv"
)

// Take ownership of files, and then give them 0600 file permissions
func SecureFiles(filepaths []string) (err error) {
	log.Printf("sec: entry")
	var currentUser *user.User
	currentUser, err = user.Current()
	log.Printf("sec: p1")
	if err != nil {
		return err
	}
	var uid, gid int
	uid, err = strconv.Atoi(currentUser.Uid)
	log.Printf("sec: p2")
	if err != nil {
		return err
	}
	gid, err = strconv.Atoi(currentUser.Gid)
	log.Printf("sec: p3")
	if err != nil {
		return err
	}
	for _, path := range filepaths {
		err = os.Chown(
			path,
			uid,
			gid,
		)
		log.Printf("sec: p4")
		if err != nil {
			return err
		}
		err = os.Chmod(
			path,
			0600,
		)
		log.Printf("sec: p5")
		if err != nil {
			return err
		}
	}
	log.Printf("sec: p6")
	return err
}
