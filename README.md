# minimalcli

Minimal CLI framework for go.

## tl;dr

This is a framework that allows you iterate on an API and easily create a CLI and HTTP server consistently. It essentially lets you migrate from a library to a service. So, it serves the same purpose as gPRC and client libraries,  but you write the library first with little scaffolding. The premise is that this order is more natural for iteration.

Example projects using this:

  * [github.com/spudtrooper/apidumpsterfire](https://github.com/spudtrooper/apidumpsterfire)
    * Deployed: [api-dumpster-fire.herokuapp.com/](https://api-dumpster-fire.herokuapp.com/)
  * [github.com/spudtrooper/opentable](https://github.com/spudtrooper/opentable)
    * Deployed: [unofficial-opentable-api.herokuapp.com/](https://unofficial-opentable-api.herokuapp.com/)
  * [github.com/spudtrooper/opensecrets](https://github.com/spudtrooper/opensecrets)
  * [github.com/spudtrooper/resy](https://github.com/spudtrooper/resy)
    * Deployed: [unofficial-resy-api.herokuapp.com/](https://unofficial-resy-api.herokuapp.com/)
  * [github.com/spudtrooper/scplanner](https://github.com/spudtrooper/scplanner)

More details on why? [See here](https://github.com/spudtrooper/minimalcli/blob/main/use-case.md).

## Details

### CLI

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