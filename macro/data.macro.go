package macro

const (
	MOVE_MOUSE = iota
	CLICK_MOUSE_L
	CLICK_MOUSE_R
	DOWN_MOUSE_L
	DOWN_MOUSE_R
	UP_MOUSE_L
	UP_MOUSE_R
	DOWN_KEY
	UP_KEY
)

type MacroAct struct {
	MacroType  int
	Mouse_xPos int
	Mouse_yPos int
	Key        string
}

var MacroActs []MacroAct

func MakeAndAppendMacro(mro_T int, xPos int, yPos int, key string){
	act := MacroAct{mro_T, xPos, yPos, key}
	MacroActs = append(MacroActs, act)
}