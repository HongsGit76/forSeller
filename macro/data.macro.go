package macro

const (
	MOVE_MOUSE = iota // = 0
	CLICK_MOUSE_L     // = 1
	CLICK_MOUSE_R     // = 2
	DOWN_MOUSE_L      // = 3
	DOWN_MOUSE_R      // = 4
	UP_MOUSE_L        // = 5
	UP_MOUSE_R        // = 6
	DOWN_KEY          // = 7
	UP_KEY            // = 8
	SLEEP             // = 9
)

type MacroAct struct {
	MacroType  int
	Mouse_xPos int
	Mouse_yPos int
	Key        string
	SleepTime  int
}

var MacroActs []MacroAct

func MakeAndAppendMacro(mro_T int, xPos int, yPos int, key string, sleepTime int) {
	act := MacroAct{mro_T, xPos, yPos, key, sleepTime}
	MacroActs = append(MacroActs, act)
}