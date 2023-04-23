package sample

import "gitlab.com/GeorgePrathap/protobuf/pb/pb"

// NewKeyboard returns a new sample keyboard
func NewKeyboard() (keyboard *pb.Keyboard) {
	keyboard = &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return
}
