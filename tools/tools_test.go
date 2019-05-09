package tools

import (
	"testing"
	"fmt"
)

func TestTransferredToHtml(t *testing.T) {
	fmt.Println(TransferredToHtml(" + \"select * from users;"))
}
