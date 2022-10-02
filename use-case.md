# minimalcli use case

Why use this? If you're creating an API and you want to generate a CLI and HTTP server with minimal effort while you iterate.

## tl;dr

Create https://unofficial-opentable-api.herokuapp.com/, which is like http://opentable.herokuapp.com/api but more, with next-to-no additional effort than just calling that API.

## Overview

I reverse-engineered opentable's API, because I wanted to scrape their NYC data to answer questions like:

* [What is the most frequent menu item?](https://github.com/spudtrooper/opentable/blob/main/output/menu-item-histogram/index.md) - spoiler, it's Caeser Salad
* [What is the most expensive menu item?](https://github.com/spudtrooper/opentable/blob/main/output/sort-by-price/index.md) - spoiler, it's all booze

Why not use http://opentable.herokuapp.com/api? It doesn't expose menu items and I like doing things like this.

## The Process

I start by converting curl requests in chrome dev console to code as explained [here](https://spudtrooper.github.io/articles/fromcurltogo/). As I convert a request, add it to  [api/core.go](https://github.com/spudtrooper/opentable/blob/main/api/core.go). When I create "derived" functionality, add that to [api/extended.go](https://github.com/spudtrooper/opentable/blob/main/api/extended.go). e.g. `Search` would go in *core* and `SearchAll` that makes a bunch of `Search` calls goes in *extended*.

As I iterate, I like to keep a CLI version going so I can tinker--you could also just write tests. [Here](https://github.com/spudtrooper/opentable/blob/d0e34fba56619538709d51a2aa57b253b91e3294/cli/main.go) is an example of what an itermediate state looked like.

But I wanted something like http://opentable.herokuapp.com/api without any effort. So, I added the [handler package](https://github.com/spudtrooper/minimalcli/tree/main/handler) (terrible name!) to allow you to get a CLI and little HTTP server with about the same effort. So, instead of hard-coding calls to the client in the CLI, I split it up this way:
  * [handlers/handlers.go](https://github.com/spudtrooper/opentable/blob/main/handlers/handlers.go) has the calls that translate either flags or request params to calls to the API. This is what we previously in [cli/main.go](https://github.com/spudtrooper/opentable/blob/d0e34fba56619538709d51a2aa57b253b91e3294/cli/main.go)
  * [cli/main.go](https://github.com/spudtrooper/opentable/blob/main/cli/main.go) uses the handlers plus a little boilerplate to generate a CLI
  * [frontend/server.go](https://github.com/spudtrooper/opentable/blob/main/frontend/server.go) uses the handlers plus a little boilerplate to generate a little HTTP server.

End result: https://unofficial-opentable-api.herokuapp.com/ with ~no additional effort.