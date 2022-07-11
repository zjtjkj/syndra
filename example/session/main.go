package main

import (
	"fmt"
	"github.com/zjtjkj/syndra/utils/session"
)

func main() {
	// init manager
	manager, err := session.NewManager(session.Memory, 172800)
	if err != nil {
		panic(err)
	}
	// create new session
	sess1, err := manager.SessionStart("")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sess1 id: %s, name: %s, age: %d\n", sess1.SessionID(), sess1.Get("name"), sess1.Get("age"))
	_ = sess1.Set("name", "Tom")
	_ = sess1.Set("age", 18)
	// get session
	sess2, err := manager.SessionStart(sess1.SessionID())
	if err != nil {
		panic(err)
	}
	fmt.Printf("sess2 id: %s, name: %s, age: %d\n", sess2.SessionID(), sess2.Get("name"), sess2.Get("age"))
	// destroy session
	err = manager.SessionDestroy(sess2.SessionID())
	if err != nil {
		panic(err)
	}
	// get session again
	sess3, err := manager.SessionStart(sess2.SessionID())
	if err != nil {
		panic(err)
	}
	fmt.Printf("sess3 id: %s, name: %s, age: %d\n", sess3.SessionID(), sess3.Get("name"), sess3.Get("age"))
}
