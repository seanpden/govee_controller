compile:
	go build -C govee/ -o ../build/govee_controller

run: compile 
	./build/govee_controller

clean:
	rm build/govee_controller

