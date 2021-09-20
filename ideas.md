## What we need:

1. A way to remove the total match table
   1. **DONE**!
   2. Add back in when we get agents
      1. Player agents are missing from our models
2. A way to link the map name and the round scores to the player data
3. A way to link the time-series to everything else
4. A way to extract the time-series
   1. **DONE**
      1. A way to identify who is defense/attack
5. A way to extract the econ 
   1. **DONE**
6. A way to link the econ to the time-series
7. A way to extract the match data
   1. **DONE**
   2. A way to identify who is defense/attack first
8. A way to extract links from the event page (i.e. event > matches list)

9. Hierarchy of data / files
   1. Match
      1. Teams
      2. Maps
      3. Match score
      4. Total player scores (can extract)
   2. Game
      1. Multiple for each match (BO3 or BO5)
      2. Game score
         1. Time series & final score
            1. TS to include:
               1. round wins/losses
               2. round result type (bomb, frag, time)
               3. merged with ECONOMY rounds
                  1. ASSUME this is load-out value not $ spent
10. A way to extract map choice
11. A way to get the hidden data from one page view
    1. is economy hidden or on another page?


## Analytical ideas:
- Add deaths into ACS calculation
  - Improve the formulation regarding first blood / first death
- time series predictive analysis
  - i.e. model can predict potential outcome round by round
    - should be 100% accurate if one team reaches 13 rounds
    - pistols (i.e. round 1 and 13) should be huge factors?
- player clustering
  - analyzing average player performance to showcase when they "pop off"
- We will probably need a parser to develop team histories... i.e. like Veritas (won/lost, focus/opp)

## Refactoring:
[] fix types to solid int for player stats but maintain float64