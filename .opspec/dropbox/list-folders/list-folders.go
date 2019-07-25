package main

import (
	//"fmt"
	"github.com/zachpuck/dokuweb-api/pkg/dropbox"
)

func main() {

	f := dropbox.New()

	f.ListFolder()

	//f := dropbox.NewClient()
	//
	//items, err := f.GetAllInPath("")
	//if err != nil {
	//	fmt.Println("unable to get all items with error: ", err)
	//}
	//
	//for i := range items {
	//
	//}
	//fmt.Println("ITEMS: ", items)
}