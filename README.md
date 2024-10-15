# Project Instructions

## To run the project

You can run the project using either of the following methods:

1. **Make**  
   ```bash
   make
   ```

2. **Docker Compose**  
   ```bash
   docker-compose up -d --build
   ```
![image](https://github.com/user-attachments/assets/7243818e-f11a-4bb7-a9fb-5dd2116e7343)

## Collection Schema

The collection schema for storing match-related data includes the following fields:

- **HomeTeam**
- **AwayTeam**
- **MatchDate**
- **Events** (All events are stored here during the match)

## Event Types

The following event types are defined for a match:

1. **MatchStarted**  
   Raised when the match starts.
   
2. **MatchFinished**  
   Raised when the match ends.

3. **Score**  
   Raised when a score happens.

### Possible Future Events

Other events can be implemented easily as needed. For example:

1. **MatchStopped**  
   Raised when the match is paused.

2. **PlayerSubstituted**  
   Raised when a player is substituted.

...and more.
