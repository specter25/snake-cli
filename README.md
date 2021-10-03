# Snake Cli 


## Terminal-based Snake game

It took me approx 5 hours to build the game . I have used two external libraries [termloop](github.com/JoelOtter/termloop) and [termbox](github.com/nsf/termbox-go) .The instructions to start the game are given below . I have formatted the code using goimports and golang-cilint. 

## Assumptions

- The minimum dimensions for the game to run is 70 pxels width and 25 pixels height .

<!-- ![scrrenshot](http://i.imgur.com/pHf4fjt.gif) -->

## Play

### Locally

```
$ go get github.com/specter25/snake-cli
$ $GOPATH/bin/snake-cli
```

### On Docker

```
$ docker run -ti ujjwal25/snake-cli
```
