# Inventory Management

An inventory management CLI application created for learning Event Sourcing.

## What I leaned
- How to store/retrieve data using Event Sourcing.

## Usage
```bash
# 1. Start the application.
go run ./cmd/

# 2. Select an operation and enter a command.
R: Receive Inventory
S: Ship Inventory
A: Inventory Adjustment
Q: Quantity On Hand
E: Events
Q: Quit
> r

# 3. Enter parameters.
> SKU: abc123
> Quantity: 30

# 4. The result is displayed.
abc123 Received: 30

# 5. Return to step 2.
```

## References
- https://youtu.be/AUj4M-st3ic
