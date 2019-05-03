package vclib

import (
 "github.com/vmware/govmomi/object"
)

type VMOptions struct {
 VMFolder       *Folder
 VMResourcePool *object.ResourcePool
}
