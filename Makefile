.PHONY: install installdeps check test
install installdeps check test:
	go -C scripts run ./$@
