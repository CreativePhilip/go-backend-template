package main

import (
	"fmt"
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/internal/auth/repositories"
	"log"
	"slices"
)

var chunkSize = 100

func main() {
	tx := db.Client()
	sessions := repositories.DbUserSessionRepository{Db: tx}

	fmt.Println("Fetching expired sessions ...")
	expiredSessions, err := sessions.GetAllExpired()

	if err != nil {
		log.Fatalf("Error fetching expired sessions: %v\n", err)
	}
	if len(expiredSessions) == 0 {
		fmt.Println("No expired sessions found")
		return
	}

	fmt.Printf("Found %d expired sessions\n", len(expiredSessions))

	currentChunk := 1
	chunkCount := (len(expiredSessions) / chunkSize) + 1

	for chunk := range slices.Chunk(expiredSessions, chunkSize) {
		fmt.Printf("Processing batch %d out of %d\n", currentChunk, chunkCount)

		err := sessions.DeleteByIds(sessionSliceToIds(chunk))

		if err != nil {
			log.Fatalf("Error deleting expired sessions: %v\n", err)
		}

		currentChunk++
	}

	tx.Close()

	fmt.Print("\n\n")
	fmt.Println("=========================")
	fmt.Println("Finished vacuuming expired sessions")
	fmt.Println("=========================")
}

func sessionSliceToIds(chunk []*repositories.UserSession) []uint {
	ids := make([]uint, 0)

	for _, session := range chunk {
		ids = append(ids, session.Id)
	}

	return ids
}
