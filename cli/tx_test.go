package cli

import "testing"

func TestTx(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("submit 5 -a 0xeF0b04a14e62434a99C4aF28C6dAb52ba9B1C8F3 -f 0xDC8F76075Db000Fa70fdA3AA2c95d63F22A10a67 -t 0x6a038842f9E9010624eAeB5f30ec5004C05EE21D")
	cli.TestCommand("confirm 6")
	cli.TestCommand("revoke 7")
	cli.TestCommand("execute 8")
	cli.TestCommand("check 9")

	cli.TestCommand("info 9")
	cli.TestCommand("list")
}
