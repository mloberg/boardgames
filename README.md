# Board Games

Picking a board game to play can be challenging. Figuring out what games work
with the number of people you have is difficult enough, but then factor in how
much time and brain space you have. Often picking the game ends up being the most
difficult part of the night.

This grabs my collection from [BoardGameGeek](https://boardgamegeek.com/) and
provides a web page to filter and search for board games.

## Requirements

* NodeJS 16
* [pnpm](https://pnpm.io/)
* Golang 1.17
* [Hugo](https://gohugo.io/)

## Configuration

* `USERNAME` BoardGameGeek username to index
* `MEILISEARCH_HOST` Meilisearch host URL (_default_: `http://localhost:7700`)
* `MEILISEARCH_API_KEY` Meilisearch private key used for indexing
* `HUGOxPARAMSxMEILISEARCH_HOST` Meilisearch host used for frontend
* `HUGOxPARAMSxMEILISEARCH_API_KEY` Meilisearch public key used for frontend.

## Usage

### Indexer

You can build the cli with `make cli`. This will create a `bin/bg` which you can
call and use as you need.

To setup the index (or make changes to it), run `make setup`.

To update the index, run `make index`.

### Frontend

After installing packages with `pnpm install`, serve the site and rebuild changes
with `make serve`. Build the production site (with assets minified) with `make site`.
Deploy the site `make deploy`.

### Development

Lint code with `make lint`. This lints both JavaScript code (requires `pnpm install`)
and Go code (requires `golangci-lint`).
