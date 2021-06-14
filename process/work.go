package process

import (
	"fmt"
	"log"
	"strconv"

	bolt "github.com/boltdb/bolt"
)

type Validation struct {
	ProjectApproval bool
	Level           int
	Total           int
}

func scoreComputation(prevScore, hours, level, projectScore int) int {
	Score := hours * level * projectScore

	return Score
}

func FinalScore() *Validation {
	fmt.Println("---This is a work computation module---")

	var totalScore int
	var prevScore int
	var hours int        //approximation of hours spent learning
	var level int        ////the current level in the learning track
	var projectScore int //the score assigned by facilitator on independent project /PS: maximum value for projectScore is 10
	var Approval bool    //whether the project was approved or not

	//We'll be taking sample inputs that'll help us calculate the work done so far

	fmt.Println("Hours worked: ")
	fmt.Scanln(&hours)

	fmt.Println("Level: ")
	fmt.Scanln(&level)

	fmt.Println("Project Score: ")
	fmt.Scanln(&projectScore)

	fmt.Println("Has this project been approved? (true/false): ")
	fmt.Scanln(&Approval)

	db, err := bolt.Open("levels.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.View(func(tx *bolt.Tx) error {
		m := make(map[string]string)
		b := tx.Bucket([]byte("Bucket"))
		c := b.Cursor()
		k, v := c.Last()
		z := string(k)
		m[z] = string(v)
		a, e := strconv.Atoi(m[z])
		prevScore = a
		return e
	})

	db.Update(func(tx *bolt.Tx) error {
		//tx.CreateBucket([]byte("Bucket"))
		b := tx.Bucket([]byte("Bucket"))
		err := b.Put([]byte(strconv.Itoa(level)), []byte(strconv.Itoa(totalScore)))
		return err
	})

	defer db.Close()

	score := scoreComputation(prevScore, hours, level, projectScore)

	totalScore = score + totalScore
	prevScore = totalScore

	validate := &Validation{Approval, level, totalScore}
	return validate

}
