package main

import (
	"strings"
)

// Vote Model
type Vote struct {
	ID          int
	UserID      int
	CandidateID int
	Count int
	Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT votes FROM candidates WHERE id = ?", candidateID)
	row.Scan(&count)

	//count = len(votes[candidateID])
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM votes WHERE user_id =  ?", userID)
	row.Scan(&count)
	return
}

func createVote(userID int, candidateID int, keyword string, count int) {
	insert := `INSERT INTO votes (user_id, candidate_id, keyword, count) VALUES (?, ?, ?, ?)`
	update := `UPDATE candidates SET votes = votes + ? WHERE id = ?`

	db.Exec(insert, userID, candidateID, keyword, count)
	db.Exec(update, count, candidateID)

	//id, _ := result.LastInsertId()
	//
	//vote := Vote{
	//	ID: int(id),
	//	UserID: userID,
	//	CandidateID: candidateID,
	//	Keyword: keyword,
	//}

	//votes[candidateID] = append(votes[candidateID], vote)
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	args := []interface{}{}
	for _, candidateID := range candidateIDs {
		args = append(args, candidateID)
	}

	rows, err := db.Query(`
    SELECT keyword
    FROM votes
		WHERE candidate_id IN (`+strings.Join(strings.Split(strings.Repeat("?", len(candidateIDs)), ""), ",")+`)
    GROUP BY keyword
    ORDER BY SUM(count) DESC
    LIMIT 10`, args...)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			panic(err.Error())
		}
		voices = append(voices, keyword)
	}
	return
}
