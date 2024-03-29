
# toggl.com task

## Version
This project uses the go 1.21, this is the latest version as of Feb 7th, 2024.

```bash
$ go version
go version go1.21.6 darwin/arm64
```

## Timeline & Plans (How the time was used):
- Revised GoLang
- Got a look on card games, I don't play such games (used to play anime card games only), so no clue about what's common between such games (still no clue)
- Spent sometime reading GoLang current (2024) conventions in:
  - Naming variables
  - Naming files
  - Directory structuring
  - Think of a good directory structure that tells what it does without opening the files, eventually used a custom MVC
  - Search for a linter for go, to mention its importance (Even though I'm not going to use it for the sake of the demo)


## Notes:
- Some FIXME are put where improvements can happen
- MVC was used as a pattern, models are in the models directory, models DO NOT interact with the DB, they are just models,
all the DB access happens in controllers only.
- Controllers are function of `gin.context` which is the request wrapper.
- Services are basically helpers for models or external APIs if needed.
- logging includes a logger in the awesome `zap` go library
- common has the shared stuff like error codes, utility functions, some consts used across the project
- The project configuration is in .env
- Each function that does any query have the expected query count in comments, this is not tested, but should be

## Run the project
- CD into the project
- `go install`
-  `go run .`


## Endpoints as described by GIN in debug mode
```bash
[GIN-debug] GET    /ping                     --> togglDecks/routers.SetupMainRouter.func1 (3 handlers)
[GIN-debug] GET    /decks                    --> togglDecks/controllers/deck_controllers.GetAvailableDecks (3 handlers)
[GIN-debug] POST   /decks                    --> togglDecks/controllers/deck_controllers.NewDeck (3 handlers)
[GIN-debug] PUT    /decks/:id/draw           --> togglDecks/controllers/deck_controllers.DrawCards (3 handlers)
[GIN-debug] GET    /decks/:id                --> togglDecks/controllers/deck_controllers.OpenDeck (3 handlers)
```

## Docs 
- Open [swagger editor](https://editor.swagger.io/) and paste swagger.yaml to view, this uses the old swagger V2 because their online editor uses it by default. 

## Testing

Run `go test ./...`

Opinion: I found testing in go easy but lacks a good test suite, for example, In Python we have Pytest, in JS/TS we have Jest (and others), both handle everything, I didn't find a mature way to share fixtures
between ALL tests in the project, only per-package which is `TestMain(m *testing.M)` and it cann't be in the root (main) package. I think Go will benefit from a test-suit that handles everything. There's a way I could think of (using Reflection) that can do this, but since this is super out of scope, I intentionally ignored it.

Also, I did write a few tests to show a demo of how I write tests, feel free to tell me how to improve them, but also, I could think of way more tests but let's leave 'em for the intervew :)


## Disclaimer: 
1. This has things that are out of scope of the task required, I have made much more, it's fine for me, I used the task to revise Golang.
2. I revised lots of code in the last few days and I'm not sure if I'm following the best practises since I haven't worked with go for the last 1.5y.
3. I would appreciate constructive feedback in case of declination.

## Thanks!
I would like to thank you for the task, I did really want to revise GoLang but I didn't have the time, the interest in toggl forced me to put time for it.
