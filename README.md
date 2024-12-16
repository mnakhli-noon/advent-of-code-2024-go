# Advent of Code 2024 - Go Learning Journey 🎄🧊
## About This Repository
Welcome to my Advent of Code adventure! This repository is a personal learning project where I'm diving into the Go programming language by solving the Advent of Code puzzles.

### My Learning Approach 🚀
- **Day-by-Day Learning**: I'll be solving each puzzle first in my own way
- **Collaborative Improvement**: After my initial solution, I'll share my code with [Claude](https://www.claude.ai)
- **Go Language Mastery**: Seeking guidance to understand the "Go way" of solving problems

## Repository Structure
```
advent-of-code-2024/
│
├── day-01/
│   ├── solution.go      # My initial solution
│   ├── improved.go      # Refined solution after Claude's feedback, or online resources -- might not exist
│   ├── problem.md       # AoC problem description
│   └── input.txt        # Input file
│
├── day-02/
│   └── ...
└── README.md
```

## Building and Running Solutions
Each day's solution can be built and run with optional flags:

```bash
# Build and run Part One (default)
go run solution.go

# Build and run Part Two
go run solution.go -partTwo

# Use the improved solution (if improved.go exists)
go run solution.go improved.go -improved

# Use improved solution for Part Two
go run solution.go improved.go -improved -partTwo
```

### Flags
- `-partTwo`: Switches the solution to solve Part Two of the puzzle
  - When this flag is passed, the program will typically use a different logic or input processing method for Part Two of the puzzle
- `-improved`: Uses the improved solution from `improved.go` (if the file exists)

## Learning Goals
- [ ] Understand Go's syntax and idioms
- [ ] Learn best practices in Go programming
- [ ] Improve problem-solving skills
- [ ] Complete as many Advent of Code puzzles as possible

## How This Works
1. I solve the puzzle using my current knowledge
2. I consult Claude.ai for:
   - Code review
   - Idiomatic Go improvements
   - Learning opportunities
3. I may use online resources for the improved version of the solution.

## Resources
- [Advent of Code](https://adventofcode.com/)
- [Go Official Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)

## Disclaimer
This is a learning journey. Expect messy code, lots of questions, and gradual improvement! 🌱

## Contact
Feel free to follow along, provide suggestions, or share learning tips!

---
*Happy Coding and Happy Holidays!* 🎅🏼👩‍💻

*PS: this README file is AI generated*
