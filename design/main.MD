# System Design

## Major Use Cases

- As a user, I would like to be able to input my team's details from Codeforces, and gather information about our performance in prior contests.
- As a user, I would like to be able to see my individual strength and weaknesses.
- As a team, I would like to be able to see me and my teammates' weaknesses as a whole.
- As a team, we would like to see suggested problems and contests for us to practice on Codeforces.
- As a team, we would like to be able to collaborate code over a realtime online text editor, for us to easily resolve bugs as a team during contests.
- Team practice backlog should be maintained.

## Assumptions

- The application can be accessed by ~30 users at a time.
- At a single instant, around ~15 individual and team information would be received.
- At a single instant, around ~15 team suggestion requests would be received.
- At a single instant, around ~10 team collaboration instances could be received.

## Components

- Cloudfare DNS
- Frontend servers
- Team Information Servers
- Team Suggestion Servers
- Collaboration servers

