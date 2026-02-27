# Generic Unique Elements Filter

A small but powerful utility function in Go that removes duplicates from any slice of comparable types.

## How it works (The "Security Guard" Logic)
1. We use a `map[T]bool` called `seen` to keep track of elements we've already encountered.
2. For every element in the input:
   - We check if it exists in the `seen` map.
   - If it's **new**, we add it to the `result` slice and mark it as `true` in the map.
   - If it's **already there**, we skip it.

## Technical Skills Used
- **Generics**: Using `[T comparable]` to allow the function to work with integers, strings, etc.
- **Maps**: Efficiently checking for existence (O(1) complexity).
- **Slices**: Dynamically building the unique result.

## Example
Input: `[1, 2, 2, 3]` -> Output: `[1, 2, 3]`