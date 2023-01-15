package main

func main() {
	if_else()
	switch_example()
	response := get_web_page("https://pokeapi.co/api/v2/pokemon/1")
	parse_response_body_to_json_advanced(response)
	getPokemonsSequential()
}
