exec_file         =        system-monitoring
bin_path          =        ../../bin
install:
	go build -o ${bin_path}/${exec_file}
clean:
	rm ${bin_path}/${exec_file}