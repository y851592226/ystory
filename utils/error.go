package utils

import (
	"errors"
)

var NFD = errors.New("Not Found")

var SPE = errors.New("Spider Error")

type controller interface {
	Abort(code string)
}

func CkeckError(this controller, err error) {
	if err == NFD {
		this.Abort("404")
	} else if err != nil {
		panic(err)
	}
}
