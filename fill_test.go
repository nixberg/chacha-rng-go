package chacharand

import "testing"

func TestFillUint8(t *testing.T) {
	rng := Zero20(0)

	expectedA := [65]uint8{
		0x76, 0xb8, 0xe0, 0xad, 0xa0, 0xf1, 0x3d, 0x90, 0x40, 0x5d, 0x6a, 0xe5, 0x53, 0x86,
		0xbd, 0x28, 0xbd, 0xd2, 0x19, 0xb8, 0xa0, 0x8d, 0xed, 0x1a, 0xa8, 0x36, 0xef, 0xcc,
		0x8b, 0x77, 0x0d, 0xc7, 0xda, 0x41, 0x59, 0x7c, 0x51, 0x57, 0x48, 0x8d, 0x77, 0x24,
		0xe0, 0x3f, 0xb8, 0xd8, 0x4a, 0x37, 0x6a, 0x43, 0xb8, 0xf4, 0x15, 0x18, 0xa1, 0x1c,
		0xc3, 0x87, 0xb6, 0x69, 0xb2, 0xee, 0x65, 0x86, 0x9f,
	}
	expectedB := [9]uint8{0x55, 0x51, 0x38, 0x7a, 0x98, 0xba, 0x97, 0x7c, 0x73}

	a := make([]uint8, 65)
	b := make([]uint8, 9)

	rng.FillUint8(a[:])
	rng.FillUint8(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}

func TestFillUint8Logic(t *testing.T) {
	expected := [65]uint8{
		0x76, 0xb8, 0xe0, 0xad, 0xa0, 0xf1, 0x3d, 0x90, 0x40, 0x5d, 0x6a, 0xe5, 0x53, 0x86,
		0xbd, 0x28, 0xbd, 0xd2, 0x19, 0xb8, 0xa0, 0x8d, 0xed, 0x1a, 0xa8, 0x36, 0xef, 0xcc,
		0x8b, 0x77, 0x0d, 0xc7, 0xda, 0x41, 0x59, 0x7c, 0x51, 0x57, 0x48, 0x8d, 0x77, 0x24,
		0xe0, 0x3f, 0xb8, 0xd8, 0x4a, 0x37, 0x6a, 0x43, 0xb8, 0xf4, 0x15, 0x18, 0xa1, 0x1c,
		0xc3, 0x87, 0xb6, 0x69, 0xb2, 0xee, 0x65, 0x86, 0x9f,
	}

	for count := range expected {
		rng := Zero20(0)

		slice := make([]uint8, count)
		rng.FillUint8(slice)

		for i, expected := range expected[:count] {
			if slice[i] != expected {
				t.Fail()
			}
		}
	}
}

func TestFillUint16(t *testing.T) {
	rng := Zero20(0)

	expectedA := [33]uint16{
		0xb876, 0xade0, 0xf1a0, 0x903d, 0x5d40, 0xe56a, 0x8653, 0x28bd, 0xd2bd, 0xb819, 0x8da0,
		0x1aed, 0x36a8, 0xccef, 0x778b, 0xc70d, 0x41da, 0x7c59, 0x5751, 0x8d48, 0x2477, 0x3fe0,
		0xd8b8, 0x374a, 0x436a, 0xf4b8, 0x1815, 0x1ca1, 0x87c3, 0x69b6, 0xeeb2, 0x8665, 0x079f,
	}
	expectedB := [5]uint16{0x5155, 0x7a38, 0xba98, 0x7c97, 0x2d73}

	a := make([]uint16, 33)
	b := make([]uint16, 5)

	rng.FillUint16(a[:])
	rng.FillUint16(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}

func TestFillUint16Logic(t *testing.T) {
	expected := [33]uint16{
		0xb876, 0xade0, 0xf1a0, 0x903d, 0x5d40, 0xe56a, 0x8653, 0x28bd, 0xd2bd, 0xb819, 0x8da0,
		0x1aed, 0x36a8, 0xccef, 0x778b, 0xc70d, 0x41da, 0x7c59, 0x5751, 0x8d48, 0x2477, 0x3fe0,
		0xd8b8, 0x374a, 0x436a, 0xf4b8, 0x1815, 0x1ca1, 0x87c3, 0x69b6, 0xeeb2, 0x8665, 0x079f,
	}

	for count := range expected {
		rng := Zero20(0)

		slice := make([]uint16, count)
		rng.FillUint16(slice)

		for i, expected := range expected[:count] {
			if slice[i] != expected {
				t.Fail()
			}
		}
	}
}

func TestFillUint32(t *testing.T) {
	rng := Zero20(0)

	expectedA := [17]uint32{
		0xade0b876, 0x903df1a0, 0xe56a5d40, 0x28bd8653, 0xb819d2bd, 0x1aed8da0, 0xccef36a8,
		0xc70d778b, 0x7c5941da, 0x8d485751, 0x3fe02477, 0x374ad8b8, 0xf4b8436a, 0x1ca11815,
		0x69b687c3, 0x8665eeb2, 0xbee7079f,
	}
	expectedB := [5]uint32{0x7a385155, 0x7c97ba98, 0x0d082d73, 0xa0290fcb, 0x6965e348}

	a := make([]uint32, 17)
	b := make([]uint32, 5)

	rng.FillUint32(a[:])
	rng.FillUint32(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}

func TestFillUint64(t *testing.T) {
	rng := Zero20(0)

	expectedA := [9]uint64{
		0x903df1a0ade0b876, 0x28bd8653e56a5d40, 0x1aed8da0b819d2bd, 0xc70d778bccef36a8,
		0x8d4857517c5941da, 0x374ad8b83fe02477, 0x1ca11815f4b8436a, 0x8665eeb269b687c3,
		0x7a385155bee7079f,
	}
	expectedB := [2]uint64{0x0d082d737c97ba98, 0x6965e348a0290fcb}

	a := make([]uint64, 9)
	b := make([]uint64, 2)

	rng.FillUint64(a[:])
	rng.FillUint64(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}

func TestFillFloat32(t *testing.T) {
	rng := Zero20(0)

	expectedA := [17]float32{
		0.679210186, 0.563445151, 0.896154225, 0.159141898, 0.719143987, 0.105187237,
		0.800525069, 0.777549207, 0.485736907, 0.551885068, 0.249513865, 0.215985775,
		0.955936611, 0.111833096, 0.412941396, 0.524992824, 0.745712698,
	}
	expectedB := [3]float32{0.477421820, 0.486690164, 0.050906003}

	a := make([]float32, 17)
	b := make([]float32, 3)

	rng.FillFloat32(a[:])
	rng.FillFloat32(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}

func TestFillFloat64(t *testing.T) {
	rng := Zero20(0)

	expectedA := [9]float64{
		0.56344518826324730, 0.15914191768880792, 0.10518727468306754, 0.77754923976038692,
		0.55188508738897202, 0.21598581789282933, 0.11183310066255192, 0.52499286514076893,
		0.47742184012795619,
	}
	expectedB := [2]float64{0.05090602941035527, 0.41171093485914778}

	a := make([]float64, 9)
	b := make([]float64, 2)

	rng.FillFloat64(a[:])
	rng.FillFloat64(b[:])

	for i, expected := range expectedA {
		if a[i] != expected {
			t.Fail()
		}
	}
	for i, expected := range expectedB {
		if b[i] != expected {
			t.Fail()
		}
	}
}
