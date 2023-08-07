package pad

import "testing"

func TestPushButton(t *testing.T) {
	testCases := []struct {
		name    string
		buttons []PadButton
		want    uint8
	}{
		{
			name:    "single",
			buttons: []PadButton{BUTTON_A},
			want:    uint8(BUTTON_A),
		},
		{
			name:    "all",
			buttons: []PadButton{BUTTON_A, BUTTON_B, BUTTON_SELECT, BUTTON_START, BUTTON_UP, BUTTON_DOWN, BUTTON_LEFT, BUTTON_RIGHT},
			want:    uint8(BUTTON_A) | uint8(BUTTON_B) | uint8(BUTTON_SELECT) | uint8(BUTTON_START) | uint8(BUTTON_UP) | uint8(BUTTON_DOWN) | uint8(BUTTON_LEFT) | uint8(BUTTON_RIGHT),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pad := Pad{}
			for _, button := range tc.buttons {
				pad.PushButton(button)
			}
			if pad.ButtonStatus != tc.want {
				t.Errorf("pad.ButtonStatus = %v, want %v", pad.ButtonStatus, tc.want)
			}
		})
	}
}

func TestReleaseButton(t *testing.T) {
	testCases := []struct {
		name    string
		push    []PadButton
		release []PadButton
		want    uint8
	}{
		{
			name:    "single",
			push:    []PadButton{BUTTON_A, BUTTON_B},
			release: []PadButton{BUTTON_A},
			want:    uint8(BUTTON_B),
		},
		{
			name:    "all",
			push:    []PadButton{BUTTON_A, BUTTON_B, BUTTON_SELECT, BUTTON_START, BUTTON_UP, BUTTON_DOWN, BUTTON_LEFT, BUTTON_RIGHT},
			release: []PadButton{BUTTON_A, BUTTON_B, BUTTON_SELECT, BUTTON_START, BUTTON_UP, BUTTON_DOWN, BUTTON_LEFT, BUTTON_RIGHT},
			want:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pad := Pad{}
			for _, button := range tc.push {
				pad.PushButton(button)
			}
			for _, button := range tc.release {
				pad.ReleaseButton(button)
			}
			if pad.ButtonStatus != tc.want {
				t.Errorf("pad.ButtonStatus = %v, want %v", pad.ButtonStatus, tc.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
	testCases := []struct {
		name   string
		push   []PadButton
		strobe bool
		wants  []uint8
	}{
		{
			name:   "strobe disabled (all)",
			push:   []PadButton{BUTTON_A, BUTTON_B, BUTTON_SELECT, BUTTON_START, BUTTON_UP, BUTTON_DOWN, BUTTON_LEFT, BUTTON_RIGHT},
			strobe: false,
			wants:  []uint8{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:   "strobe disabled (button B)",
			push:   []PadButton{BUTTON_B},
			strobe: false,
			wants:  []uint8{0, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:   "strobe enabled (button A)",
			push:   []PadButton{BUTTON_A},
			strobe: true,
			wants:  []uint8{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:   "strobe enabled (button B)",
			push:   []PadButton{BUTTON_B},
			strobe: true,
			wants:  []uint8{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pad := Pad{
				Strobe:  tc.strobe,
				ReadIdx: 0,
			}
			for _, button := range tc.push {
				pad.PushButton(button)
			}

			for _, want := range tc.wants {
				read := pad.Read()
				if read != want {
					t.Errorf("pad.Read() = %v, want %v", read, want)
				}
			}
		})
	}
}

func TestSetStrobe(t *testing.T) {
	testCases := []struct {
		name        string
		strobe      bool
		idx         uint8
		wantStrobe  bool
		wantReadIdx uint8
	}{
		{
			name:        "strobe enabled",
			idx:         7,
			strobe:      true,
			wantStrobe:  true,
			wantReadIdx: 0,
		},
		{
			name:        "strobe disabled",
			idx:         7,
			strobe:      false,
			wantStrobe:  false,
			wantReadIdx: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pad := Pad{
				ReadIdx: tc.idx,
			}
			pad.SetStrobe(tc.strobe)
			if pad.Strobe != tc.wantStrobe {
				t.Errorf("pad.Strobe = %v, want %v", pad.Strobe, tc.wantStrobe)
			}
			if pad.ReadIdx != tc.wantReadIdx {
				t.Errorf("pad.ReadIdx = %v, want %v", pad.ReadIdx, tc.wantReadIdx)
			}
		})
	}
}
