package main

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"

	"poker-go-cli/game"
	"poker-go-cli/game/player"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type GameUI struct {
	window     fyne.Window
	game       *game.Game
	player     *player.Player
	betEntry   *widget.Entry
	handLabel  *widget.Label
	potLabel   *widget.Label
	stateLabel *widget.Label
}

func main() {
	pokerApp := app.New()
	window := pokerApp.NewWindow("Poker Game")
	window.Resize(fyne.NewSize(800, 600))

	// Create welcome screen
	welcomeText := canvas.NewText("Welcome to Poker Game", color.White)
	welcomeText.TextSize = 24
	welcomeText.TextStyle.Bold = true

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter your name")

	chipsEntry := widget.NewEntry()
	chipsEntry.SetPlaceHolder("Enter starting chips")

	startButton := widget.NewButton("Start Game", func() {
		if nameEntry.Text == "" || chipsEntry.Text == "" {
			dialog.ShowError(
				errors.New("Please enter your name and starting chips"),
				window,
			)
			return
		}

		chips, err := strconv.Atoi(chipsEntry.Text)
		if err != nil {
			dialog.ShowError(
				errors.New("Please enter a valid number for chips"),
				window,
			)
			return
		}

		// Create game UI
		gameUI := &GameUI{
			window: window,
			player: player.NewPlayer(nameEntry.Text, chips),
		}

		// Create AI players
		aiPlayers := []*player.Player{
			player.NewPlayer("AI 1", 1000),
			player.NewPlayer("AI 2", 1000),
			player.NewPlayer("AI 3", 1000),
		}

		// Create game with all players
		allPlayers := append([]*player.Player{gameUI.player}, aiPlayers...)
		gameUI.game = game.NewGame(allPlayers)

		// Show game screen
		showGameScreen(window, gameUI)
	})

	content := container.NewVBox(
		welcomeText,
		widget.NewLabel("Player Name:"),
		nameEntry,
		widget.NewLabel("Starting Chips:"),
		chipsEntry,
		startButton,
	)

	window.SetContent(content)
	window.ShowAndRun()
}

func showGameScreen(window fyne.Window, gameUI *GameUI) {
	// Create game UI elements
	gameUI.handLabel = widget.NewLabel("Your Hand: ")
	gameUI.potLabel = widget.NewLabel("Pot: 0")
	gameUI.stateLabel = widget.NewLabel("Game State: Pre-Flop")
	gameUI.betEntry = widget.NewEntry()
	gameUI.betEntry.SetPlaceHolder("Enter bet amount")

	// Create game controls
	dealButton := widget.NewButton("Deal Cards", func() {
		gameUI.game.DealCards()
		updateGameUI(gameUI)
	})

	flopButton := widget.NewButton("Deal Flop", func() {
		gameUI.game.DealFlop()
		updateGameUI(gameUI)
	})

	turnButton := widget.NewButton("Deal Turn", func() {
		gameUI.game.DealTurn()
		updateGameUI(gameUI)
	})

	riverButton := widget.NewButton("Deal River", func() {
		gameUI.game.DealRiver()
		updateGameUI(gameUI)
	})

	betButton := widget.NewButton("Place Bet", func() {
		bet, err := strconv.Atoi(gameUI.betEntry.Text)
		if err != nil {
			dialog.ShowError(errors.New("Please enter a valid bet amount"), window)
			return
		}

		if gameUI.player.PlaceBet(bet) {
			gameUI.game.Pot += bet
			updateGameUI(gameUI)
		} else {
			dialog.ShowError(errors.New("Not enough chips"), window)
		}
	})

	// Layout game UI
	gameContent := container.NewVBox(
		gameUI.handLabel,
		gameUI.potLabel,
		gameUI.stateLabel,
		widget.NewLabel("Your Chips: "+strconv.Itoa(gameUI.player.Chips)),
		gameUI.betEntry,
		betButton,
		dealButton,
		flopButton,
		turnButton,
		riverButton,
	)

	window.SetContent(gameContent)
}

func updateGameUI(gameUI *GameUI) {
	// Update hand display
	handStr := "Your Hand: "
	for _, card := range gameUI.player.Hand {
		handStr += card.String() + " "
	}
	gameUI.handLabel.SetText(handStr)

	// Update pot
	gameUI.potLabel.SetText(fmt.Sprintf("Pot: %d", gameUI.game.Pot))

	// Update game state
	stateStr := "Game State: "
	switch gameUI.game.State {
	case game.PreFlop:
		stateStr += "Pre-Flop"
	case game.Flop:
		stateStr += "Flop"
	case game.Turn:
		stateStr += "Turn"
	case game.River:
		stateStr += "River"
	case game.Showdown:
		stateStr += "Showdown"
	}
	gameUI.stateLabel.SetText(stateStr)
}