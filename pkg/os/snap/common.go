package snap

import (
	"fmt"
	"github.com/linuxsuren/http-downloader/pkg/exec"
)

// SnapName is the name of the snap
const SnapName = "snap"

// CommonInstaller is the installer of a common snap
type CommonInstaller struct {
	Name   string
	Args   []string `yaml:"args"`
	Execer exec.Execer
}

// Available check if support current platform
func (d *CommonInstaller) Available() (ok bool) {
	if d.Execer.OS() == "linux" {
		_, err := d.Execer.LookPath(SnapName)
		ok = err == nil
	}
	return
}

// Install installs the target package
func (d *CommonInstaller) Install() (err error) {
	args := []string{"install", d.Name}
	args = append(args, d.Args...)
	if err = d.Execer.RunCommand(SnapName, args...); err != nil {
		return
	}
	return
}

// Uninstall uninstalls the target package
func (d *CommonInstaller) Uninstall() (err error) {
	err = d.Execer.RunCommand(SnapName, "remove", d.Name)
	return
}

// WaitForStart waits for the service be started
func (d *CommonInstaller) WaitForStart() (ok bool, err error) {
	ok = true
	return
}

// Start starts the target service
func (d *CommonInstaller) Start() error {
	fmt.Println("not supported yet")
	return nil
}

// Stop stops the target service
func (d *CommonInstaller) Stop() error {
	fmt.Println("not supported yet")
	return nil
}
