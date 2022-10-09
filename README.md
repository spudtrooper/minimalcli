# minimalcli

Minimal CLI framework for go.

## tl;dr

Why use this? [See here](https://github.com/spudtrooper/minimalcli/blob/main/use-case.md).

Example projects using this:

  * [github.com/spudtrooper/apidumpsterfire](https://github.com/spudtrooper/apidumpsterfire)
  * [github.com/spudtrooper/opentable](https://github.com/spudtrooper/opentable)
  * [github.com/spudtrooper/resy](https://github.com/spudtrooper/resy)


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