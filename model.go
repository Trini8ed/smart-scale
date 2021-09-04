package main

type Order struct {
	Weight int    `json:"weight" firestore:"weight"`
	Units  string `json:"units" firestore:"units"`
}
