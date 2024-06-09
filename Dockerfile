# Dockerfile
FROM golang:1.20

# Install air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# Install bash-completion and other necessary packages
RUN apt-get update && \
    apt-get install -y bash-completion curl && \
    rm -rf /var/lib/apt/lists/*

# Add git-completion script
RUN curl -o /etc/bash_completion.d/git-completion.bash https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash

# Set up bash to use bash-completion and git-completion
RUN echo "if [ -f /etc/bash_completion ]; then . /etc/bash_completion; fi" >> /etc/bash.bashrc && \
    echo "if [ -f /etc/bash_completion.d/git-completion.bash ]; then . /etc/bash_completion.d/git-completion.bash; fi" >> /etc/bash.bashrc

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go install -v golang.org/x/tools/gopls@latest

CMD ["air"]
