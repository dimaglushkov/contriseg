package image

import (
	"fmt"
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/image/animations"
	"strings"
)

type AnimationIterator func(cal internal.Calendar) []internal.Calendar

var animationsMap = map[string]AnimationIterator{
	"bfs":  animations.CalendarBFSIterations,
	"move": animations.CalendarMoveColLeftIterations,
	"cbc":  animations.CalendarColByColIterations,
}

func GetAvailableAnimations() []string {
	availableAnimations := make([]string, 0, len(animationsMap))
	for k := range animationsMap {
		availableAnimations = append(availableAnimations, k)
	}
	return availableAnimations
}

func GetAnimationIterator(animation string) (AnimationIterator, error) {
	var iter AnimationIterator
	var ok bool

	if iter, ok = animationsMap[strings.ToLower(animation)]; !ok {
		return nil, fmt.Errorf("unknown framing animation: %s, available animations are: %s (case insesnsetive)", animation, strings.Join(GetAvailableAnimations(), ", "))
	}

	return iter, nil
}
