##### Generate proto base message
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/agent.proto

##### Generate agents instance message
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/dsr_agent/dsr_agent.proto


files/agent_id/