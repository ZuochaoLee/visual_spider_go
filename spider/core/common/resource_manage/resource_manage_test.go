//
package resource_manage_test

import (
	"../resource_manage"
	"testing"
)

func TestResourceManage(t *testing.T) {
	var mc *resource_manage.ResourceManageChan
	mc = resource_manage.NewResourceManageChan(1)
	mc.GetOne()
	println("incr")
	mc.FreeOne()
	println("decr")
	mc.GetOne()
	println("incr")
}
