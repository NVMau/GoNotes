package main

import (
	"github.com/NVMau/GoNotes/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8002")
}
