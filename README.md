# Regenerate protobuf

protoc --go_out=. *.proto

# Run
go build . && ./google_news_parser


