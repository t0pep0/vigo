package main

import (
	"github.com/nsf/termbox-go"

//	"github.com/nsf/tulib"
//	"strconv"
)

type NormalMode struct {
	stub_overlay_mode
	editor *editor
}

func NewNormalMode(editor *editor) NormalMode {
	m := NormalMode{editor: editor}
	m.editor.SetStatus("Normal")
	return m
}

func (m NormalMode) OnKey(ev *termbox.Event) {
	g := m.editor
	v := g.active.leaf

	switch ev.Ch {
	case 'a':
		v.on_vcommand(vcommand_move_cursor_forward, 0)
		g.setMode(NewInsertMode(g))
	case 'A':
		v.on_vcommand(vcommand_move_cursor_end_of_line, 0)
		g.setMode(NewInsertMode(g))
	case 'h':
		v.on_vcommand(vcommand_move_cursor_backward, 0)
	case 'i':
		g.setMode(NewInsertMode(g))
	case 'j':
		v.on_vcommand(vcommand_move_cursor_next_line, 0)
	case 'k':
		v.on_vcommand(vcommand_move_cursor_prev_line, 0)
	case 'l':
		v.on_vcommand(vcommand_move_cursor_forward, 0)
	case 'w':
		v.on_vcommand(vcommand_move_cursor_word_forward, 0)
	case 'b':
		v.on_vcommand(vcommand_move_cursor_word_backward, 0)
	case 'x':
		v.on_vcommand(vcommand_delete_rune, 0)
	case '0':
		v.on_vcommand(vcommand_move_cursor_beginning_of_line, 0)
	case '$':
		v.on_vcommand(vcommand_move_cursor_end_of_line, 0)
	case ':':
		g.setMode(NewCommandMode(g, m))
	}
}

func (m NormalMode) Exit() {
}
