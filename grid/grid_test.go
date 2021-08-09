package grid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, abs(-5))
	assert.Equal(t, 5, abs(5))
	assert.Equal(t, 0, abs(-0))
	assert.Equal(t, 69, abs('E'))
}

func TestDistanceTo(t *testing.T) {
	a := &Point{0, 0}
	b := &Point{1, 1}

	assert.Equal(t, 2, a.DistanceTo(b))

	a = &Point{0, 0}
	b = &Point{4, 4}

	assert.Equal(t, 8, a.DistanceTo(b))

	a = &Point{4, 1}
	b = &Point{1, 5}

	assert.Equal(t, 7, a.DistanceTo(b))
}

func TestRemovePath(t *testing.T) {
	path := NewPath(&Point{1, 1}, &Point{2, 1}, &Point{1, 2}, &Point{4, 3})
	fmt.Println(path)

	fmt.Println(path.Remove(&Point{2, 1}))

	t.Fail() //todo add more cases
}

func TestStringifyPathTo(t *testing.T) {
	t.Fail() //todo add more cases
	a := &Point{0,0}
	b := &Point{1,3}
	output := a.StringifyPathTo(b)

	assert.Equal(t, "ENNN", output)

	a = &Point{2,3}
	b =	&Point{4,1}
	output = a.StringifyPathTo(b)

	assert.Equal(t, "EESS", output)
}

func TestStringifyPath(t *testing.T) {
	t.Fail() //todo add more cases
	path := NewPath(&Point{0,0},
		&Point{1,3},
		&Point{4,4},
		&Point{4,2},
		&Point{4,2},
		&Point{0,1},
		&Point{3,2},
		&Point{2,3},
		&Point{4,1})

	assert.Equal(t, "ENNNDEEENDSSDDWWWWSDEEENDWNDEESSD", path.StringifyPath())
}
