package marathon

import (
	"github.com/Broderick-Westrope/tetrigo/cmd/tui/common"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	ForceQuit        key.Binding
	Exit             key.Binding
	Help             key.Binding
	Left             key.Binding
	Right            key.Binding
	Clockwise        key.Binding
	CounterClockwise key.Binding
	SoftDrop         key.Binding
	HardDrop         key.Binding
	Hold             key.Binding
}

func constructKeyMap(keys *common.Keys) *keyMap {
	return &keyMap{
		ForceQuit:        common.ConstructKeyBinding(keys.ForceQuit, "force quit"),
		Exit:             common.ConstructKeyBinding(keys.Exit, "exit"),
		Help:             common.ConstructKeyBinding(keys.Help, "help"),
		Left:             common.ConstructKeyBinding(keys.Left, "move left"),
		Right:            common.ConstructKeyBinding(keys.Right, "move right"),
		Clockwise:        common.ConstructKeyBinding(keys.RightRotate, "rotate clockwise"),
		CounterClockwise: common.ConstructKeyBinding(keys.LeftRotate, "rotate counter-clockwise"),
		SoftDrop:         common.ConstructKeyBinding(keys.Down, "toggle soft drop"),
		HardDrop:         common.ConstructKeyBinding(keys.Up, "hard drop"),
		Hold:             common.ConstructKeyBinding(keys.Submit, "hold"),
	}
}

func (k *keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Exit,
		k.Help,
	}
}

func (k *keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Exit,
			k.Help,
			k.Left,
		},
		{
			k.Right,
			k.Clockwise,
			k.CounterClockwise,
		},
		{
			k.SoftDrop,
			k.HardDrop,
			k.Hold,
		},
	}
}
