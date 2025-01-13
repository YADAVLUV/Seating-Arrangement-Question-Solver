Seating Arrangement Solver
Overview
This Go program solves the seating arrangement for a group of people based on specified constraints. The seating arrangement can be either linear or circular. The program takes input from the user about the type of arrangement, the number of people, and any constraints on seating (e.g., who should sit next to whom and on which side).

Features
Linear and Circular Arrangements: The program supports both linear and circular seating arrangements.
Constraints: You can provide constraints that define who should sit next to whom and whether they should sit to the left or right of the other person.
Flexible Input: The program allows users to specify the seating arrangement interactively via the console.
Getting Started
Prerequisites
Ensure you have Go installed on your machine. You can check by running:

bash
Copy code
go version
If you don't have Go installed, download and install it from Go's official website.

Running the Program
Clone or download the project files to your local machine.
Open a terminal and navigate to the project directory.
Run the program with the following command:
bash
Copy code
go run main.go
Follow the prompts to enter the arrangement type, number of people, and any constraints.
Example Input/Output
Input:

sql
Copy code
Enter the arrangement type (linear/circular): circular
Enter the number of people: 5
Enter the number of constraints: 2
Enter constraints in the format: <person> <side> <next_to>
Example: A left B (meaning A sits to the left of B)
Constraint 1: A left B
Constraint 2: C right B
Output:

less
Copy code
Seating Arrangement: [C A B Person4 Person5]
How it Works
The program takes the following steps:

Input Validation: It validates the arrangement type (linear or circular), number of people, and the constraints.
Seating Logic:
Find Position: It tries to find an existing position for a person based on their constraints.
Place Person: It places people based on whether they need to sit to the left or right of another person.
Fill Remaining Seats: If any seats are still unoccupied after processing the constraints, they are filled with placeholder names (e.g., Person1, Person2, etc.).
Circular or Linear Logic: The program adjusts seat placement based on whether the arrangement is linear or circular.
Functions
SolveSeatingArrangement: Main function to solve the seating arrangement.
findPosition: Helper function to find the current position of a person in the seating array.
placeFirstAvailable: Places a person in the first available seat.
placeAt: Places a person at a specific seat based on the position and arrangement type.
Error Handling
The program includes basic error handling to ensure that:

The arrangement type is either "linear" or "circular".
The number of people is greater than zero.
The constraints are correctly formatted.
Extending the Program
You can easily extend this program by:

Adding more complex constraints.
Adding support for additional seating arrangements.
Integrating with a user interface (UI) for easier interaction.
