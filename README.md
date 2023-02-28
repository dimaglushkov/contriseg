# contriseg
Contriseg allows anyone to generate animations with their github contributions charts with an extreme ease. 

<img src="https://dimaglushkov.xyz/static/contriseg_move.gif" width="100%" >
<img src="https://dimaglushkov.xyz/static/contriseg_bfs.gif" width="100%" >
<img src="https://dimaglushkov.xyz/static/contriseg_cbc.gif" width="100%" >

### Installation & Launching (without docker)
0. Install `go`
1. Clone or download this repo: 
```bash
git clone git@github.com:dimaglushkov/contriseg.git
```
2. Generate your [GitHub Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-personal-access-token-classic) with only `read:user` privileges
3. Create your `.env` file from the [`.env.example`]((https://github.com/dimaglushkov/contriseg/blob/main/.env.example)) with the following content:
```bash
GITHUB_USERNAME=<target user>
GITHUB_TOKEN=<previously generated token>
TARGET_LOCATION=<output file location>
ANIMATION=<animation type ()predefined: bfs, cbc, move)>
```
4. You need to load above env variables before running contriseg. You can use something like [python-dotenv](https://github.com/theskumar/python-dotenv/).
5. Download dependencies, build, and run the application:
```bash
go run .
```

### Compressing resulting GIF
Since a generated GIF could be big and take too long to be downloaded, you can compress it using amazing [`gifsicle`](https://github.com/kohler/gifsicle).
To do that just run
```bash
./bin/compress.sh
```

### Creating custom animations
To create a new animation you only need to do two things:
1. Implement function of type 
```go
type Iterator func(cal internal.Calendar) []internal.Calendar
```
which should return a list of continuously changing `Calendar` (it's basically frames for Calendar).

Check out [`anim.go`](https://github.com/dimaglushkov/contriseg/blob/main/internal/image/anim.go) to see predefined animations.


2. After animation is completed, insert it's alias and name to [iterationsMap](https://github.com/dimaglushkov/contriseg/blob/main/internal/image/anim.go#L12)
```go
var animationsMap = map[string]AnimationIterator{
	"bfs":  CalendarBFSIterations,
	"move": CalendarMoveColLeftIterations,
	"cbc":  CalendarColByColIterations,
	"your_alias": YourAnimationFunction
}

```

That's it, now you can use your own animation by just editing `.env` file and changing `ANIMATION` value to the alias of the recently created animation 
