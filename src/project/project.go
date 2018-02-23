package project

import (
	"../config"
)

type Project struct{}

func Discover(path string, settings config.Settings) Project {
}

func (p Project) Launch() error {

}
