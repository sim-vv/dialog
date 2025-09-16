//go:build darwin

package dialog

func (b *MsgBuilder) yesNo() bool {
	return false
}

func (b *MsgBuilder) info() {
	return
}

func (b *MsgBuilder) error() {
	return
}

func (b *FileBuilder) load() (string, error) {
	return "", nil
}

func (b *FileBuilder) loads() ([]string, error) {
	return []string{}, nil
}

func (b *FileBuilder) save() (string, error) {
	return "", nil
}

func (b *DirectoryBuilder) browse() (string, error) {
	return "", nil
}
