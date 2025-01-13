package main
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// SolveSeatingArrangement solves seating based on constraints
func SolveSeatingArrangement(arrangementType string, n int, constraints []map[string]string) ([]string, error) {
	// Validate inputs
	if arrangementType != "linear" && arrangementType != "circular" {
		return nil, errors.New("invalid arrangement type: must be 'linear' or 'circular'")
	}
	if n <= 0 {
		return nil, errors.New("number of people must be greater than zero")
	}

	// Initialize the seating arrangement
	seats := make([]string, n)

	// Apply constraints
	for _, constraint := range constraints {
		person := constraint["person"]
		nextTo := constraint["next_to"]
		side := constraint["side"] // "left" or "right"

		// Find the position of the referenced person
		pos := findPosition(seats, nextTo)
		if pos == -1 {
			// If the referenced person is not yet seated, place them
			pos = placeFirstAvailable(seats, nextTo)
		}

		// Place the current person based on the side
		if side == "left" {
			placeAt(seats, (pos-1+n)%n, person, arrangementType == "circular")
		} else if side == "right" {
			placeAt(seats, (pos+1)%n, person, arrangementType == "circular")
		} else {
			return nil, errors.New("invalid side: must be 'left' or 'right'")
		}
	}

	// Fill remaining empty seats
	for i := 0; i < n; i++ {
		if seats[i] == "" {
			seats[i] = fmt.Sprintf("Person%d", i+1) // Fill remaining seats with placeholders
		}
	}

	return seats, nil
}

// Helper: Find the position of a person in the seating arrangement
func findPosition(seats []string, person string) int {
	for i, p := range seats {
		if p == person {
			return i
		}
	}
	return -1
}

// Helper: Place a person in the first available seat
func placeFirstAvailable(seats []string, person string) int {
	for i, seat := range seats {
		if seat == "" {
			seats[i] = person
			return i
		}
	}
	return -1
}

// Helper: Place a person at a specific position
func placeAt(seats []string, pos int, person string, isCircular bool) {
	if isCircular || (pos >= 0 && pos < len(seats)) {
		seats[pos] = person
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get arrangement type
	fmt.Print("Enter the arrangement type (linear/circular): ")
	arrangementType, _ := reader.ReadString('\n')
	arrangementType = strings.TrimSpace(arrangementType)

	// Get number of people
	fmt.Print("Enter the number of people: ")
	var n int
	fmt.Scanf("%d\n", &n)

	// Get constraints
	constraints := []map[string]string{}
	fmt.Print("Enter the number of constraints: ")
	var c int
	fmt.Scanf("%d\n", &c)

	fmt.Println("Enter constraints in the format: <person> <side> <next_to>")
	fmt.Println("Example: A left B (meaning A sits to the left of B)")
	for i := 0; i < c; i++ {
		fmt.Printf("Constraint %d: ", i+1)
		input, _ := reader.ReadString('\n')
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input format. Try again.")
			i--
			continue
		}
		constraints = append(constraints, map[string]string{
			"person":   parts[0],
			"side":     parts[1],
			"next_to":  parts[2],
		})
	}

	// Solve seating arrangement
	seats, err := SolveSeatingArrangement(arrangementType, n, constraints)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Seating Arrangement:", seats)
	}
}
