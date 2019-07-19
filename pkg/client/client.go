package client

import "fmt"

func getJs(view string, name string) string {
	js := `
	<script>
		var chunk = document.querySelector('[name="%s"]');
		chunk.innerHTML = '%s';
	</script>
	`

	return fmt.Sprintf(js, view, name)

}

func GetView(view string, name string) string {
	js := getJs(view, name)

	return js
}
