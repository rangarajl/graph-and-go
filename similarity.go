package main

import (
    "context"
    "fmt"
    "log"

    "github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
    // Neo4j connection details
    uri := "neo4j://localhost:7687"
    username := "neo4j"
    password := "password"

    // Connect to Neo4j
    driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
    if err != nil {
        log.Fatal("Error connecting to Neo4j:", err)
    }
    defer driver.Close()
    
    session := driver.NewSession(neo4j.SessionConfig{})
    defer session.Close()

    // Example: Calculate similarity between two players
    result, err := session.Run("MATCH (p1:Player)-[:PLAYED_AGAINST]-(p2:Player) " +
        "RETURN p1.name AS player1, p2.name AS player2, " +
        "gds.alpha.similarity.pearson([p1.homeRuns, p1.battingAverage], [p2.homeRuns, p2.battingAverage]) AS similarity", nil)
    if err != nil {
        log.Fatal("Error querying similarity:", err)
    }

    for result.Next() {
        fmt.Printf("Player 1: %s, Player 2: %s, Similarity: %f\n", result.Record().GetByIndex(0), result.Record().GetByIndex(1), result.Record().GetByIndex(2))
    }

    if err := result.Err(); err != nil {
        log.Fatal(err)
    }

    // Example: Aggregate team success (simplified example)
    // Assuming success is measured by the sum of home runs
    aggResult, err := session.Run("MATCH (p:Player)-[:PLAYS_FOR]->(t:Team) " +
        "RETURN t.name AS team, sum(p.homeRuns) AS totalHomeRuns", nil)
    if err != nil {
        log.Fatal("Error querying team success:", err)
    }

    for aggResult.Next() {
        fmt.Printf("Team: %s, Total Home Runs: %d\n", aggResult.Record().GetByIndex(0), aggResult.Record().GetByIndex(1))
    }

    if err := aggResult.Err(); err != nil {
        log.Fatal(err)
    }
}
