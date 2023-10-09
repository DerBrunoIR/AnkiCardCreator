# AnkiCardCreator

This repo is meant to contain scripts to extract anki cards from different sources. At the moment only one source is supported.

# Sources

### pkg.go.dev (GoLang package documentation)
AnkiConnect addon is necessary.


# Public Decks

### StdLib 

available at: [AnkiWebLink](https://ankiweb.net/shared/info/1050662530)
version: Go@1.21.1 

Hi,

I created this deck mainly for personal use. This deck contains all packages (even internal ones) from https://pkg.go.dev/std. For each package i genrated for each variable div, constant div, function div and type div a single card by a carful selection of html elements from the standards libaries webpage. All links should work. For easy access i mirrored the hirachical structure from the standard libary.

statistics:

- min = 0 cards/deck (e.g. Go::StdLib::log::slog::internal::benchmarks)
- max = 290 cards/deck (Go::StdLib::syscall)
- std = 23.52 cards/deck
- avg = 13.35 cards/deck
- 260 diffrent decks
- in total 3471 cards

## random example
  
`sync.atomic.Uintptr` 

- ![Front](https://github.com/DerBrunoIR/AnkiCardCreator/assets/95578637/0a21eb67-07e0-461e-957e-ef959b949cd1)
- ![Back](https://github.com/DerBrunoIR/AnkiCardCreator/assets/95578637/a2278989-fc10-4584-9ae3-908c01b633d6)


This deck contains content published under a BSD license from "https://go.dev/".
