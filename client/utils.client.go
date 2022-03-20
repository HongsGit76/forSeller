package client

type windowSize struct {
	width  float32
	height float32
}

var WindowSize = windowSize{450, 450}
var ButtonSize = windowSize{20, 100}

const DefaultModeString = "아래 버튼 중 수행할 매크로를 선택하세요."
const KeyboardModeString = "키보드 매크로를 입력하세요."
const MouseModeString = "(F7)을 눌러 현재 마우스 위치를 복사합니다."
const StopString = "매크로를 멈추려면 \"s\"를 눌러 주세요."
const NoMacroListString = "매크로를 먼저 생성하고 실행해 주세요."

const LMouseDown = "   왼쪽 마우스 누르기  "
const RMouseDown = "   오른쪽 마우스 누르기"
const LMouseUp = "   왼쪽 마우스 떼기     "
const RMouseUp = "   오른쪽 마우스 떼기    "
const LMouseClick = "   왼쪽 마우스 클릭     "
const RMouseClick = "   오른쪽 마우스 클릭    "

var KeyboardMode = false
var MouseMode = false