# Bees In The Trap 🐝

You must destroy a hive of bees before they sting you to death.

## Game Rules

- **Your Stats:**
  - Starting HP: 100
  - Goal: Destroy the hive before you die

- **The Hive Contains:**
  - 1 Queen Bee (100 HP, deals 10 damage)
  - 5 Worker Bees (75 HP each, deal 5 damage each)
  - 25 Drone Bees (60 HP each, deal 1 damage each)

- **Gameplay Mechanics:**
  - Take turns attacking the hive
  - Random bee selection for attacks
  - 20% chance to miss your attack
  - 15% chance for bees to miss their sting
  - If the Queen dies, all other bees die too!

## Features

- Turn-based strategic gameplay
- Colorful terminal interface
- Random targeting system
- Auto-play mode with easy toggle
- Detailed game status display
- End-game summary
- Concurrent processing using goroutines
- Cross-platform support (Windows, Linux, macOS)

## Technical Implementation

- Written in Go 1.18
- Uses standard library only (no external dependencies)
- Implements concurrent programming patterns
- Clean code architecture with separation of concerns
- Comprehensive test coverage
- Cross-platform compatibility

## Installation

### Prerequisites

- Go 1.18 or higher
- Git

### Getting Started

```bash
# Clone the repository
git clone https://github.com/chichi-ohio/BeesInTheTrap.git
cd BeesInTheTrap

# Build for your platform
make build

# Run tests
make test
```

## Running the Game

### Using Make Commands

```bash
# For macOS
make run-mac

# For Linux
make run

# For Windows (run the executable directly)
./bin/bees-in-the-trap-windows.exe
```

### Game Commands

- `hit` - Attack the beehive
- `auto` - Toggle auto-play mode (press Enter to stop)
- `status` - Display current game status
- `exit` - Quit the game

## Testing

Run the comprehensive test suite:

```bash
make test
```

## Project Structure

```
BeesInTheTrap/
├── cmd/
│   └── game/          # Main game entry point
├── internal/
│   └── game/          # Game logic and components
├── bin/               # Compiled binaries
├── Makefile          # Build and run commands
└── README.md         # This file
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Game Preview

```
=== BEES IN THE TRAP ===
Turn: 1

Player Status:
  Health: 100/100 HP
  Hits: 0

Hive Status:
  Queen Bees:  1/1
  Worker Bees: 5/5
  Drone Bees:  25/25
``` 
