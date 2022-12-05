
# compile to assembly
#go build -gcflags -S main.go
go tool compile -S -N -l main.go


# compile  detail target   静态单赋值（Static Single Assignment、SSA）的特性
#GOSSAFUNC=main go build main.go