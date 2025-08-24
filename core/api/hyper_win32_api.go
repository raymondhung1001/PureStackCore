package api

import (
	"fmt"
	"sync"

	"github.com/go-ole/go-ole"
)

type HyperWin32Api struct {
	initialized bool
	locator *ole.IDispatch
	service *ole.IDispatch
	mutex sync.Mutex
	logger *
}