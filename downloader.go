/**

NOT COMPLETED YET!

*/
package main

import "fmt"

// Interface for files
type TGFile interface {
	Download() bool
	GetID() string
	GetType() string
}

// if the file photo?
type TGPhoto struct {
	ID     string
	Width  int
	Height int
	Size   int
}

func (f TGPhoto) Download() bool {
	fmt.Println("TO IMPLEMENT")
	return true
}

func (f TGPhoto) GetID() string {
	return f.ID
}

func (f TGPhoto) GetType() string {
	return "Photo"
}

type TGVideo struct {
	Duration int
	Width    int
	Height   int
	MimeType string
	Thumb    TGPhoto
	ID       string
	Size     int
}

func (f TGVideo) Download() bool {
	fmt.Println("VIDEO: TO BE IMPLEMENTED")
	return true
}

func (f TGVideo) GetID() string {
	return f.ID
}

func (f TGVideo) GetType() string {
	return "Video"
}
