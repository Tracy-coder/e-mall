package data

import "fmt"

const (
	RankKey = "rank"
)

func ProductViewKey(id uint64) string {
	return fmt.Sprintf("view:product:%s", id)
}
