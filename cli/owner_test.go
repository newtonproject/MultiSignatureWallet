package cli

import "testing"

func TestOwner(t *testing.T) {
	cli := NewCLI()

	cli.TestCommand("owner list")
	cli.TestCommand("owner add 0xF67cD3b82491cB41aaC69ab579670D1839006476")
	cli.TestCommand("owner remove 0xF67cD3b82491cB41aaC69ab579670D1839006476")
	cli.TestCommand("owner replace 0xF67cD3b82491cB41aaC69ab579670D1839006476 0x536e9f4e54F2A32BB47F7223a7b621AFe509cCb2")
	cli.TestCommand("owner check 0xF67cD3b82491cB41aaC69ab579670D1839006476")
}
