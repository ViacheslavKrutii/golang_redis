package main

import (
	"Proj/golang_pub-sub_observe/objects"
	"Proj/golang_pub-sub_observe/state"
)

func observerExample() {
	state := state.NewState()
	bob := objects.CreatePlayer("Bob", state)
	alice := objects.CreatePlayer("Alice", state)
	ignat := objects.CreatePlayer("Ignat", state)

	bob.CreateLobby()
	ignat.CreateLobby()
	bob.InvitePlayer(alice)
	ignat.InvitePlayer(alice)
	alice.CheckInvites()
	bob.Move()
	bob.Move()
	bob.Move()
	bob.Move()
	bob.Move()

}

func main() {
	observerExample()

}
