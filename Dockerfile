FROM golang:1.21

RUN useradd -m -s /bin/zsh go

# RUN echo "deb http://deb.debian.org/debian stretch-backports main" | tee -a /etc/apt/sources.list && \
#     mkdir -p /usr/share/man/man1 && \
RUN    apt update && apt install -y --no-install-recommends \
    git \
    ca-certificates  ca-certificates-java \
    default-jre \
    zsh \
    curl \
    wget \
    migrate

ENV JAVA_HOME="/usr/lib/jvm/java-17-openjdk-amd64/"

USER go

WORKDIR /home/go/app

RUN bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/zdharma-continuum/zinit/HEAD/scripts/install.sh)" -- \
    bash -c zinit self-update

RUN echo "zinit light zdharma/fast-syntax-highlighting" >>~/.zshrc && \
    echo "zinit light zsh-users/zsh-autosuggestions" >>~/.zshrc && \
    echo "zinit light zsh-users/zsh-completions" >>~/.zshrc && \
    echo "zinit ice depth=1; zinit light romkatv/powerlevel10k" >>~/.zshrc && \
    echo 'HISTFILE=~/zsh/.zsh_history\nHISTSIZE=10000' >>~/.zshrc && \
    echo "function save_history_after_command {\nhistory >> ~/zsh/.zsh_history\n}\nprecmd_functions+=(save_history_after_command)" >>~/.zshrc


CMD ["sh", "-c", "go mod download && tail -f /dev/null" ]

