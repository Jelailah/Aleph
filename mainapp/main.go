package main

import (
	"fmt"
	"strconv"

	w "fake.com/Buycoins/process"
)

var CreateBlock []byte

//this will create a new block if work has been done.
func ProofOfWork(WorkScore string, Approval bool, transaction string, Difficulty int) {

	if (Approval) && (len(WorkScore) >= Difficulty) {
		chain := w.InitBlockChain()
		chain.AddBlock(transaction)

		fmt.Println(chain)

	} else {
		fmt.Println("You need to do more work")
	}
}

func main() {
	workdone := w.FinalScore()
	workdoneLevel := workdone.Level
	workScore := strconv.Itoa(workdone.Total)

	Difficulty := 1 + workdoneLevel

	transaction := "Great job buddy!"
	Approval := workdone.ProjectApproval
	ProofOfWork(workScore, Approval, transaction, Difficulty)

	fmt.Println("New Block Created!")
}
