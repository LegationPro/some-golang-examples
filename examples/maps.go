package examples

import "fmt"

func Maps() {
	m := make(map[string]string)

	m["username"] = "jeff"
	m["last_name"] = "jefferson"

	m_any := make(map[any]any)

	m_any[1] = 2
	m_any["n"] = "ok"

	username := m["username"]
	fmt.Printf("Username is: %v\n", username)

	any_val := m_any["n"]

	fmt.Printf("From any: %v\n", any_val)
	fmt.Printf("Another from any: %v\n", m_any[1])

	for k, v := range m {
		fmt.Printf("key: %v\nvalue: %v\n", k, v)
	}
}
