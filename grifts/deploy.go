package grifts

import (
	"os"
	"os/exec"

	. "github.com/markbates/grift/grift"
)

var _ = Add("deploy", func(c *Context) error {
	return run(exec.Command("ssh", "root@gopheracademy.com", "'/root/go/src/github.com/gopheracademy/gcon/deploy.sh'"))

})

var _ = Add("remote", func(c *Context) error {
	return run(exec.Command("ssh", "root@gopheracademy.com", "'/root/go/src/github.com/gopheracademy/gcon/grift.sh'", os.Args[2]))
	return nil

})

func run(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
