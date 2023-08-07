package pad

type PadButton int

// memory: $4016(1P), $4017(2P)
// button: A, B, SELECT, START, UP, DOWN, LEFT, RIGHT
// bit   : 0, 1,      2,     3,  4,    5,    6,     7
// 1: pressed, 0: released
const (
	BUTTON_A PadButton = 1 << iota
	BUTTON_B
	BUTTON_SELECT
	BUTTON_START
	BUTTON_UP
	BUTTON_DOWN
	BUTTON_LEFT
	BUTTON_RIGHT
)

type Pad struct {
	ButtonStatus uint8
	ReadIdx      uint8
	Strobe       bool
}

func (p *Pad) PushButton(button PadButton) {
	switch button {
	case BUTTON_A:
		p.ButtonStatus |= uint8(BUTTON_A)
	case BUTTON_B:
		p.ButtonStatus |= uint8(BUTTON_B)
	case BUTTON_SELECT:
		p.ButtonStatus |= uint8(BUTTON_SELECT)
	case BUTTON_START:
		p.ButtonStatus |= uint8(BUTTON_START)
	case BUTTON_UP:
		p.ButtonStatus |= uint8(BUTTON_UP)
	case BUTTON_DOWN:
		p.ButtonStatus |= uint8(BUTTON_DOWN)
	case BUTTON_LEFT:
		p.ButtonStatus |= uint8(BUTTON_LEFT)
	case BUTTON_RIGHT:
		p.ButtonStatus |= uint8(BUTTON_RIGHT)
	}
}

func (p *Pad) ReleaseButton(button PadButton) {
	switch button {
	case BUTTON_A:
		p.ButtonStatus &= ^uint8(BUTTON_A)
	case BUTTON_B:
		p.ButtonStatus &= ^uint8(BUTTON_B)
	case BUTTON_SELECT:
		p.ButtonStatus &= ^uint8(BUTTON_SELECT)
	case BUTTON_START:
		p.ButtonStatus &= ^uint8(BUTTON_START)
	case BUTTON_UP:
		p.ButtonStatus &= ^uint8(BUTTON_UP)
	case BUTTON_DOWN:
		p.ButtonStatus &= ^uint8(BUTTON_DOWN)
	case BUTTON_LEFT:
		p.ButtonStatus &= ^uint8(BUTTON_LEFT)
	case BUTTON_RIGHT:
		p.ButtonStatus &= ^uint8(BUTTON_RIGHT)
	}
}

// ReadIdxで指定したボタンの状態を返す
// （ButtonsStatusをReadIdxだけ右にシフトして、最下位ビットを取得する）
// Storobeがfalseの場合はReadIdxをインクリメントし、8になった場合は0に戻す
// 8つのボタンの状態を順番に読み取る
func (p *Pad) Read() uint8 {
	ret := (p.ButtonStatus >> p.ReadIdx) & 1
	if !p.Strobe {
		p.ReadIdx++
		p.ReadIdx %= 8
	}

	return ret
}

func (p *Pad) SetStrobe(flag bool) {
	p.Strobe = flag
	if flag {
		p.ReadIdx = 0
	}
}
