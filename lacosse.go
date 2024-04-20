CREATE (:Player {name: 'Player A', position: 'Attack'})
CREATE (:Player {name: 'Player B', position: 'Midfield'})
CREATE (:Player {name: 'Player C', position: 'Defense'})
CREATE (:Player {name: 'Player D', position: 'Goalkeeper'})
CREATE (:Game {date: '2023-04-01', opponent: 'Team X'})


### Low Submarine Shots and Rainbow Passes
MATCH (p1:Player {name: 'Player A'}), (g:Game {date: '2023-04-01'})
CREATE (p1)-[:LOW_SUBMARINE_SHOT {outcome: 'goal'}]->(g)

MATCH (p2:Player {name: 'Player B'}), (g:Game {date: '2023-04-01'})
CREATE (p2)-[:RAINBOW_PASS {outcome: 'complete'}]->(p1)

### Faceoffs, Turnovers, and Saves
MATCH (p3:Player {name: 'Player C'}), (g:Game {date: '2023-04-01'})
CREATE (p3)-[:FACEOFF_WON]->(g)

MATCH (p1:Player {name: 'Player A'}), (g:Game {date: '2023-04-01'})
CREATE (p1)-[:TURNOVER {cause: 'interception'}]->(g)

MATCH (p4:Player {name: 'Player D'}), (g:Game {date: '2023-04-01'})
CREATE (p4)-[:SAVE {shotBy: 'Opponent Player'}]->(g)


