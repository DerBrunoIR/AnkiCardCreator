# AnkiCardCreator

This repo contains a script to extract anki cards from go standard libary.
AnkiConnect addon is necessary.

## [Deck]( https://ankiweb.net/shared/decks?search=GoLang%20Standard%20Libary)
version: Go@1.21.1 Hi,

I created this deck mainly for personal use. This deck contains all packages (even internal ones) from https://pkg.go.dev/std. For each package i genrated for all sub packages and for each variable block, constant block, function and type a single card by a carful selection of html elements from the standards libaries webpage. All links should work. For easy access i mirrored the hirachical structure from the standard libary.

statistics:

- min = 0 cards/deck (e.g. Go::StdLib::log::slog::internal::benchmarks)
- max = 290 cards/deck (Go::StdLib::syscall)
- std = 23.52 cards/deck
- avg = 13.35 cards/deck
- 260 diffrent decks
- in total 3471 cards, seems like a lot, however the average is pretty low
- 
random example cards ( since the web preview seems not to work with HTML cards ):
- hash.adler32.New front: https://pasteboard.co/H3FqayVfVYcP.png
- back: https://pasteboard.co/5AK8T4914V8u.png
sync.atomic.Uintptr
- front: https://pasteboard.co/SaHUXpEXCGlk.png
- back: https://pasteboard.co/wAIGyiVvMRGW.png

This deck contains content published under a BSD license from "https://go.dev/".
