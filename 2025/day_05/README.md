# Day 05: Fresh Ingredient Database

## Challenge

We have a simple text file (problem input) that represents a database of ingredient IDs.

The database operates on ingredient IDs. It consists of a list of fresh ingredient ID ranges, a blank line, and a list of available ingredient IDs.

### Example Input

```
3-5
10-14
16-20
12-18

1
5
8
11
17
32
```

### Rules

- The fresh ID ranges are **inclusive**: the range `3-5` means that ingredient IDs 3, 4, and 5 are all fresh
- The ranges can also **overlap**: an ingredient ID is fresh if it is in any range

### Example Analysis

- Ingredient ID **1** is **spoiled** because it does not fall into any range
- Ingredient ID **5** is **fresh** because it falls into range 3-5
- Ingredient ID **8** is **spoiled**
- Ingredient ID **11** is **fresh** because it falls into range 10-14
- Ingredient ID **17** is **fresh** because it falls into range 16-20 as well as range 12-18
- Ingredient ID **32** is **spoiled**

So, in this example, **3** of the available ingredient IDs are fresh.

## Implementation Status

### Completed
- ✅ Test infrastructure
- ✅ Structs defined ([`FreshRange`](day_05/parser.go:7), [`Database`](day_05/parser.go:12))
- ✅ Helper function [`parseRange()`](day_05/parser.go:45) implemented

### TODO (Your Implementation)
- ❌ [`ParseInput()`](day_05/parser.go:28) - Parse the input file into a Database struct
- ❌ [`Database.IsFresh()`](day_05/parser.go:20) - Check if an ingredient ID is fresh
- ❌ [`CountFreshIngredients()`](day_05/solver.go:3) - Count total fresh ingredients

## Running Tests

```bash
cd day_05
go test -v
```

All tests should fail until you implement the functions above.