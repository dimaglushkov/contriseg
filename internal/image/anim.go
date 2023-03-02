package image

import (
	"fmt"
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/image/animation"
	"strings"
)

type AnimationFunc func(cal internal.Calendar) []internal.Calendar

var animationsMap = map[string]AnimationFunc{
	"bfs":  animation.DrawBFS,
	"move": animation.MoveColLeft,
	"cbc":  animation.DrawColByColLeft,
}

func GetAvailableAnimations() []string {
	anims := make([]string, 0, len(animationsMap))
	for k := range animationsMap {
		anims = append(anims, k)
	}
	return anims
}

func GetAnimationIterator(animAlias string) (AnimationFunc, error) {
	var anim AnimationFunc
	var ok bool

	if anim, ok = animationsMap[strings.ToLower(animAlias)]; !ok {
		return nil, fmt.Errorf(
			"unknown animation: %s, available animations are: %s (case insesnsetive)",
			animAlias,
			strings.Join(GetAvailableAnimations(), ", "),
		)
	}

	return anim, nil
}
