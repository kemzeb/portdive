<p align="center">
<img src="https://raw.githubusercontent.com/kemzeb/portdive/main/assets/portdive.gif" height = 458 width = 696>
</p>

---

portdive is a CLI-based game inspired by an old classic video game called [Splinter Cell: Chaos Theory](https://en.wikipedia.org/wiki/Tom_Clancy%27s_Splinter_Cell:_Chaos_Theory). More specifically, the hacking mini-game players were able to play. In its current state, the game is playable but in active development.

The objective of the game is to find the correct IP address using the "Pwner" device to reduce the number of options that you have to choose from. More features are to come!

The game tutorials define these addresses as "port" addresses. Though I believe this is not correct [(port addresses are numbers between 0-65,535)](https://en.wikipedia.org/wiki/Port_(computer_networking)#Port_number), I decided to honor this mistake by including it in the name of the project. However, the minigame may be referring to attempting to exploit a service known as [port fowarding](https://en.wikipedia.org/wiki/Port_forwarding#Purpose).

## To-Do
- [ ] Reduce tight coupling and other major refactoring work
- [ ] Update game UI to be more appealing
- [ ] Add more features that were seen in the original minigame
- [ ] Introduce a mechanism to generate IP addresses such that the game is still winnable (rather than using a static set of addresses)
- [ ] Ponder introducing a round system where finding the correct IP address becomes progressively harder

## Installation
Simply clone this repository using `git clone` or use Github's ZIP download feature.

## Usage
To play the game, make sure you are in the root directory of the project. Afterward, execute the following command into your command-line session:

` go run main.go`

Assuming that you have an installation of Go on your machine, you should be able to start playing the game!

### Keys
Movement within the IP address list and the Pwner device is done by holding `Control` and using basic ***Vim*** keybinds of movement within  text:

Move ***UP*** in the IP address list:   `Control+J`

Move ***DOWN*** in the IP address list: `Control+K`

Move ***RIGHT*** in the Pwner device:   `Control+L`

Move ***LEFT*** in the Pwner device:    `Control+H`


To make your choice with either the address list or the Pwner, you hold `Control` and also use the following:

Make a decision in the ***address list***: `Control+M`

Make a decision in the ***Pwner device***: `Control+P`

To exit the game, use `Control+C`

The reason behind the use of `Control` is due to the nature of the keybinds available with the ***termloop*** API.

