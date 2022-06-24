package main

import "github.com/bitfield/script"

func main() {
	script.Exec("bash -c 'ssh -T git@github.com'").Stdout()
}
