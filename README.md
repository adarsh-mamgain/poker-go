# Poker Go

A command-line poker game implementation in Go with a graphical user interface built using Fyne.

## Features

- Interactive poker game with a modern GUI
- Support for multiple AI players
- Realistic poker gameplay including:
  - Pre-flop, flop, turn, and river phases
  - Betting system
  - Hand evaluation
  - Pot management
- Customizable starting chips
- Player name customization

## Prerequisites

- Go 1.16 or higher
- Fyne GUI toolkit dependencies

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/poker-go.git
cd poker-go
```

2. Install dependencies:

```bash
go mod download
```

3. Run the game:

```bash
go run main.go
```

## How to Play

1. Launch the game
2. Enter your name and starting chips
3. Click "Start Game" to begin
4. Use the game controls to:
   - Deal cards
   - Place bets
   - Progress through game phases (flop, turn, river)
   - View your hand and the pot

## Project Structure

```
poker-go/
├── main.go          # Main application entry point
├── game/            # Core game logic
│   └── player/      # Player-related functionality
├── go.mod           # Go module definition
└── go.sum           # Go module checksums
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
