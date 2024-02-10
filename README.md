
## toggl.com task

This project uses the go 1.18 since it's the latest version as of 

Timeline & Plans (How the time was used):
- Revised GoLang
- Got a look on card games, I don't play such games (used to play anime card games only), so no clue about what's common between such games (still no clue)
- Spent sometime reading GoLang current (2024) conventions in:
  - Naming variables
  - Naming files
  - Directory structuring
  - Think of a good directory structure that tells what it does without opening the files, eventually used a custom MVC
  - Search for a linter for go, to mention its importance (Even though I'm not going to use it for the sake of the demo)

Disclaimer:
- I revised lots of code in the last few days and I'm not sure if I'm following the best practises since I haven't worked with go for the last 1.5y.

Conventions used:
- Some FIXME are put where improvements can happen


- MVC was used as a pattern, models are in the models directory, models DO NOT interact with the DB, they are just models,
all the DB access happens in controllers only.
- Controllers are function of `gin.context` which is the request wrapper.
- Services are basically helpers for models or external APIs if needed.
- logging includes a logger in the awesome `zap` go library
- common has the shared stuff like error codes, utility functions, some consts used across the project
- The project configuration is in .env

Notes:
- Each function that does any query have the expected query count in comments, this is tested
- 

## Run the project
- CD into the project and run `go run .`


## Endpoints as described by GIN in debug mode
```bash
[GIN-debug] GET    /ping                     --> togglDecks/routers.SetupMainRouter.func1 (3 handlers)
[GIN-debug] GET    /decks                    --> togglDecks/controllers/deck_controllers.GetAvailableDecks (3 handlers)
[GIN-debug] POST   /decks                    --> togglDecks/controllers/deck_controllers.NewDeck (3 handlers)
[GIN-debug] PUT    /decks/:id/draw           --> togglDecks/controllers/deck_controllers.DrawCards (3 handlers)
[GIN-debug] GET    /decks/:id                --> togglDecks/controllers/deck_controllers.OpenDeck (3 handlers)
```

## Docs 
- Open [swagger editor](https://editor.swagger.io/) and paste swagger.yaml to view


Disclaimer: 
1. This is out of scope of the task required, I have made much more, it's fine for me, I used the task to revise Golang.
2. Anything missing is intentionally ignored, I didn't want to make it more late than this
