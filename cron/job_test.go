package cron

import (
	"testing"
	"time"
)

func TestHappyJobOrder(t *testing.T) {
	results := make(map[string]int)
	f := func(j *Job) error {
		if _, ok := results[j.Name]; ok {
			results[j.Name] = results[j.Name] + 1
		} else {
			results[j.Name] = 1
		}
		return nil
	}

	// timeSleep = timeWinner + 2
	var timeSleep, timeWinner, timeLoser time.Duration = time.Second * 7, time.Second * 1, time.Second * 5
	var winnerScore, loserScore int

	jobWinner := NewJob("jobWinner", timeWinner, f)
	jobLoser := NewJob("jobLoser", timeLoser, f)

	testService.Add(jobWinner)
	testService.Add(jobLoser)
	errCh := testService.Start()
	go func() {
		for {
			jE := <-errCh
			if jE.E != nil {
				t.Errorf("found error running job(%s): %v", jE.Job.Name, jE.E)
			} else {
				//t.Logf("Succes: %s", jE.Job.Name)
			}
		}
	}()
	time.Sleep(timeSleep)

	// Check that results[jobWinner] > 5
	if w, ok := results["jobWinner"]; ok {
		winnerScore = w
		if winnerScore >= 4 {
			t.Logf("Winner score: %d", winnerScore)
		} else {
			t.Errorf("failed %v second (%v second sleep) score: %d", timeWinner, timeSleep, winnerScore)
		}
	} else {
		t.Errorf("unable to find scores for 'jobWinner'")
	}

	// Check that results[jobLoser] == 1
	if l, ok := results["jobLoser"]; ok {
		loserScore = l
		if loserScore == 1 {
			t.Logf("Loser score: %d", loserScore)
		} else {
			t.Errorf("failed %v second (%v second sleep) score: %d", timeLoser, timeSleep, loserScore)
		}
	} else {
		t.Errorf("unable to find scores for 'jobLoser'")
	}

	// Check that winnerScore < loserScore
	if winnerScore <= loserScore {
		t.Errorf("winnerScore(%v) <= loserScore(%v)", winnerScore, loserScore)
	}

	// timeWinner = 1 sec
	// timeLoser  = 5 sec
	// timeSleep  = 7 sec
	if timeWinner > timeLoser {
		t.Errorf("invalid test parameters timeLoser(%v) < timeWinner(%v)", timeLoser, timeWinner)
	}
	if timeLoser > timeSleep {
		t.Errorf("invalid test parameters timeWinner(%v) < timeSleep(%v)", timeWinner, timeSleep)
	}
}
