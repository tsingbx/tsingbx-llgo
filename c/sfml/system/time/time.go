package time

import (
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs sfml-system); -lsfml-system"
)

type Time struct {
	m_microseconds int64
}

//go:linkname Init C._ZN2sf4TimeC1Ev
func Init(time *Time) {
	time.m_microseconds = 0
}

//go:linkname AsSeconds C._ZNK2sf4Time9asSecondsEv
func AsSeconds(time *Time) float32

//go:linkname AsMilliseconds C._ZNK2sf4Time14asMillisecondsEv
func AsMilliseconds(time *Time) int32

//go:linkname AsMicroseconds C._ZNK2sf4Time14asMicrosecondsEv
func AsMicroseconds(time *Time) int64
