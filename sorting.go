// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	numbers := []int{5, 2, 6, 3, 1, 4}
// 	sort.Ints(numbers) // Sort in ascending order
// 	fmt.Println("Sorted Numbers:", numbers)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	names := []string{"Charlie\n", "Alice\n", "Bob\n"}
// 	sort.Strings(names) // Sort in alphabetical order
// 	fmt.Println("Sorted Names:", names)
// }

package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
