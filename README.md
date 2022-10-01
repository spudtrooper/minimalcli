# minimalcli

Minimal CLI framework for go.

## Details

The CLI will take commands as full names or abbreviations before all the flags. Command are registered with `Register`. Help text looks something like this:

```
                    Action - Abbreviation
  ======================== - ============
                      Auth - a
                       Bid - b
                      Bids - bi
                      Help - h
                   Resolve - r
```

So, with this example you could invoke it as:

```
<app-name> Auth --some_flag --another_flag arg-1 arg-2
<app-name> auth --some_flag --another_flag arg-1 arg-2
<app-name> a --some_flag --another_flag arg-1 arg-2
```

## Usage

* For a simple CLI see `app/app_test.go`.
* For a combination of CLI and HTTP, see the following example:
  * Define a set of [Handlers](https://github.com/spudtrooper/opentable/blob/main/handlers/handlers.go)
  * Use those handlers to create a [CLI](https://github.com/spudtrooper/opentable/blob/main/cli/main.go)
  * Use the same handlers to create an [HTTP server](https://github.com/spudtrooper/opentable/blob/main/frontend/server.go)

