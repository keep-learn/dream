#编译程序
build:
	go build -race -o bin/dream .

#编译并导出
build-with-export:
	go build -race -o bin/dream .
	bin/dream export "X 信息、Y 信息；甲类、乙类"

