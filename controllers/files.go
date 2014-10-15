package controllers

import (
	"github.com/Marktown/frontend/backends/file_system"
	"text/scanner"
)

type PathItem struct {
	Path  string     `json:"name"`
	Items []PathItem `json:"items"`
}

type Textfile struct {
	Data string `json:"data"`
}

type FilesController struct {
	BaseController
}

func (this *FilesController) Prepare() {
	this.SetupEnv()
	this.Layout = "layouts/default.html.tpl"
}

func (this *FilesController) Index() {
	this.TplNames = "files/index.html.tpl"
	pathItems := PathItem{"/", []PathItem{
		PathItem{"/foo", []PathItem{}},
		PathItem{"/bar", []PathItem{
			PathItem{"/bar/bax", []PathItem{}},
			PathItem{"/bar/baz", []PathItem{}},
		}},
		PathItem{"/fax", []PathItem{
			PathItem{"/fax/ban", []PathItem{}},
			PathItem{"/fax/bam", []PathItem{}},
		}},
		PathItem{"/faz", []PathItem{}},
	}}
	this.Data["json"] = &pathItems
	//this.ServeJson()

	fs := file_system.NewFileStore()
	list, err := fs.ReadDir("")
	//TODO error handling
	if err != nil {
		return
	}
	//TODO use ReadDirTree for recursive reading, must be implemented
	testPathItems := PathItem{"/", []PathItem{}}
	items := make([]PathItem, len(list))
	for index, file := range list {
		items[index] = PathItem{file.Name(), []PathItem{}}
	}
	testPathItems.Items = items
	this.Data["json"] = &testPathItems
	this.ServeJson()
}

func (this *FilesController) New() {
	this.TplNames = "files/new.html.tpl"
}

func (this *FilesController) Read() {
	textfile := Textfile{}
	//dataAsBytes, err := ioutil.ReadFile("tests/assets/testfile.md")
	fs := file_system.NewFileStore()
	fs.RootPath = "tests/assets/"
	reader, err := fs.ReadFile("testfile.md")
	if err != nil {
		return
	}

	textfile.Data = ""
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace = 1
	tok := s.Scan()
	for tok != scanner.EOF {
		textfile.Data += s.TokenText()
		tok = s.Scan()
	}
	//textfile.Data = string(dataAsBytes)
	this.Data["json"] = &textfile
	this.ServeJson()
}

func (this *FilesController) Update() {
	this.TplNames = "files/update.html.tpl"
}

func (this *FilesController) Delete() {
	this.TplNames = "files/delete.html.tpl"
}
