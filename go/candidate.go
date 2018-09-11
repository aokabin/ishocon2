package main

import "fmt"

// Candidate Model
type Candidate struct {
	ID             int
	Name           string
	PoliticalParty string
	Votes int
	Sex            string
}

// CandidateElectionResult type
type CandidateElectionResult struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
	VoteCount      int
}

// PartyElectionResult type
type PartyElectionResult struct {
	PoliticalParty string
	VoteCount      int
}

func getAllCandidate() (candidates []Candidate) {
	rows, err := db.Query("SELECT * FROM candidates")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex, &c.Votes)
		if err != nil {
			panic(err.Error())
		}
		candidates = append(candidates, c)
	}
	return
}

func getCandidate(candidateID int) (c Candidate, err error) {
	row := db.QueryRow("SELECT * FROM candidates WHERE id = ?", candidateID)
	err = row.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex, &c.Votes)
	return
}

func getCandidateByName(name string) (c Candidate, err error) {
	c, ok := candidateByName[name]
	if !ok {
		err = fmt.Errorf("Error: %s", "Can not find candidate")
	}
	return
}

func getAllPartyName() (partyNames []string) {
	rows, err := db.Query("SELECT political_party FROM candidates GROUP BY political_party")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		partyNames = append(partyNames, name)
	}
	return
}

func getCandidatesByPoliticalParty(party string) (candidates []Candidate) {
	rows, err := db.Query("SELECT * FROM candidates WHERE political_party = ?", party)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex, &c.Votes)
		if err != nil {
			panic(err.Error())
		}
		candidates = append(candidates, c)
	}
	return
}

func getAllElectionResult() (result []CandidateElectionResult) {
	query := `SELECT c.id, c.name, c.political_party, c.sex, IFNULL(v.count, 0)
	FROM candidates AS c
	LEFT OUTER JOIN
	(SELECT candidate_id, SUM(count) AS count
	FROM votes
	GROUP BY candidate_id) AS v
	ON c.id = v.candidate_id
	ORDER BY v.count DESC`

	rows, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		r := CandidateElectionResult{}
		err = rows.Scan(&r.ID, &r.Name, &r.PoliticalParty, &r.Sex, &r.VoteCount)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, r)
	}
	return
}

func getVoteCountByPartyName(partyName string) int {
	query := `SELECT SUM(v.count)
	FROM candidates AS c
	LEFT OUTER JOIN
	(SELECT candidate_id, SUM(count) AS count
	FROM votes
	GROUP BY candidate_id) AS v
	ON c.id = v.candidate_id
	WHERE political_party = ?
	ORDER BY v.count DESC`

	rows, err := db.Query(query, partyName)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}
	return count
}
