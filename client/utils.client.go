package client

type windowSize struct {
	width  float32
	height float32
}

var WindowSize = windowSize{450, 450}
var ButtonSize = windowSize{20, 100}

var DefaultModeString = "아래 버튼 중 수행할 매크로를 선택하세요."
var KeyboardModeString = "키보드 매크로를 입력하세요."
var MouseModeString = "(F7)을 눌러 현재 마우스 위치를 복사합니다."

var LMouseDown = "   왼쪽 마우스 누르기  "
var RMouseDown = "   오른쪽 마우스 누르기"
var LMouseUp = "   왼쪽 마우스 떼기     "
var RMouseUp = "   오른쪽 마우스 떼기    "
var LMouseClick = "   왼쪽 마우스 클릭     "
var RMouseClick = "   오른쪽 마우스 클릭    "

var KeyboardMode = false
var MouseMode = false