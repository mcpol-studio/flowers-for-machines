package mod_event_client_to_server

import (
	"github.com/OmineDev/flowers-for-machines/core/py_rpc/mod_event/client_to_server/minecraft"
	mei "github.com/OmineDev/flowers-for-machines/core/py_rpc/mod_event/interface"
)

// Minecraft Package
type Minecraft struct{ mei.Default }

// Return the package name of m
func (m *Minecraft) PackageName() string {
	return "Minecraft"
}

// Return a pool/map that contains all the module of m
func (m *Minecraft) ModulePool() map[string]mei.Module {
	return map[string]mei.Module{
		"vipEventSystem": &minecraft.VIPEventSystem{Module: &mei.DefaultModule{}},
		"preset":         &minecraft.Preset{Module: &mei.DefaultModule{}},
		"aiCommand":      &minecraft.AICommand{Module: &mei.DefaultModule{}},
	}
}
