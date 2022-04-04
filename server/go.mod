module example.com/server

go 1.18

replace example.com/rest_server => ./rest_server

require example.com/rest_server v0.0.0-00010101000000-000000000000
