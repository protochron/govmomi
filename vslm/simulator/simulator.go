package simulator

import (
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm/methods"
	"github.com/vmware/govmomi/vslm/types"
)

var content = types.VslmServiceInstanceContent{
	AboutInfo: types.VslmAboutInfo{
		Name:         "VSLM",
		InstanceUuid: "",
	},
	SessionManager:        vim.ManagedObjectReference{Type: "PbmSessionManager", Value: "SessionManager"},
	VStorageObjectManager: vim.ManagedObjectReference{Type: "VcenterVStorageObjectManager", Value: "VStorageObjectManager"},
}

func ServiceContent() types.VslmServiceInstanceContent {
	return content
}

type ServiceInstance struct {
	vim.ManagedObjectReference

	Content types.VslmServiceInstanceContent
}

func (s *ServiceInstance) RetrieveContent(_ *types.RetrieveContent) soap.HasFault {
	return &methods.RetrieveContentBody{
		Res: &types.RetrieveContentResponse{
			Returnval: s.Content,
		},
	}
}
