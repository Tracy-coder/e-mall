run:
	go build && ./e-mall

idls:
	hz update -idl ./idl/user.proto

ent:
	go generate ./data/ent
.PHONY:
	run idls