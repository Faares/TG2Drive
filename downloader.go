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

/*
"video":{"duration":10,"width":1920,"height":1080,"mime_type":"video/mp4","thumb":{"file_id":"AAQEABMgXx8bAAQ-1LxCgAtwK-sgAAIC","file_size":18290,"width":320,"height":180},"file_id":"BAADBAADgQcAAm_0wFMJa8z-C8n7jQI","file_size":11256084}}}]
*/
type TGVideo struct {
	Duration int
	Width    int
	height   int
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
