package opts

import (
	"fmt"
	"testing"
)

func TestO(t *testing.T) {
	opts := NewOpts().Where("id = ?", 1).Where("name = ?", "张三")
	opts = opts.Where("", "")
	fmt.Println(len(opts))
}
