package dockerContainer

type ContainerRemoveRequest struct {
	RemoveVolumes bool `json:"removeVolumes" validate:"required"`
	RemoveLinks   bool `json:"removeLinks" validate:"required"`
	Force         bool `json:"force" validate:"required"`
}

type ContainerRenameRequest struct {
	NewName string `json:"newName" validate:"required"`
}
