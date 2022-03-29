module github.com/domjesus/banco

go 1.17

replace github.com/domjesus/banco/contas => ./contas

replace github.com/domjesus/banco/clientes => ./clientes

require github.com/domjesus/banco/contas v0.0.0-00010101000000-000000000000

require github.com/domjesus/banco/clientes v0.0.0-00010101000000-000000000000 // indirect
