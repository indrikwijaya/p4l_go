package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	fmt.Println("Let's simulate an election!")
	// first, let's seed the PRNG
	rand.Seed(time.Now().UnixNano())

	// next, read in the electoral votes
	electoralVotes := ReadElectoralVotes("electoralVotes.txt")

	filename := "debates.txt"
	// read in polls
	polls := ReadPollingData(filename)

	numTrials := 100000
	marginOfError := 0.1

	probability1, probability2, tieProbability := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("estimated probability of a candidate 1 win is", probability1)
	fmt.Println("estimated probability of a candidate 2 win is", probability2)
	fmt.Println("estimated probability of a tie is", tieProbability)
}

// SimulateMultipleElections takes polling data as a map of state names to floats (for candidate 1), along with a map of state names to electoral votes, a number fo trials, and a margin of error in the polls
//It returns 3 values: the estimated probabilities of candidate 1 winning, candidate 2 winning, and a tie.
func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint, numTrials int, marginOfError float64) (float64, float64, float64) {
		// keep track of no of simulated elections won by each candidate (and ties)
		winCount1 := 0
		winCount2 := 0
		tieCount := 0 //oh no!

		//simulate an election n times and update the counts each times
		for i := 0; i < numTrials; i++ {
				// call SimulateOneElection as a subroutine
				votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)
				// did candidate 1 or candidate 2 win?
				if votes1 > votes2 {
						winCount1++
				} else if votes1 < votes2 {
						winCount2++
				} else { //tie!
						tieCount++
				}
		}
		probability1 := float64(winCount1)/float64(numTrials)
		probability2 := float64(winCount2)/float64(numTrials)
		tieProbability := float64(tieCount)/float64(numTrials)
		return probability1, probability2, tieProbability
}

//SimulateOneElection takes a map of state names to polling % along w a map of electoral college votes and a margin of error.
//It returns the no of EC votes for each of the 2 candidates.
func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, marginOfError float64) (uint, uint){
		var collegeVotes1 uint = 0
		var collegeVotes2 uint = 0

		// range over all the states, and simulate the election in each state.
		for state := range polls {
				poll := polls[state] // current polling value in the state for candidate 1
				numVotes := electoralVotes[state]
				// let's adjust polling value w some randomized "noise"
				adjustedPoll := AddNoise(poll, marginOfError)
				// now we check who won state based on the adjusted poll ...
				if adjustedPoll >= 0.5 {
						//candidate 1 wins! give the EC votes for the state
						collegeVotes1 += numVotes
				} else {
						//candidate 2 wins!
						collegeVotes2 += numVotes
				}
		}

		return collegeVotes1, collegeVotes2
}

//AddNoise takes a polling for candidate 1 and a margin of error. It returns an adjusted polling value for candidate 1 aft adding a random noise.
func AddNoise(poll float64, marginOfError float64) float64 {
		x := rand.NormFloat64()// random number from standard normal distribution
		x = x/2 // x has ~95% chance of being btw -1 and 1
		x = x*marginOfError // x has 95% chance of being -marginOfError and marginOfError
		return x + poll
}
