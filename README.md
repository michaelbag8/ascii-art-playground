# ascii-art-playground

````markdown
# ASCII Art Recoding Sessions (Go)

---

# Recoding Session 1 — Merge Banner Maps

Write a function that joins two banner maps into one.

Given two banner maps, return a new map that contains all entries from both. If a rune appears in both maps, the entry from `priority` wins over the entry from `base`. Neither input map should be modified.

```go
func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string
````

## Examples

* MergeBanners(base, empty) → copy of base
* MergeBanners(empty, priority) → copy of priority
* If both have 'A', priority['A'] is used
* Base-only runes must still exist

## Rules

* Return a new map allocation
* Do not modify inputs
* Copy slices to avoid shared references

---

# Recoding Session 2 — Trim Trailing Spaces from Art Rows

Write a function that trims trailing spaces from each row of ASCII art.

```go
func TrimArtRows(rows []string) []string
```

## Examples

```go
TrimArtRows([]string{"_ ", "| |"})
→ []string{"_", "| |"}
```

## Rules

* Only remove trailing spaces
* Preserve leading spaces
* Return new slice
* Same length as input

---

# Recoding Session 3 — Pad Art Rows

Write a function that pads ASCII art rows to a given width.

```go
func PadArtRows(rows []string, width int) []string
```

## Examples

```go
PadArtRows([]string{"hi", "there"}, 8)
→ []string{"hi      ", "there   "}
```

## Rules

* Only pad (no truncation)
* width <= 0 → return copy as-is
* Return new slice

---

# Recoding Session 4 — Normalize Art Width

Write a function that makes all rows equal width based on the longest row.

```go
func NormalizeArtWidth(rows []string) []string
```

## Examples

```go
NormalizeArtWidth([]string{"a", "abc"})
→ []string{"a  ", "abc"}
```

## Rules

* Find max length
* Pad all rows
* Return new slice

---

# Recoding Session 5 — Reverse Art Rows

Write a function that reverses ASCII art rows.

```go
func ReverseArtRows(rows []string) []string
```

## Examples

```go
ReverseArtRows([]string{"top", "bottom"})
→ []string{"bottom", "top"}
```

## Rules

* Do not modify input
* Return new slice

---

# Recoding Session 6 — Count Visible Characters

Write a function that counts visible characters in ASCII art (ignore spaces).

```go
func CountVisibleChars(rows []string) int
```

## Examples

```go
CountVisibleChars([]string{"A A", "###"})
→ 5
```

## Rules

* Ignore `' '`
* Count all other characters

---

# Recoding Session 7 — Join Art Horizontally

Write a function that joins two ASCII art blocks side by side.

```go
func JoinArtHorizontal(left []string, right []string) []string
```

## Examples

```go
JoinArtHorizontal([]string{"A", "B"}, []string{"1", "2"})
→ []string{"A1", "B2"}
```

## Rules

* Equal length slices
* Do not modify inputs
* Return new slice

```
```
