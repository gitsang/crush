package editor

import (
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/crush/internal/config"
)

type EditorKeyMap struct {
	AddFile     key.Binding
	SendMessage key.Binding
	OpenEditor  key.Binding
	Newline     key.Binding
}

func DefaultEditorKeyMap() EditorKeyMap {
	return EditorKeyMap{
		AddFile: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "add file"),
		),
		SendMessage: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "send"),
		),
		OpenEditor: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "open editor"),
		),
		Newline: key.NewBinding(
			key.WithKeys("shift+enter", "ctrl+j"),
			// "ctrl+j" is a common keybinding for newline in many editors. If
			// the terminal supports "shift+enter", we substitute the help text
			// to reflect that.
			key.WithHelp("ctrl+j", "newline"),
		),
	}
}

func NewEditorKeyMap(cfg *config.Config) EditorKeyMap {
	keyMap := DefaultEditorKeyMap()

	if cfg != nil && cfg.Options != nil && cfg.Options.TUI != nil {
		if cfg.Options.TUI.SendKey != "" {
			keyMap.SendMessage = key.NewBinding(
				key.WithKeys(cfg.Options.TUI.SendKey),
				key.WithHelp(cfg.Options.TUI.SendKey, "send"),
			)
		}
	}

	return keyMap
}

// KeyBindings implements layout.KeyMapProvider
func (k EditorKeyMap) KeyBindings() []key.Binding {
	return []key.Binding{
		k.AddFile,
		k.SendMessage,
		k.OpenEditor,
		k.Newline,
		AttachmentsKeyMaps.AttachmentDeleteMode,
		AttachmentsKeyMaps.DeleteAllAttachments,
		AttachmentsKeyMaps.Escape,
	}
}

type DeleteAttachmentKeyMaps struct {
	AttachmentDeleteMode key.Binding
	Escape               key.Binding
	DeleteAllAttachments key.Binding
}

// TODO: update this to use the new keymap concepts
var AttachmentsKeyMaps = DeleteAttachmentKeyMaps{
	AttachmentDeleteMode: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r+{i}", "delete attachment at index i"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc", "alt+esc"),
		key.WithHelp("esc", "cancel delete mode"),
	),
	DeleteAllAttachments: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("ctrl+r+r", "delete all attachments"),
	),
}
