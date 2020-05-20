# Overview

This part of the application will comprise of services responsible for getting the information of a team.

## Major Use Cases

- Get names of the team members
- Get handles of the team members
- Get team name
- Get list of problems solved by the team members combined
- Get topic-wise list of problems solved by the team
- Rate a team by topic-wise.
- Get a team's top 5 strongest topics
- Get a team's top 5 weakest topics
- Get a team's percentage of AC, WA, TLE, MLE, Compilation Error.
- Get a team's list of attempted unsolved problems
- Get a team's estimated Codeforces rating.

## Objects in the system

- Team
- Individual
- Problem
- Topic
- Verdicts

## Relations between the classes

- Team to Individual: Many to Many
- Team to Problem: Many to Many
- Team to Topic: Many to Many
- Team to Verdicts: Many to Many
- Individual to Problem: Many to Many
- Individual to Topic: Many to Many
- Individual to Verdicts: Many to Many
- Topic to Problem: One to Many
- Problem to Verdicts: Many to Many

## Classes design

- Team:
  - Attributes:
    - teamID
    - teamName
    - teamMembers
    - listOfProblems
    - topicWiseSolved
    - topicWiseTotal
    - acceptedProblems
    - teamRating
    - strongestTopics
    - weakestTopics
  - Methods:
    - getTeamName
    - getTeamMembers
    - getTeamID
    - getListOfProblems
    - getTopicWiseStatistics
    - getTeamRating
    - getAcceptedProblems
    - getStrongestTopics
    - getWeakestTopics
    - calculateRating
    - calculateStrongestAndWeakest
