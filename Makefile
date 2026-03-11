GO_VERSION := 1.25^  # <1>

.PHONY: install-go init-go

setup: # <2>
    install-go
    init-go

#TODO add MacOS support
install-go: # <3>
    wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
    sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
    rm go$(GO_VERSION).linux-amd64.tar.gz

init-go: # <4>
    echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
    echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc