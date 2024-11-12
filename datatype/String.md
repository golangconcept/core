# String in go lang

Go's standard library provides a rich set of functions for working with strings.

Most of these are part of the strings package, which includes utilities for string `manipulation`, `searching`, `splitting`, `replacing`, and more.

**Here is a list of some of the most commonly used in-built string functions:**

1. strings.Contains(s, substr string) bool
2. strings.ContainsAny(s, chars string) bool
3. strings.HasPrefix(s, prefix string) bool
4. strings.HasSuffix(s, suffix string) bool
5. strings.Index(s, substr string) int
6. strings.LastIndex(s, substr string) int
7. strings.Split(s, sep string) [] string
8. strings.SplitN(s, sep string, n int) []string
9. strings.SplitAfter(s, sep string) []string
10. strings.Trim(s, cutset string) string
11. strings.TrimSpace(s string) string
12. strigns.ToLower(s string) string
13. strings.ToUpper(s string) string
14. strings.ToTitle(s string) string
15. strings.Map(fn func(r rune) rune, s string) string
16. strings.Replace(s, old, new string, n int) string
17. strings.Repeat(s string, count int) string
18. strings.Join(a[] string, sep string) string
19. strings.Compare(a, b string) int
20. strings.Count(s, substr string) int
21. strings.TrimLeft(s, cutset string) string
22. strings.TrimRight(s, cutset string) string
23. strings.Fields(s string) [] string
24. strings.IndexByte(s string, c byte) int
25. strings.ToValidUTF8(s, replacement string) string

## Start Small Projects

Once you grasp the basics, try building small projects to apply what you’ve learned:

A command-line tool (e.g., a task manager or file manipulator).
A simple web server (e.g., a REST API).
A chat application using Goroutines and Channels.

## Practice Concurrency

Go's concurrency model is one of its strongest features. To understand it fully:

- Study the `Go Scheduler and how goroutines are scheduled`.
- Implement more complex concurrent solutions, like `web scrapers`, or even explore `Worker Pools`.

### Note: Performance and Optimizations

Once you’re comfortable with Go basics, start learning about performance optimizations, profiling tools (like pprof), and benchmarking. Go is known for its speed, and understanding how to write efficient Go code is a valuable skill.

https://www.toptal.com/golang/golang-oop-tutorial
https://betterprogramming.pub/memory-optimization-and-garbage-collector-management-in-go-71da4612a960

https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html
