# A dumb generals.io bot in Go

~~It doesn't even do random moves yet.~~ It totally runs a random strategy.

But does have most of the logistics wired in.

```bash
# just prep it
git clone https://github.com/bryanhelmig/generalsbot
cd generalsbot
go get

# just test it
go test

# just run it
go run main.go core.go constants.go strategy.go types.go
# now join the printed URL on the screen
```

Forgive the Go newbishness.

### Features

While it runs, it prints the grid to stdout:

![](https://cdn.zapier.com/storage/photos/142c31ba28fececb40b4e6b49c2d325c.png)

![](https://cdn.zapier.com/storage/photos/8d419bbdc24fd632d62a4e968b4587f3.png)

### Upcoming

Pluggable strategies with preferences game stage. For example:

1. No enemies seen yet? `Explore` strategy will want to make the moves.
2. Seen an enemy but haven't seen seen their general? `Hunt` strategy should be preferred.
3. Seen a general? `BuildThenRush` strategy will be preferred.
4. Enemy seen us and still alive? `Defend` strategy will be preferred.
5. etc...
