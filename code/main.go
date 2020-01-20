package main

import (
	"container/list"
	"fmt"
)

const (
	ANDROID_PLATFORM = 1
	IOS_PLATFORM     = 2
)

type (
	Button interface {
		click()
	}
	Scroll interface {
		scroll()
	}
	Menu interface {
		selectItem()
	}
	GdiFactory interface {
		createButton() *Button
		createMenu() *Menu
		createScroll() *Scroll
	}
)

/*
---------------------------------------------------------------
IOS IMP
*/
type IOSFactory struct {
}

func NewIOSFactory() GdiFactory {
	var g GdiFactory
	g = IOSFactory{}
	return g
}

func (I IOSFactory) createButton() *Button {
	var b Button = IOSButton{}
	return &b
}

func (I IOSFactory) createMenu() *Menu {
	var m Menu = IOSMenu{}
	return &m
}

func (I IOSFactory) createScroll() *Scroll {
	var s Scroll = IOSScroll{}
	return &s
}

type (
	IOSButton struct {
	}
	IOSMenu struct {
	}
	IOSScroll struct {
	}
)

func (I IOSButton) click() {
	fmt.Println("click from ios button")
}
func (I IOSScroll) scroll() {
	fmt.Println("scroll from ios scroll")
}
func (I IOSMenu) selectItem() {
	fmt.Println("selectItem from ios menu")
}

/*
---------------------------------------------------------------
Android IMP
*/
type AndroidFactory struct {
}

func (I AndroidFactory) CreateButton() *Button {
	panic("implement me")
}

func NewAndroidFactory() GdiFactory {
	var g GdiFactory
	g = AndroidFactory{}
	return g
}
func (I AndroidFactory) createButton() *Button {
	var b Button = AndroidButton{}
	return &b
}

func (I AndroidFactory) createMenu() *Menu {
	var m Menu = AndroidMenu{}
	return &m
}

func (I AndroidFactory) createScroll() *Scroll {
	var s Scroll = AndroidScroll{}
	return &s
}

type (
	AndroidButton struct {
	}
	AndroidMenu struct {
	}
	AndroidScroll struct {
	}
)

func (I AndroidButton) click() {
	fmt.Println("click from android button")
}
func (I AndroidScroll) scroll() {
	fmt.Println("scroll from android scroll")
}
func (I AndroidMenu) selectItem() {
	fmt.Println("selectItem from android menu")
}

/*
------------------------------------------------------
GDI Configurator
*/
type GdiConfigurator struct {
	Button Button
	Menu   Menu
	Scroll Scroll
}

func NewGdiConfigurator(platform int) *GdiConfigurator {
	var f GdiFactory
	switch platform {
	case ANDROID_PLATFORM:
		f = NewAndroidFactory()
	case IOS_PLATFORM:
		f = NewIOSFactory()
	}
	return &GdiConfigurator{
		Button: *f.createButton(),
		Menu:   *f.createMenu(),
		Scroll: *f.createScroll(),
	}

}
func Load(conf *GdiConfigurator) {
	conf.Scroll.scroll()
	conf.Button.click()
	conf.Menu.selectItem()
}

/*
Signleton
*/
type singleton struct {
	platformItems list.List
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{} // Это НЕ потоко-безопасно
		instance.platformItems.Init()
	}
	return instance
}
func printList(items list.List) {
	for e := items.Front(); e != nil; e = e.Next() {
		fmt.Println("%#v", e.Value)
	}
}
func main() {
	c := NewGdiConfigurator(IOS_PLATFORM)
	Load(c)
	GetInstance().platformItems.PushBack(c.Menu)
	printList(GetInstance().platformItems)
}
