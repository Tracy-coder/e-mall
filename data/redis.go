package data

import "fmt"

func ProductViewKey(id uint64) string {
	return fmt.Sprintf("view:product:%s", id)
}
