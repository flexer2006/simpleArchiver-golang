package vlcPack

type MockEncoder struct {
	encodeFunc func(string) (string, error)
}

func (m *MockEncoder) Encode(str string) (string, error) {
	return m.encodeFunc(str)
}
